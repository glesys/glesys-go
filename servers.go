package glesys

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/imdario/mergo"
)

// ServerService provides functions to interact with servers
type ServerService struct {
	client clientInterface
}

// Server is a simplified version of a server
type Server struct {
	DataCenter string `json:"datacenter"`
	Hostname   string `json:"hostname"`
	ID         string `json:"serverid"`
	Platform   string `json:"platform"`
}

// User represents a system user when creating servers (currently supported in KVM)
type User struct {
	Username     string   `json:"username"`
	PublicKeys   []string `json:"sshkeys"`
	Password     string   `json:"password"`
}

// ServerDetails is a more complete representation of a server
type ServerDetails struct {
	CPU         int    `json:"cpucores"`
	Bandwidth   int    `json:"bandwidth"`
	DataCenter  string `json:"datacenter"`
	Description string `json:"description"`
	Hostname    string `json:"hostname"`
	ID          string `json:"serverid"`
	IPList      []IP   `json:"iplist"`
	Platform    string `json:"platform"`
	Memory      int    `json:"memorysize"`
	State       string `json:"state"`
	Storage     int    `json:"disksize"`
	Template    string `json:"templatename"`
}

// IsLocked returns true if the server is currently locked, false otherwise
func (sd *ServerDetails) IsLocked() bool {
	return sd.State == "locked"
}

// IsRunning returns true if the server is currently running, false otherwise
func (sd *ServerDetails) IsRunning() bool {
	return sd.State == "running"
}

// CreateServerParams is used when creating a new server
type CreateServerParams struct {
	Bandwidth    int    `json:"bandwidth"`
	CampaignCode string `json:"campaigncode,omitempty"`
	CPU          int    `json:"cpucores"`
	DataCenter   string `json:"datacenter"`
	Description  string `json:"description,omitempty"`
	Hostname     string `json:"hostname"`
	IPv4         string `json:"ip"`
	IPv6         string `json:"ipv6"`
	Memory       int    `json:"memorysize"`
	Password     string `json:"rootpassword,omitempty"`
	Platform     string `json:"platform"`
	PublicKey    string `json:"sshkey,omitempty"`
	Storage      int    `json:"disksize"`
	Template     string `json:"templatename"`
	Users        []User `json:"users"`
}

// EditServerParams is used when editing an existing server
type EditServerParams struct {
	Bandwidth   int    `json:"bandwidth,omitempty"`
	CPU         int    `json:"cpucores,omitempty"`
	Description string `json:"description,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
	Memory      int    `json:"memorysize,omitempty"`
	Storage     int    `json:"disksize,omitempty"`
}

// WithDefaults populates the parameters with default values. Existing
// parameters will not be overwritten.
func (p CreateServerParams) WithDefaults() CreateServerParams {
	defaults := CreateServerParams{
		Bandwidth:  100,
		CPU:        2,
		DataCenter: "Falkenberg",
		Hostname:   generateHostname(),
		IPv4:       "any",
		IPv6:       "any",
		Memory:     2048,
		Platform:   "OpenVZ",
		Storage:    50,
		Template:   "Debian 8 64-bit",
	}
	mergo.Merge(&p, defaults)
	return p
}

// DestroyServerParams is used when destroying a server
type DestroyServerParams struct {
	KeepIP bool `json:"keepip"`
}

// StopServerParams is used when stopping a server. Supported types are `soft`
// `hard` and `reboot`.
type StopServerParams struct {
	Type string `json:"type"`
}

// Create creates a new server
func (s *ServerService) Create(context context.Context, params CreateServerParams) (*ServerDetails, error) {
	data := struct {
		Response struct {
			Server ServerDetails
		}
	}{}
	err := s.client.post(context, "server/create", &data, params)
	return &data.Response.Server, err
}

// Destroy deletes a server
func (s *ServerService) Destroy(context context.Context, serverID string, params DestroyServerParams) error {
	return s.client.post(context, "server/destroy", nil, struct {
		DestroyServerParams
		ServerID string `json:"serverid"`
	}{params, serverID})
}

// Details returns detailed information about one server
func (s *ServerService) Details(context context.Context, serverID string) (*ServerDetails, error) {
	data := struct {
		Response struct {
			Server ServerDetails
		}
	}{}
	err := s.client.get(context, fmt.Sprintf("server/details/serverid/%s/includestate/yes", serverID), &data)
	return &data.Response.Server, err
}

// Edit modifies a server
func (s *ServerService) Edit(context context.Context, serverID string, params EditServerParams) (*ServerDetails, error) {
	data := struct {
		Response struct {
			Server ServerDetails
		}
	}{}
	err := s.client.post(context, "server/edit", &data, struct {
		EditServerParams
		ServerID string `json:"serverid"`
	}{params, serverID})
	return &data.Response.Server, err
}

// List returns a list of servers
func (s *ServerService) List(context context.Context) (*[]Server, error) {
	data := struct {
		Response struct {
			Servers []Server
		}
	}{}
	err := s.client.get(context, "server/list", &data)
	return &data.Response.Servers, err
}

// Start turns on a server
func (s *ServerService) Start(context context.Context, serverID string) error {
	return s.client.post(context, "server/start", nil, map[string]string{"serverid": serverID})
}

// Stop turns off a server
func (s *ServerService) Stop(context context.Context, serverID string, params StopServerParams) error {
	return s.client.post(context, "server/stop", nil, struct {
		StopServerParams
		ServerID string `json:"serverid"`
	}{params, serverID})
}

func generateHostname() string {
	adjectives := []string{"autumn", "hidden", "bitter", "misty", "silent",
		"empty", "dry", "dark", "summer", "icy", "delicate", "quiet", "white",
		"cool", "spring", "winter", "patient", "twilight", "dawn", "crimson",
		"wispy", "weathered", "blue", "billowing", "broken", "cold", "damp",
		"falling", "frosty", "green", "long", "late", "lingering", "bold", "little",
		"morning", "muddy", "old", "red", "rough", "still", "small", "sparkling",
		"throbbing", "shy", "wandering", "withered", "wild", "black", "young",
		"holy", "solitary", "fragrant", "aged", "snowy", "proud", "floral",
		"restless", "divine", "polished", "ancient", "purple", "lively", "nameless"}
	nouns := []string{"waterfall", "river", "breeze", "moon", "rain", "wind",
		"sea", "morning", "snow", "lake", "sunset", "pine", "shadow", "leaf",
		"dawn", "glitter", "forest", "hill", "cloud", "meadow", "sun", "glade",
		"bird", "brook", "butterfly", "trout", "bush", "dew", "dust", "field",
		"fire", "flower", "firefly", "feather", "grass", "haze", "mountain",
		"night", "pond", "darkness", "snowflake", "silence", "sound", "sky",
		"shape", "surf", "thunder", "violet", "water", "wildflower", "wave",
		"water", "resonance", "sun", "wood", "dream", "cherry", "tree", "fog",
		"frost", "voice", "paper", "frog", "smoke", "star"}

	rand.Seed(time.Now().UTC().UnixNano())

	sections := []string{
		adjectives[rand.Intn(len(adjectives))],
		nouns[rand.Intn(len(nouns))],
		strconv.Itoa(100 + rand.Intn(899)),
	}

	return strings.Join(sections, "-")
}
