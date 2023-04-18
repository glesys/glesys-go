package glesys

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockHTTPClient struct {
	body        string
	lastRequest *http.Request
	statusCode  int
}

func (c *mockHTTPClient) Do(request *http.Request) (*http.Response, error) {
	response := http.Response{
		StatusCode: c.statusCode,
		Body:       io.NopCloser(bytes.NewBufferString(c.body)),
	}
	c.lastRequest = request
	return &response, nil
}

func TestRequestHasCorrectHeaders(t *testing.T) {
	client := NewClient("project-id", "api-key", "test-application/0.0.1")

	request, err := client.newRequest(context.Background(), "GET", "/", nil)
	assert.NoError(t, err)

	assert.Equal(t, "application/json", request.Header.Get("Content-Type"), "header Content-Type is correct")
	assert.Equal(t, "test-application/0.0.1 glesys-go/7.0.1", request.Header.Get("User-Agent"), "header User-Agent is correct")

	assert.NotEmpty(t, request.Header.Get("Authorization"), "header Authorization is not empty")
}

func TestEmptyUserAgent(t *testing.T) {
	client := NewClient("project-id", "api-key", "")

	request, err := client.newRequest(context.Background(), "GET", "/", nil)
	assert.NoError(t, err)
	assert.Equal(t, "glesys-go/7.0.1", request.Header.Get("User-Agent"), "header User-Agent is correct")
}

func TestGetResponseErrorMessage(t *testing.T) {
	json := `{ "response": {"status": { "code": 400, "text": "Unauthorized" } } }`
	response := http.Response{
		Body:       io.NopCloser(bytes.NewBufferString(json)),
		StatusCode: 400,
	}
	err := handleResponseError(&response)
	assert.Equal(t, "Request failed with HTTP error: 400 (Unauthorized)", err.Error(), "error message is correct")
}

func TestDoDoesNotReturnErrorIfStatusIs200(t *testing.T) {
	payload := `{ "response": { "hello": "world" } }`
	client := Client{httpClient: &mockHTTPClient{body: payload, statusCode: 200}}

	request, _ := client.newRequest(context.Background(), "GET", "/", nil)
	err := client.do(request, nil)

	assert.NoError(t, err, "do does not return an error")
}

func TestDoReturnsErrorIfStatusIsNot200(t *testing.T) {
	payload := `{ "response": { "foo": "bar" } }`
	client := Client{httpClient: &mockHTTPClient{body: payload, statusCode: 500}}

	request, _ := client.newRequest(context.Background(), "GET", "/", nil)
	err := client.do(request, nil)

	assert.Error(t, err, "do returns an error")
}

func TestDoDecodesTheJsonResponseIntoAStruct(t *testing.T) {
	payload := `{ "response": { "message": "Hello World" } }`
	client := Client{httpClient: &mockHTTPClient{body: payload, statusCode: 200}}

	request, _ := client.newRequest(context.Background(), "GET", "/", nil)

	data := struct {
		Response struct {
			Message string
		}
	}{}
	err := client.do(request, &data)

	assert.NoError(t, err)
	assert.Equal(t, "Hello World", data.Response.Message, "JSON was parsed correctly")
}

func TestSetBaseURL(t *testing.T) {
	client := NewClient("project-id", "api-key", "test-application/0.0.1")

	url := "https://dev-api.glesys.test"
	err := client.SetBaseURL(url)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, client.BaseURL.String(), url, "invalid baseurl returned")
}

func TestGet(t *testing.T) {
	payload := `{ "response": { "message": "Hello World" } }`
	mockClient := mockHTTPClient{body: payload, statusCode: 200}
	client := Client{httpClient: &mockClient}

	data := struct{}{}
	client.get(context.Background(), "/foo", data)

	assert.Equal(t, "GET", mockClient.lastRequest.Method, "method used is correct")
}

func TestPost(t *testing.T) {
	payload := `{ "response": { "message": "Hello World" } }`
	mockClient := mockHTTPClient{body: payload, statusCode: 200}
	client := Client{httpClient: &mockClient}

	client.post(context.Background(), "/foo", nil, struct{ Foo string }{Foo: "bar"})

	params := struct{ Foo string }{}
	json.NewDecoder(mockClient.lastRequest.Body).Decode(&params)

	assert.Equal(t, "POST", mockClient.lastRequest.Method, "method used is correct")
	assert.Equal(t, "bar", params.Foo, "params are correct")
}
