package glesys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerDisk_Create(t *testing.T) {
	c := &mockClient{body: `{ "response": { "disk": { "id": "aaaaa-bbbbbb-cccccc",
		"sizeingib": 200,
		"name": "Diskett",
		"scsiid": 1,
		"type": "gold"
		}}}`}
	s := ServerDisksService{client: c}

	params := CreateServerDiskParams{
		ServerID:  "wps123456",
		Name:      "Diskett",
		SizeInGIB: 200,
	}

	disk, _ := s.Create(context.Background(), params)

	assert.Equal(t, "serverdisk/create", c.lastPath, "correct path is used")
	assert.Equal(t, 200, disk.SizeInGIB, "size is correct")
	assert.Equal(t, "Diskett", disk.Name, "correct name variable")
	assert.Equal(t, "gold", disk.Type, "correct type variable")
}

func TestServerDisk_UpdateName(t *testing.T) {
	c := &mockClient{body: `{ "response": { "disk": { "id": "aaaaa-bbbbbb-cccccc",
		"sizeingib": 200,
		"name": "Extradisk",
		"scsiid": 1
		}}}`}
	s := ServerDisksService{client: c}

	params := EditServerDiskParams{
		ID:   "aaaaa-bbbbbb-cccccc",
		Name: "Extradisk",
	}

	disk, _ := s.UpdateName(context.Background(), params)

	assert.Equal(t, "serverdisk/updatename", c.lastPath, "correct path is used")
	assert.Equal(t, 200, disk.SizeInGIB, "size is correct")
	assert.Equal(t, "Extradisk", disk.Name, "correct name variable")
}

func TestServerDisk_Reconfigure(t *testing.T) {
	c := &mockClient{body: `{ "response": { "disk": { "id": "aaaaa-bbbbbb-cccccc",
		"sizeingib": 250,
		"name": "Diskett",
		"scsiid": 1
		}}}`}
	s := ServerDisksService{client: c}

	params := EditServerDiskParams{
		ID:        "aaaaa-bbbbbb-cccccc",
		SizeInGIB: 250,
	}

	disk, _ := s.Reconfigure(context.Background(), params)

	assert.Equal(t, "serverdisk/reconfigure", c.lastPath, "correct path is used")
	assert.Equal(t, 250, disk.SizeInGIB, "size is correct")
	assert.Equal(t, "Diskett", disk.Name, "correct name variable")
}

func TestServerDisk_Delete(t *testing.T) {
	c := &mockClient{}
	s := ServerDisksService{client: c}

	id := "aaaaa-bbbbbb-cccccc"

	_ = s.Delete(context.Background(), id)

	assert.Equal(t, "serverdisk/delete", c.lastPath, "correct path is used")
}

func TestServerDisk_Limits(t *testing.T) {
	c := &mockClient{body: `{ "response": { "limits": { "minsizeingib": 10,
		"maxsizeingib": 1024,
		"currentnumdisks": 1,
		"maxnumdisks": 3
		}}}`}
	s := ServerDisksService{client: c}

	limits, _ := s.Limits(context.Background(), "wps12345")

	assert.Equal(t, "serverdisk/limits", c.lastPath, "correct path is used")
	assert.Equal(t, 1024, limits.MaxSizeInGIB, "max size is correct")
	assert.Equal(t, 3, limits.MaxNumDisks, "max number of disks correct")
}
