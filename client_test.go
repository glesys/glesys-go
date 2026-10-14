package glesys

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
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

func TestClientRequest(t *testing.T) {
	var paths []string
	var bodies []map[string]interface{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.EscapedPath())
		var body map[string]interface{}
		_ = json.NewDecoder(r.Body).Decode(&body)
		bodies = append(bodies, body)
		_, _ = io.WriteString(w, `{"response":{"status":{"code":200},"id":9007199254740993123,"future":{"a":true}}}`)
	}))
	defer server.Close()

	client := NewClient("project-id", "api-key", "test")
	assert.NoError(t, client.SetBaseURL(server.URL))
	client.httpClient = server.Client()

	raw, err := client.Request(context.Background(), "post", "email/quota", map[string]interface{}{"emailaccount": "a@example.com"})
	assert.NoError(t, err)
	assert.Contains(t, string(raw), "9007199254740993123")
	assert.Contains(t, string(raw), "future")
	assert.Equal(t, "a@example.com", bodies[0]["emailaccount"])

	_, err = client.Request(context.Background(), "GET", "email/list", map[string]interface{}{"filter": "a/b@c", "domainname": "example.com"})
	assert.NoError(t, err)
	assert.Equal(t, "/email/list/domainname/example.com/filter/a%2Fb@c", paths[1])
}

func TestClientRequestValidation(t *testing.T) {
	c := &mockHTTPClient{body: `{"response":{}}`, statusCode: 200}
	client := NewClient("project-id", "api-key", "test")
	client.httpClient = c

	for _, endpoint := range []string{"https://evil.test/x", "//evil.test/x", "../email/quota", "email/../quota", "email/quota?x=y", "email/%2e%2e", "email\\quota", ""} {
		_, err := client.Request(context.Background(), "POST", endpoint, nil)
		assert.Error(t, err, endpoint)
	}
	_, err := client.Request(context.Background(), "GET", "email/list", map[string]interface{}{"bad/name": "x"})
	assert.Error(t, err)
	_, err = client.Request(context.Background(), "GET", "email/list", map[string]interface{}{"x": ".."})
	assert.Error(t, err)
	_, err = client.Request(context.Background(), "GET", "email/list", map[string]interface{}{"x": []string{"a"}})
	assert.Error(t, err)
	_, err = client.Request(context.Background(), "DELETE", "email/list", nil)
	assert.Error(t, err)
	assert.Nil(t, c.lastRequest, "invalid requests must not reach the network")
}
