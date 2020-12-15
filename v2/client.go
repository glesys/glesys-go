package glesys

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const version = "2.5.0"

type httpClientInterface interface {
	Do(*http.Request) (*http.Response, error)
}

type clientInterface interface {
	get(ctx context.Context, path string, v interface{}) error
	post(ctx context.Context, path string, v interface{}, params interface{}) error
}

// Client is used to interact with the GleSYS API
type Client struct {
	apiKey     string
	baseURL    *url.URL
	httpClient httpClientInterface
	project    string
	userAgent  string

	DNSDomains      *DNSDomainService
	EmailDomains    *EmailDomainService
	IPs             *IPService
	LoadBalancers   *LoadBalancerService
	ObjectStorages  *ObjectStorageService
	Servers         *ServerService
	Networks        *NetworkService
	NetworkAdapters *NetworkAdapterService
}

// NewClient creates a new Client for interacting with the GleSYS API. This is
// the main entrypoint for API interactions.
func NewClient(project, apiKey, userAgent string) *Client {
	baseURL, _ := url.Parse("https://api.glesys.com")

	c := &Client{
		apiKey:     apiKey,
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
		project:    project,
		userAgent:  userAgent,
	}

	c.DNSDomains = &DNSDomainService{client: c}
	c.EmailDomains = &EmailDomainService{client: c}
	c.IPs = &IPService{client: c}
	c.LoadBalancers = &LoadBalancerService{client: c}
	c.ObjectStorages = &ObjectStorageService{client: c}
	c.Servers = &ServerService{client: c}
	c.Networks = &NetworkService{client: c}
	c.NetworkAdapters = &NetworkAdapterService{client: c}

	return c
}

func (c *Client) get(ctx context.Context, path string, v interface{}) error {
	request, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return err
	}
	return c.do(request, v)
}

func (c *Client) post(ctx context.Context, path string, v interface{}, params interface{}) error {
	request, err := c.newRequest(ctx, "POST", path, params)
	if err != nil {
		return err
	}
	return c.do(request, v)
}

func (c *Client) newRequest(ctx context.Context, method, path string, params interface{}) (*http.Request, error) {
	u, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	if c.baseURL != nil {
		u = c.baseURL.ResolveReference(u)
	}

	buffer := new(bytes.Buffer)

	if params != nil {
		err = json.NewEncoder(buffer).Encode(params)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(method, u.String(), buffer)
	if err != nil {
		return nil, err
	}

	userAgent := strings.TrimSpace(fmt.Sprintf("%s glesys-go/%s", c.userAgent, version))

	request = request.WithContext(ctx)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", userAgent)
	request.SetBasicAuth(c.project, c.apiKey)

	return request, nil
}

func (c *Client) do(request *http.Request, v interface{}) error {
	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return c.handleResponseError(response)
	}

	return c.parseResponseBody(response, v)
}

func (c *Client) handleResponseError(response *http.Response) error {
	data := struct {
		Response struct {
			Status struct {
				Text string `json:"text"`
			} `json:"status"`
		} `json:"response"`
	}{}

	err := c.parseResponseBody(response, &data)
	if err != nil {
		return err
	}

	return fmt.Errorf("Request failed with HTTP error: %v (%v)", response.StatusCode, strings.TrimSpace(data.Response.Status.Text))
}

func (c *Client) parseResponseBody(response *http.Response, v interface{}) error {
	if v == nil {
		return nil
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}
