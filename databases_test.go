package glesys

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDatabaseCreate(t *testing.T) {
	c := &mockClient{body: `{"response": {"database": {
		"id": "db-1234",
		"name": "myTestDb",
		"engine": "mysql",
		"plan": {"key": "plan-1core-4gib-25gib"}
	}}}`}
	d := DatabaseService{client: c}

	params := CreateDatabaseParams{
		Name:    "myTestDb",
		PlanKey: "plan-1core-4gib-25gib",
		Engine:  "mysql",
	}

	database, _ := d.Create(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "database/create", c.lastPath, "path used is correct")
	assert.Equal(t, "myTestDb", database.Name, "Database name is correct")
	assert.Equal(t, "db-1234", database.ID, "Database id is correct")
	assert.Equal(t, "mysql", database.Engine, "Engine is correct")
	assert.Equal(t, "plan-1core-4gib-25gib", database.Plan.Key, "Plan key is correct")

}

func TestDatabasesDelete(t *testing.T) {
	c := &mockClient{}
	d := DatabaseService{client: c}

	err := d.Delete(context.Background(), "db-1234")

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "database/delete", c.lastPath, "path used is correct")
	assert.Equal(t, nil, err, "Should not get error")
}

func TestUpdateAllowlist(t *testing.T) {
	c := &mockClient{body: `{"response": {"database": {
		"id": "db-1234",
		"name": "myTestDb",
		"engine": "mysql",
		"allowlist": ["127.0.0.1", "127.0.0.2"]
	}}}`}
	d := DatabaseService{client: c}

	database, _ := d.UpdateAllowlist(context.Background(), UpdateAllowlistParams{
		ID:        "db-1234",
		AllowList: []string{"127.0.0.1", "127.0.0.2"},
	})

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "database/updateallowlist", c.lastPath, "path used is correct")
	assert.Equal(t, "myTestDb", database.Name, "Database name is correct")
	assert.Equal(t, "db-1234", database.ID, "Database id is correct")
	assert.Equal(t, []string{"127.0.0.1", "127.0.0.2"}, database.Allowlist, "Allowlist is correct")

}

func TestDatabaseList(t *testing.T) {
	c := &mockClient{body: `{"response": {"databases": [{
		"id": "db-1234",
		"name": "myTestDb",
		"engine": "mysql"
	},{
		"id": "db-56789",
		"name": "mySecondDb",
		"engine": "postgres"
	}
	]}}`}
	d := DatabaseService{client: c}

	databases, _ := d.List(context.Background())

	assert.Equal(t, "GET", c.lastMethod, "method is used correct")
	assert.Equal(t, "database/list", c.lastPath, "path used is correct")
	assert.Equal(t, 2, len(*databases), "Size of slice is correct name is correct")
	assert.Equal(t, "db-1234", (*databases)[0].ID, "First database returned")

}

func TestDatabaseDetails(t *testing.T) {
	c := &mockClient{body: `{"response": {"database": {
		"id": "db-1234",
		"name": "myTestDb",
		"engine": "mysql",
		"maintenancewindow": {"durationinminutes": 42}
	}}}`}
	d := DatabaseService{client: c}

	database, _ := d.Details(context.Background(), "db-1234")

	assert.Equal(t, "GET", c.lastMethod, "method is used correct")
	assert.Equal(t, "database/details/id/db-1234", c.lastPath, "path used is correct")
	assert.Equal(t, "myTestDb", database.Name, "Database name is correct")
	assert.Equal(t, "db-1234", database.ID, "Database id is correct")
	assert.Equal(t, 42, database.MaintenanceWindow.DurationInMinutes, "Engine is correct")

}

func TestDatabaseConnectionString(t *testing.T) {
	c := &mockClient{body: `{"response": {"connectiondetails": {
		"connectionstring": "mysql://dbadmin:password@db-12345.database-v1.glesys.com:3306/defaultdb?ssl-mode=required" }}}`}
	d := DatabaseService{client: c}

	database, _ := d.ConnectionString(context.Background(), "db-12345")

	assert.Equal(t, "GET", c.lastMethod, "method is used correct")
	assert.Equal(t, "database/connectiondetails/id/db-12345", c.lastPath, "path used is correct")
	assert.Equal(t, "mysql://dbadmin:password@db-12345.database-v1.glesys.com:3306/defaultdb?ssl-mode=required", database.ConnectionString, "Database name is correct")

}

func TestDatabasePlans(t *testing.T) {
	c := &mockClient{body: `{"response": {"plans": [
      {
        "key": "plan-1core-4gib-25gib",
        "cpucores": 1,
        "memoryingib": 4,
        "storageingib": 25
      },
      {
        "key": "plan-1core-4gib-50gib",
        "cpucores": 1,
        "memoryingib": 4,
        "storageingib": 50
      }]}}`}
	d := DatabaseService{client: c}

	plans, _ := d.ListPlans(context.Background())

	assert.Equal(t, "GET", c.lastMethod, "method is used correct")
	assert.Equal(t, "database/listplans", c.lastPath, "path used is correct")
	assert.Equal(t, 2, len(*plans), "Size of slice is correct name is correct")
	assert.Equal(t, "plan-1core-4gib-25gib", (*plans)[0].Key, "First database returned")

}

func TestEstimatedCost(t *testing.T) {
	c := &mockClient{body: `{"response": {"billing": {
      "currency": "SEK",
      "current": {
        "price": 0,
        "discount": 0,
        "total": 0
      },
      "estimated": {
        "price": 355.51,
        "discount": 0,
        "total": 355.51
      },
      "diff": {
        "price": 355.51,
        "discount": 0,
        "total": 355.51
      }}}}`}
	d := DatabaseService{client: c}

	billing, _ := d.EstimatedCost(context.Background(), EstimatedCostParams{
		PlanKey: "plan-1core-4gib-25gib",
	})

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "database/estimatedcost", c.lastPath, "path used is correct")
	assert.Equal(t, "SEK", billing.Currency, "Database name is correct")
	assert.Equal(t, 355.51, billing.Estimated.Price, "Database id is correct")
	assert.Equal(t, 0, billing.Diff.Discount, "Database id is correct")

}
