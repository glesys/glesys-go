package glesys

import (
	"context"
	"time"
)

// BlockStorageService provides functions to interact with serverdisks
type BlockStorageService struct {
	client clientInterface
}

// CreateServerDiskParams specifies the details for a new serverdisk
type CreateBlockStorageParams struct {
	Name       string `json:"name"`
	Datacenter string `json:"datacenter"`
	SizeInGIB  int    `json:"sizeingib"`
	Tier       string `json:"tier"`
}

// ServerDiskDetails represents any extra disks for a server
type BlockStorageVolumeDetails struct {
	VolumeID      string                 `json:"id"`
	Name          string                 `json:"name,omitempty"`
	SizeInGIB     int                    `json:"sizeingib"`
	WWN           int                    `json:"WWN,omitempty"`
	Status        string                 `json:"status"`
	Datacenterkey string                 `json:"datacenterkey,omitempty"`
	Wwn           string                 `json:"wwn,omitempty"`
	Createdat     time.Time              `json:"createdat,omitempty"`
	Updatedat     time.Time              `json:"updatedat,omitempty"`
	Tier          BlockStorageVolumeTier `json:"tier,omitempty"`
	Attachments   []struct {
		ServerID   string `json:"serverid,omitempty"`
		DeviceName string `json:"devicename,omitempty"`
	} `json:"attachments,omitempty"`
}

type BlockStorageVolumeTier struct {
	Slug         string `json:"slug,omitempty"`
	Name         string `json:"name,omitempty"`
	MinsizeInGIB int    `json:"minsizeingib,omitempty"`
	MaxsizeInGIB int    `json:"maxsizeingib,omitempty"`
	Iopsread     int    `json:"iopsread,omitempty"`
	Iopswrite    int    `json:"iopswrite,omitempty"`
}

// ServerDiskReconfigureParams parameters for updating a ServerDisk
type EditBlockStorageParams struct {
	VolumeID  string `json:"volumeid"`
	Name      string `json:"name,omitempty"`
	SizeInGIB int    `json:"sizeingib,omitempty"`
}

type EstimateCostBlockStorageParams struct {
	ProjectKey string `json:"projectkey,omitempty"`
	VolumeID   string `json:"volumeid,omitempty"`
	SizeInGIB  int    `json:"sizeingib,omitempty"`
	Tier       string `json:"tier,omitempty"`
}

// Create - Creates an additional serverdisk using CreateServerDiskParams
func (s *BlockStorageService) Create(context context.Context, params CreateBlockStorageParams) (*BlockStorageVolumeDetails, error) {
	data := struct {
		Response struct {
			Volume BlockStorageVolumeDetails
		}
	}{}
	err := s.client.post(context, "blockstorage/create", &data, params)
	return &data.Response.Volume, err
}

// Details - Returns a BlockStorageVolume
func (s *BlockStorageService) Details(context context.Context, volumeID string) (*BlockStorageVolumeDetails, error) {
	data := struct {
		Response struct {
			Volume BlockStorageVolumeDetails
		}
	}{}
	err := s.client.post(context, "blockstorage/details", &data, struct {
		VolumeID string `json:"volumeid"`
	}{volumeID})
	return &data.Response.Volume, err
}

// List - Lists BlockStorageVolumes in project
func (s *BlockStorageService) List(context context.Context, projectKey string) (*[]BlockStorageVolumeDetails, error) {
	data := struct {
		Response struct {
			Volumes []BlockStorageVolumeDetails
		}
	}{}
	err := s.client.post(context, "blockstorage/list", &data, struct {
		ProjectKey string `json:"projectkey"`
	}{projectKey})
	return &data.Response.Volumes, err
}

// UpdateName - Modifies a serverdisk name using EditServerDiskParams
func (s *BlockStorageService) UpdateName(context context.Context, params EditBlockStorageParams) (*BlockStorageVolumeDetails, error) {
	data := struct {
		Response struct {
			Volume BlockStorageVolumeDetails
		}
	}{}
	err := s.client.post(context, "blockstorage/edit", &data, params)
	return &data.Response.Volume, err
}

// Reconfigure - Modifies a serverdisk using EditBlockStorageParams
func (s *BlockStorageService) Resize(context context.Context, params EditBlockStorageParams) (*BlockStorageVolumeDetails, error) {
	data := struct {
		Response struct {
			Volume BlockStorageVolumeDetails
		}
	}{}
	err := s.client.post(context, "blockstorage/resize", &data, params)
	return &data.Response.Volume, err
}

// Reconfigure - Modifies a serverdisk using EditBlockStorageParams
func (s *BlockStorageService) Attach(context context.Context, volumeID string, serverID string) (*BlockStorageVolumeDetails, error) {
	data := struct {
		Response struct {
			Volume BlockStorageVolumeDetails
		}
	}{}
	err := s.client.post(context, "blockstorage/attach", &data, struct {
		VolumeID string `json:"volumeid"`
		ServerID string `json:"serverid"`
	}{volumeID, serverID})
	return &data.Response.Volume, err
}

// Delete - deletes a serverdisk
func (s *BlockStorageService) Detach(context context.Context, volumeID string) error {
	return s.client.post(context, "blockstorage/detach", nil, struct {
		VolumeID string `json:"volumeid"`
	}{volumeID})
}

// Delete - deletes a serverdisk
func (s *BlockStorageService) Delete(context context.Context, volumeID string) error {
	return s.client.post(context, "blockstorage/delete", nil, struct {
		VolumeID string `json:"volumeid"`
	}{volumeID})
}

// Limits - retrieve serverdisk limits for a specific server
func (s *BlockStorageService) ListTiers(context context.Context) (*[]BlockStorageVolumeTier, error) {
	data := struct {
		Response struct {
			Tiers []BlockStorageVolumeTier
		}
	}{}
	err := s.client.get(context, "blockstorage/listtiers", &data)
	return &data.Response.Tiers, err
}

// EstimatedCost Estimate cost for a blockstorage volume, new or existing.
func (s *BlockStorageService) EstimatedCost(context context.Context, params EstimateCostBlockStorageParams) (*Billing, error) {
	data := struct {
		Response struct {
			Billing Billing
		}
	}{}
	err := s.client.post(context, "blockstorage/estimatedcost", &data, params)
	return &data.Response.Billing, err
}
