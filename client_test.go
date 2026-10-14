package glesys

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
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
	assert.Equal(t, "test-application/0.0.1 glesys-go/8.5.0", request.Header.Get("User-Agent"), "header User-Agent is correct")

	assert.NotEmpty(t, request.Header.Get("Authorization"), "header Authorization is not empty")
}

func TestEmptyUserAgent(t *testing.T) {
	client := NewClient("project-id", "api-key", "")

	request, err := client.newRequest(context.Background(), "GET", "/", nil)
	assert.NoError(t, err)
	assert.Equal(t, "glesys-go/8.5.0", request.Header.Get("User-Agent"), "header User-Agent is correct")
}

func TestGetResponseErrorMessage(t *testing.T) {
	json := `{ "response": {"status": { "code": 400, "text": "Unauthorized" } } }`
	response := http.Response{
		Body:       io.NopCloser(bytes.NewBufferString(json)),
		StatusCode: 400,
	}
	err := handleResponseError(&response)
	assert.Equal(t, "request failed with HTTP error: 400 (Unauthorized)", err.Error(), "error message is correct")
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

type trackedBody struct {
	io.Reader
	closed bool
}

func (b *trackedBody) Close() error { b.closed = true; return nil }

func TestParseResponseBody(t *testing.T) {
	for _, tc := range []struct {
		name, body string
		result     interface{}
		wantErr    bool
	}{
		{"nil result", `{"response":{}}`, nil, false},
		{"valid", `{"response":{"status":{"code":200}}}`, &struct{}{}, false},
		{"malformed", "oops", nil, true},
		{"empty", "", nil, true},
		{"provider error on HTTP 200", `{"response":{"status":{"code":409,"text":"conflict"}}}`, nil, true},
		{"too large", strings.Repeat(" ", maxResponseBytes+1), nil, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			b := &trackedBody{Reader: strings.NewReader(tc.body)}
			err := parseResponseBody(&http.Response{Body: b, StatusCode: 200}, tc.result)
			assert.Equal(t, tc.wantErr, err != nil)
			assert.True(t, b.closed, "body must be closed")
		})
	}
}

func TestHandleResponseErrorClosesBody(t *testing.T) {
	b := &trackedBody{Reader: strings.NewReader("not JSON")}
	err := handleResponseError(&http.Response{Body: b, StatusCode: 502})
	assert.EqualError(t, err, "request failed with HTTP error: 502 ()")
	assert.True(t, b.closed)
}
