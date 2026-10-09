package glesys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockStorageCreate(t *testing.T) {
	c := &mockClient{body: `{
		   "response": {
		     "volume": {
					  "id": "bsv-ab3ab",
					  "name": "volumeone",
					  "sizeingib": 101,
					  "status": "CREATING",
					  "datacenterkey": "dc-fbg1",
					  "wwn": "None",
					  "createdat": "2026-10-08T06:52:07+00:00",
					  "updatedat": "2026-10-08T06:52:07+00:00",
					  "tier": {
		          "slug": "balanced",
							"name": "Balanced",
							"minsizeingib": 1,
							"maxsizeingib": 2048,
							"iopsread": 12000,
							"iopswrite": 12000
		        },
						"attachments": []
		     }
		  }
		}`}
	//  }`}
	b := BlockStorageService{client: c}

	params := CreateBlockStorageParams{
		Name:       "volumeone",
		Datacenter: "dc-fbg1",
		SizeInGIB:  101,
		Tier:       "balanced",
	}

	volume, _ := b.Create(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "blockstorage/create", c.lastPath, "path used is correct")
	assert.Equal(t, "volumeone", volume.Name, "Blockstorage name is correct")
	assert.Equal(t, "bsv-ab3ab", volume.VolumeID, "Blockstorage id is correct")
	assert.Equal(t, "Balanced", volume.Tier.Name, "Attribute is correct")
	assert.Equal(t, 101, volume.SizeInGIB, "Attribute is correct")
}

func TestBlockStorageDelete(t *testing.T) {
	c := &mockClient{}
	d := BlockStorageService{client: c}

	d.Delete(context.Background(), "db-1234")

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "blockstorage/delete", c.lastPath, "path used is correct")
}

func TestBlockStorageList(t *testing.T) {
	c := &mockClient{body: `{
		  "response": {
				"volumes": [
				{
					"id": "bsv-1234",
					"name": "volume-one",
					"sizeingib": 101,
					"status": "READY",
					"datacenterkey": "dc-fbg1",
					"wwn": "3ebadbeef999abc",
					"createdat": "2026-06-15T12:35:09+00:00",
					"updatedat": "2026-06-15T12:40:34+00:00",
					"tier": {
						"slug": "capacity",
						"name": "Capacity",
						"minsizeingib": 1,
						"maxsizeingib": 2048,
						"iopsread": 1200,
						"iopswrite": 1200
					},
					"attachments":
						[
		          {"serverid": "kvm123456", "devicename": "sdb"}
		        ]
				  },
					{
						"id": "bsv-56789",
						"name": "volume-two",
						"sizeingib": 101,
						"status": "READY",
						"datacenterkey": "dc-fbg1",
						"wwn": "3ebadbeef9a9123",
						"createdat": "2026-06-25T10:35:09+00:00",
						"updatedat": "2026-06-25T12:40:34+00:00",
						"tier": {
							"slug": "capacity",
							"name": "Capacity",
							"minsizeingib": 1,
							"maxsizeingib": 2048,
							"iopsread": 1200,
							"iopswrite": 1200
						},
						"attachments":
							[{"serverid": "kvm123456", "devicename": "sdc"}]
					}]
		  }
		}`}

	bsv := BlockStorageService{client: c}

	volumes, err := bsv.List(context.Background(), "cl12345")

	assert.NoError(t, err)
	assert.NotNil(t, volumes)
	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "blockstorage/list", c.lastPath, "path used is correct")
	assert.Equal(t, 2, len(*volumes), "Size of slice is correct name is correct")
	assert.Equal(t, "bsv-1234", (*volumes)[0].VolumeID, "ID is correct")
}

func TestBlockStorageDetails(t *testing.T) {
	c := &mockClient{body: `{"response": {"volume": {
		"id": "bsv-ab123ba",
		"name": "volumeone",
		"sizeingib": 101,
		"tier": {
		  "slug": "capacity",
		  "name": "Capacity",
		  "minsizeingib": 1,
		  "maxsizeingib": 2048,
		  "iopsread": 1200,
		  "iopswrite": 1200
		}
	}}}`}
	b := BlockStorageService{client: c}

	volume, err := b.Details(context.Background(), "bsv-ab123ba")

	assert.NoError(t, err)
	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "blockstorage/details", c.lastPath, "path used is correct")
	assert.Equal(t, "volumeone", volume.Name, "Blockstorage name is correct")
	assert.Equal(t, "bsv-ab123ba", volume.VolumeID, "Blockstorage id is correct")
	assert.Equal(t, 101, volume.SizeInGIB, "Size is correct")
}

func TestBlockStorageTiers(t *testing.T) {
	c := &mockClient{body: `{"response": {"tiers": [
		  {"slug": "velocity",
		    "name": "Velocity",
		    "minsizeingib": 1,
		    "maxsizeingib": 2048,
		    "iopsread": 120000,
		    "iopswrite": 30000},
		   {"slug": "balanced",
		    "name": "Balanced",
		    "minsizeingib": 1,
		    "maxsizeingib": 2048,
		    "iopsread": 12000,
		    "iopswrite": 12000},
		   {"slug": "capacity",
		    "name": "Capacity",
		    "minsizeingib": 1,
		    "maxsizeingib": 2048,
		    "iopsread": 1200,
		    "iopswrite": 1200
		  }]}}`}
	b := BlockStorageService{client: c}

	tiers, _ := b.ListTiers(context.Background())

	assert.Equal(t, "GET", c.lastMethod, "method is used correct")
	assert.Equal(t, "blockstorage/listtiers", c.lastPath, "path used is correct")
	assert.Equal(t, 3, len(*tiers), "Size of slice is correct name is correct")
	assert.Equal(t, "Velocity", (*tiers)[0].Name, "First tier returned")
}

func TestBlockStorageEstimatedCost(t *testing.T) {
	c := &mockClient{body: `{"response": {"billing": {
      "currency": "SEK",
      "current": {
        "price": 0,
        "discount": 0,
        "total": 0
      },
      "estimated": {
        "price": 180.0,
        "discount": 0,
        "total": 180.0
		  },
      "diff": {
        "price": 180.0,
        "discount": 0,
        "total": 180.0
	    }}}}`}
	b := BlockStorageService{client: c}

	billing, _ := b.EstimatedCost(context.Background(), EstimateCostBlockStorageParams{
		SizeInGIB: 200,
		Tier:      "balanced",
	})

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "blockstorage/estimatedcost", c.lastPath, "path used is correct")
	assert.Equal(t, "SEK", billing.Currency, "Blockstorage name is correct")
	assert.Equal(t, 180.0, billing.Estimated.Price, "Blockstorage id is correct")
	assert.Equal(t, 0, billing.Diff.Discount, "Blockstorage id is correct")
}
