package glesys

import "context"

// ServerDisksService provides functions to interact with serverdisks
type ServerDisksService struct {
	client clientInterface
}

// CreateServerDiskParams specifies the details for a new serverdisk
type CreateServerDiskParams struct {
	Name      string `json:"name"`
	ServerID  string `json:"serverid"`
	SizeInGIB int    `json:"sizeingib"`
}

// ServerDiskDetails represents any extra disks for a server
type ServerDiskDetails struct {
	ID        string `json:"id"`
	Name      string `json:"name,omitempty"`
	SizeInGIB int    `json:"sizeingib"`
	SCSIID    int    `json:"scsiid"`
}

// ServerDiskReconfigureParams parameters for updating a ServerDisk
type EditServerDiskParams struct {
	ID        string `json:"id"`
	Name      string `json:"name,omitempty"`
	SizeInGIB int    `json:"sizeingib,omitempty"`
}

// ServerDiskLimitsDetails represents the disk limits for a server
type ServerDiskLimitsDetails struct {
	MinSizeInGIB    int `json:"minsizeingib"`
	MaxSizeInGIB    int `json:"maxsizeingib"`
	MaxNumDisks     int `json:"maxnumdisks"`
	CurrentNumDisks int `json:"currentnumdisks"`
}

// Create - Creates an additional serverdisk using CreateServerDiskParams
func (s *ServerDisksService) Create(context context.Context, params CreateServerDiskParams) (*ServerDiskDetails, error) {
	data := struct {
		Response struct {
			Disk ServerDiskDetails
		}
	}{}
	err := s.client.post(context, "serverdisk/create", &data, params)
	return &data.Response.Disk, err
}

// UpdateName - Modifies a serverdisk name using EditServerDiskParams
func (s *ServerDisksService) UpdateName(context context.Context, params EditServerDiskParams) (*ServerDiskDetails, error) {
	data := struct {
		Response struct {
			Disk ServerDiskDetails
		}
	}{}
	err := s.client.post(context, "serverdisk/updatename", &data, params)
	return &data.Response.Disk, err
}

// Reconfigure - Modifies a serverdisk using EditServerDiskParams
func (s *ServerDisksService) Reconfigure(context context.Context, params EditServerDiskParams) (*ServerDiskDetails, error) {
	data := struct {
		Response struct {
			Disk ServerDiskDetails
		}
	}{}
	err := s.client.post(context, "serverdisk/reconfigure", &data, params)
	return &data.Response.Disk, err
}

// Delete - deletes a serverdisk
func (s *ServerDisksService) Delete(context context.Context, diskID string) error {
	return s.client.post(context, "serverdisk/delete", nil, struct {
		DiskID string `json:"id"`
	}{diskID})
}

// Limits - retrieve serverdisk limits for a specific server
func (s *ServerDisksService) Limits(context context.Context, serverID string) (*ServerDiskLimitsDetails, error) {
	data := struct {
		Response struct {
			Limits ServerDiskLimitsDetails
		}
	}{}
	err := s.client.post(context, "serverdisk/limits", &data,
		struct {
			ServerID string `json:"serverid"`
		}{serverID})
	return &data.Response.Limits, err
}
