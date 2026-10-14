package glesys

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
)

const version = "8.5.0"

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
	BaseURL    *url.URL
	httpClient httpClientInterface
	project    string
	userAgent  string

	Databases       *DatabaseService
	DNSDomains      *DNSDomainService
	EmailDomains    *EmailDomainService
	IPs             *IPService
	LoadBalancers   *LoadBalancerService
	ObjectStorages  *ObjectStorageService
	PrivateNetworks *PrivateNetworkService
	Servers         *ServerService
	ServerDisks     *ServerDisksService
	Networks        *NetworkService
	NetworkAdapters *NetworkAdapterService
	NetworkCircuits *NetworkCircuitService
}

// NewClient creates a new Client for interacting with the GleSYS API. This is
// the main entrypoint for API interactions.
func NewClient(project, apiKey, userAgent string) *Client {
	BaseURL, _ := url.Parse("https://api.glesys.com")

	c := &Client{
		apiKey:     apiKey,
		BaseURL:    BaseURL,
		httpClient: http.DefaultClient,
		project:    project,
		userAgent:  userAgent,
	}

	c.Databases = &DatabaseService{client: c}
	c.DNSDomains = &DNSDomainService{client: c}
	c.EmailDomains = &EmailDomainService{client: c}
	c.IPs = &IPService{client: c}
	c.LoadBalancers = &LoadBalancerService{client: c}
	c.ObjectStorages = &ObjectStorageService{client: c}
	c.PrivateNetworks = &PrivateNetworkService{client: c}
	c.Servers = &ServerService{client: c}
	c.ServerDisks = &ServerDisksService{client: c}
	c.Networks = &NetworkService{client: c}
	c.NetworkAdapters = &NetworkAdapterService{client: c}
	c.NetworkCircuits = &NetworkCircuitService{client: c}

	return c
}

// SetBaseURL can be used to set a custom BaseURL
func (c *Client) SetBaseURL(bu string) error {
	url, err := url.Parse(bu)
	if err != nil {
		return err
	}
	c.BaseURL = url
	return nil
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

	if c.BaseURL != nil {
		u = c.BaseURL.ResolveReference(u)
	}

	buffer := new(bytes.Buffer)

	if params != nil {
		err = json.NewEncoder(buffer).Encode(params)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequestWithContext(ctx, method, u.String(), buffer)
	if err != nil {
		return nil, err
	}

	userAgent := strings.TrimSpace(fmt.Sprintf("%s glesys-go/%s", c.userAgent, version))

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
		return handleResponseError(response)
	}

	return parseResponseBody(response, v)
}

func handleResponseError(response *http.Response) error {
	data := struct {
		Response struct {
			Status struct {
				Text string `json:"text"`
			} `json:"status"`
		} `json:"response"`
	}{}

	err := parseResponseBody(response, &data)
	if err != nil {
		return err
	}

	return fmt.Errorf("request failed with HTTP error: %v (%v)", response.StatusCode, strings.TrimSpace(data.Response.Status.Text))
}

func parseResponseBody(response *http.Response, v interface{}) error {
	if v == nil {
		return nil
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}

var endpointPattern = regexp.MustCompile(`^[A-Za-z0-9_-]+/[A-Za-z0-9_-]+$`)
var parameterPattern = regexp.MustCompile(`^[A-Za-z0-9_-]+$`)

// ValidateEndpoint checks that endpoint is a relative service/operation such as
// "server/list", never a URL.
func ValidateEndpoint(endpoint string) error {
	if !endpointPattern.MatchString(endpoint) {
		return fmt.Errorf("endpoint must be a relative service/operation without URL syntax or traversal")
	}
	return nil
}

// Request performs a raw API request for endpoints not yet wrapped by this
// package and returns the complete JSON response. With GET, params (strings,
// json.Number or bool values) are sent as /name/value path segments; with POST
// they are sent as the JSON body.
func (c *Client) Request(ctx context.Context, method, endpoint string, params map[string]any) (json.RawMessage, error) {
	if err := ValidateEndpoint(endpoint); err != nil {
		return nil, err
	}
	method = strings.ToUpper(method)
	var data json.RawMessage
	switch method {
	case http.MethodPost:
		var payload any
		if params != nil {
			payload = params
		}
		err := c.post(ctx, endpoint, &data, payload)
		return data, err
	case http.MethodGet:
		keys := make([]string, 0, len(params))
		for k := range params {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			if !parameterPattern.MatchString(k) {
				return nil, fmt.Errorf("invalid GET parameter name")
			}
			var value string
			switch v := params[k].(type) {
			case string:
				value = v
			case json.Number:
				value = v.String()
			case bool:
				if v {
					value = "true"
				} else {
					value = "false"
				}
			default:
				return nil, fmt.Errorf("GET parameters must be strings, JSON numbers or booleans")
			}
			if value == "." || value == ".." {
				return nil, fmt.Errorf("invalid GET parameter value")
			}
			endpoint += "/" + k + "/" + url.PathEscape(value)
		}
		err := c.get(ctx, endpoint, &data)
		return data, err
	default:
		return nil, fmt.Errorf("method must be GET or POST")
	}
}
