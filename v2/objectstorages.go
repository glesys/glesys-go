package glesys

import (
	"context"
)

// ObjectStorageService provides functions to interact with Networks
type ObjectStorageService struct {
	client clientInterface
}

// ObjectStorageInstance represents a Object Storage Instance
type ObjectStorageInstance struct {
	Created     string                    `json:"created"`
	Credentials []ObjectStorageCredential `json:"credentials"`
	DataCenter  string                    `json:"datacenter"`
	Description string                    `json:"description,omitempty"`
	InstanceID  string                    `json:"id"`
}

// ObjectStorageCredential represents a credential for an Object Storage Instance
type ObjectStorageCredential struct {
	AccessKey    string `json:"accesskey"`
	Created      string `json:"created"`
	CredentialID string `json:"id"`
	Description  string `json:"description,omitempty"`
	SecretKey    string `json:"secretkey"`
}

// CreateObjectStorageInstanceParams is used when creating a new instance
type CreateObjectStorageInstanceParams struct {
	DataCenter  string `json:"datacenter"`
	Description string `json:"description,omitempty"`
}

// EditObjectStorageInstanceParams is used when editing an existing instance
type EditObjectStorageInstanceParams struct {
	Description string `json:"description,omitempty"`
	InstanceID  string `json:"instanceid"`
}

// CreateObjectStorageCredentialParams is used when creating a new credential
type CreateObjectStorageCredentialParams struct {
	InstanceID  string `json:"instanceid"`
	Description string `json:"description,omitempty"`
}

// DeleteObjectStorageCredentialParams is used when creating a new credential
type DeleteObjectStorageCredentialParams struct {
	InstanceID   string `json:"instanceid"`
	CredentialID string `json:"credentialid"`
}

// CreateInstance creates a new Object Storage Instance
func (s *ObjectStorageService) CreateInstance(context context.Context, params CreateObjectStorageInstanceParams) (*ObjectStorageInstance, error) {
	data := struct {
		Response struct {
			Instance ObjectStorageInstance
		}
	}{}
	err := s.client.post(context, "objectstorage/createinstance", &data, params)
	return &data.Response.Instance, err
}

// InstanceDetails returns detailed information about an Object Storage Instance
func (s *ObjectStorageService) InstanceDetails(context context.Context, instanceID string) (*ObjectStorageInstance, error) {
	data := struct {
		Response struct {
			Instance ObjectStorageInstance
		}
	}{}
	err := s.client.post(context, "objectstorage/instancedetails", &data, struct {
		InstanceID string `json:"instanceid"`
	}{instanceID})
	return &data.Response.Instance, err
}

// DeleteInstance deletes an Object Storage Instance
// !!!THIS WILL DELETE ALL CREDENTIALS AND DATA FOR THIS INSTANCE!!!
func (s *ObjectStorageService) DeleteInstance(context context.Context, instanceID string) error {
	return s.client.post(context, "objectstorage/deleteinstance", nil, struct {
		InstanceID string `json:"instanceid"`
	}{instanceID})
}

// EditInstance Updates the description for an instance
func (s *ObjectStorageService) EditInstance(context context.Context, params EditObjectStorageInstanceParams) (*ObjectStorageInstance, error) {
	data := struct {
		Response struct {
			Instance ObjectStorageInstance
		}
	}{}
	err := s.client.post(context, "objectstorage/editinstance", &data, struct {
		EditObjectStorageInstanceParams
	}{params})
	return &data.Response.Instance, err
}

// ListInstances returns a list of Object Storage Instances  available under your account
func (s *ObjectStorageService) ListInstances(context context.Context) (*[]ObjectStorageInstance, error) {
	data := struct {
		Response struct {
			Instances []ObjectStorageInstance
		}
	}{}

	err := s.client.get(context, "objectstorage/listinstances", &data)
	return &data.Response.Instances, err
}

// CreateCredential creates a Credential for an Object Storage Instance
func (s *ObjectStorageService) CreateCredential(context context.Context, params CreateObjectStorageCredentialParams) (*ObjectStorageCredential, error) {
	data := struct {
		Response struct {
			Credential ObjectStorageCredential
		}
	}{}
	err := s.client.post(context, "objectstorage/createcredential", &data, params)
	return &data.Response.Credential, err
}

// DeleteCredential deletes a Credential for an Object Storage Instance
func (s *ObjectStorageService) DeleteCredential(context context.Context, params DeleteObjectStorageCredentialParams) error {
	return s.client.post(context, "objectstorage/deletecredential", nil, params)
}
