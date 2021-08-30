package glesys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectStorageInstance_Create(t *testing.T) {
	c := &mockClient{body: `{ "response": { "instance": { "id": "os-ab123", "created": "2020-05-29T13:37:00+00:00",
		"description": "OSI Test", "datacenter": "dc-sto1", "credentials": [{"id": "376af022-c81e-4a02-9f28-0dce6ba24387",
		"accesskey": "ABC123QWERTYAOEU987", "description": "None", "created": "2020-05-29T13:37:00+00:00",
		"secretkey": "someverylongrandomkey123445s3cure"}]}}}`}
	s := ObjectStorageService{client: c}

	params := CreateObjectStorageInstanceParams{
		DataCenter:  "dc-sto1",
		Description: "OSI Test",
	}

	instance, _ := s.CreateInstance(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "objectstorage/createinstance", c.lastPath, "path used is correct")
	assert.Equal(t, "os-ab123", instance.InstanceID, "Objectstorage Instance ID is correct")
	assert.Equal(t, "dc-sto1", instance.DataCenter, "DataCenter is correct")
	assert.Equal(t, "OSI Test", instance.Description, "Description is correct")
}

func TestObjectStorageInstance_Delete(t *testing.T) {
	c := &mockClient{}
	s := ObjectStorageService{client: c}

	s.DeleteInstance(context.Background(), "os-ab123")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "objectstorage/deleteinstance", c.lastPath, "path used is correct")
}

func TestObjectStorageInstance_Details(t *testing.T) {
	c := &mockClient{body: `{ "response": { "instance": { "id": "os-ab123", "created": "2020-05-29T13:37:00+00:00",
		"description": "OSI Test", "datacenter": "dc-sto1", "credentials": [{"id": "376af022-c81e-4a02-9f28-0dce6ba24387",
		"accesskey": "ABC123QWERTYAOEU987", "description": "None", "created": "2020-05-29T13:37:00+00:00",
		"secretkey": "someverylongrandomkey123445s3cure"}]}}}`}
	s := ObjectStorageService{client: c}

	instance, _ := s.InstanceDetails(context.Background(), "os-ab123")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "objectstorage/instancedetails", c.lastPath, "path used is correct")
	assert.Equal(t, "dc-sto1", instance.DataCenter, "DataCenter is correct")
	assert.Equal(t, "OSI Test", instance.Description, "Description is correct")
	assert.Equal(t, "os-ab123", instance.InstanceID, "ID is correct")
	assert.Equal(t, "ABC123QWERTYAOEU987", instance.Credentials[0].AccessKey, "Access Key is correct")
}

func TestObjectStorageInstance_Edit(t *testing.T) {
	c := &mockClient{body: `{ "response": { "instance": { "id": "os-ab123", "created": "2020-05-29T13:37:00+00:00",
		"description": "My OSI Test", "datacenter": "dc-sto1", "credentials": [{"id": "376af022-c81e-4a02-9f28-0dce6ba24387",
		"accesskey": "ABC123QWERTYAOEU987", "description": "None", "created": "2020-05-29T13:37:00+00:00",
		"secretkey": "someverylongrandomkey123445s3cure"}]}}}`}
	s := ObjectStorageService{client: c}

	params := EditObjectStorageInstanceParams{
		Description: "My OSI Test",
		InstanceID:  "os-ab123",
	}

	instance, _ := s.EditInstance(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "objectstorage/editinstance", c.lastPath, "path used is correct")
	assert.Equal(t, "os-ab123", instance.InstanceID, "ID is correct")
	assert.Equal(t, "dc-sto1", instance.DataCenter, "DataCenter is correct")
	assert.Equal(t, "My OSI Test", instance.Description, "Description is correct")
}

func TestObjectStorageCredential_Create(t *testing.T) {
	c := &mockClient{body: `{ "response": { "credential": { "id": "376af022-c81e-4a02-9f28-0dce6ba24399",
		"created": "2020-05-29T23:37:00+00:00", "description": "Key2", "accesskey": "AOEUKEY123546",
		"secretkey": "superlongsecretkey123123123"} } }`}
	s := ObjectStorageService{client: c}

	params := CreateObjectStorageCredentialParams{
		InstanceID:  "os-ab123",
		Description: "Key2",
	}
	credential, _ := s.CreateCredential(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "objectstorage/createcredential", c.lastPath, "path used is correct")
	assert.Equal(t, "Key2", credential.Description, "Description is correct")
	assert.Equal(t, "376af022-c81e-4a02-9f28-0dce6ba24399", credential.CredentialID, "ID is correct")
}

func TestObjectStorageCredential_Delete(t *testing.T) {
	c := &mockClient{}

	s := ObjectStorageService{client: c}

	params := DeleteObjectStorageCredentialParams{
		CredentialID: "376af022-c81e-4a02-9f28-0dce6ba24399",
		InstanceID:   "os-ab123",
	}

	s.DeleteCredential(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "objectstorage/deletecredential", c.lastPath, "path used is correct")
}
