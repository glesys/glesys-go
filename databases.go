package glesys

import (
	"context"
	"fmt"
)

// DatabaseService provides functions to interact with Databases
type DatabaseService struct {
	client clientInterface
} // Database represents a Database
type Database struct {
	DataCenterKey string `json:"datacenterkey"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	Engine        string `json:"engine"`
	EngineVersion string `json:"engineversion"`
}

// DatabaseDetails represents the detailed version of a database
type DatabaseDetails struct {
	DataCenterKey     string       `json:"datacenterkey"`
	ID                string       `json:"id"`
	Name              string       `json:"name"`
	Engine            string       `json:"engine"`
	EngineVersion     string       `json:"engineversion"`
	Fqdn              string       `json:"fqdn"`
	Status            string       `json:"status"`
	Allowlist         []string     `json:"allowlist"`
	Plan              DatabasePlan `json:"plan"`
	MaintenanceWindow struct {
		WeekDay           string `json:"weekday"`
		StartTime         string `json:"starttime"`
		DurationInMinutes int    `json:"durationinminutes"`
	} `json:"maintenancewindow"`
}

type DatabasePlan struct {
	Key          string `json:"key"`
	CpuCores     int    `json:"cpucores"`
	MemoryInGib  int    `json:"memoryingib"`
	StorageInGib int    `json:"storageingib"`
}

type Billing struct {
	Currency string `json:"currency"`
	Current  struct {
		Price    float64 `json:"price"`
		Discount int     `json:"discount"`
		Total    float64 `json:"total"`
	} `json:"current"`
	Estimated struct {
		Price    float64 `json:"price"`
		Discount int     `json:"discount"`
		Total    float64 `json:"total"`
	} `json:"estimated"`
	Diff struct {
		Price    float64 `json:"price"`
		Discount int     `json:"discount"`
		Total    float64 `json:"total"`
	} `json:"diff"`
}

// CreateDatabaseParams is used when creating a new database
type CreateDatabaseParams struct {
	PlanKey       string `json:"plankey"`
	Engine        string `json:"engine"`
	EngineVersion string `json:"engineversion"`
	DataCenterKey string `json:"datacenterkey"`
	Name          string `json:"name"`
}

// UpdateAllowlistParams is used to update the allowlist for a database instance
type UpdateAllowlistParams struct {
	ID        string   `json:"id"`
	AllowList []string `json:"allowlist"`
}

// ConnectionDetails returns the connection string
type ConnectionDetails struct {
	ConnectionString string `json:"connectionstring"`
}

type EstimatedCostParams struct {
	ID      string `json:"databaseid,omitempty"`
	PlanKey string `json:"plankey"`
}

// Create Creates a database
func (db *DatabaseService) Create(context context.Context, params CreateDatabaseParams) (*DatabaseDetails, error) {
	data := struct {
		Response struct {
			Database DatabaseDetails
		}
	}{}
	err := db.client.post(context, "database/create", &data, params)
	return &data.Response.Database, err
}

// Delete a database
func (db *DatabaseService) Delete(context context.Context, databaseID string) error {
	return db.client.post(context, "database/delete", nil, struct {
		ID string `json:"id"`
	}{databaseID})
}

// UpdateAllowlist Update the allow list for a database instance
func (db *DatabaseService) UpdateAllowlist(context context.Context, params UpdateAllowlistParams) (*DatabaseDetails, error) {
	data := struct {
		Response struct {
			Database DatabaseDetails
		}
	}{}
	err := db.client.post(context, "database/updateallowlist", &data, params)
	return &data.Response.Database, err
}

// List returns a list of databases
func (db *DatabaseService) List(context context.Context) (*[]Database, error) {
	data := struct {
		Response struct {
			Databases []Database
		}
	}{}
	err := db.client.get(context, "database/list", &data)
	return &data.Response.Databases, err
}

// Details returns detailed information about one database
func (db *DatabaseService) Details(context context.Context, databaseID string) (*DatabaseDetails, error) {
	data := struct {
		Response struct {
			Database DatabaseDetails
		}
	}{}
	err := db.client.get(context, fmt.Sprintf("database/details/id/%s", databaseID), &data)
	return &data.Response.Database, err
}

// ConnectionString Get the connection string for a database instance.
func (db *DatabaseService) ConnectionString(context context.Context, databaseID string) (*ConnectionDetails, error) {
	data := struct {
		Response struct {
			ConnectionDetails ConnectionDetails
		}
	}{}
	err := db.client.get(context, fmt.Sprintf("database/connectiondetails/id/%s", databaseID), &data)
	return &data.Response.ConnectionDetails, err
}

// ListPlans Get a list of available databases plans.
func (db *DatabaseService) ListPlans(context context.Context) (*[]DatabasePlan, error) {
	data := struct {
		Response struct {
			Plans []DatabasePlan
		}
	}{}
	err := db.client.get(context, "database/listplans", &data)
	return &data.Response.Plans, err
}

// EstimatedCost Estimate cost for a database instance, new or existing.
func (db *DatabaseService) EstimatedCost(context context.Context, params EstimatedCostParams) (*Billing, error) {
	data := struct {
		Response struct {
			Billing Billing
		}
	}{}
	err := db.client.post(context, "database/estimatedcost", &data, params)
	return &data.Response.Billing, err
}
