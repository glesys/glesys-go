package glesys

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"dario.cat/mergo"
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
	Username   string   `json:"username"`
	PublicKeys []string `json:"sshkeys,omitempty"`
	Password   string   `json:"password,omitempty"`
}

// ServerDetails is a more complete representation of a server
type ServerDetails struct {
	AdditionalDisks []ServerDiskDetails   `json:"additionaldisks,omitempty"`
	CPU             int                   `json:"cpucores"`
	Backup          ServerBackupDetails   `json:"backup,omitempty"`
	Bandwidth       int                   `json:"bandwidth"`
	DataCenter      string                `json:"datacenter"`
	Description     string                `json:"description"`
	Hostname        string                `json:"hostname"`
	ID              string                `json:"serverid"`
	InitialTemplate ServerTemplateDetails `json:"initialtemplate,omitempty"`
	IPList          []ServerIP            `json:"iplist"`
	IsRunning       bool                  `json:"isrunning"`
	IsLocked        bool                  `json:"islocked"`
	ISOFile         string                `json:"isofile,omitempty"`
	Platform        string                `json:"platform"`
	Memory          int                   `json:"memorysize"`
	State           string                `json:"state"`
	Storage         int                   `json:"disksize"`
	Template        string                `json:"templatename"`
}

// ServerBackupDetails represent the backups for a server
type ServerBackupDetails struct {
	Enabled   string                 `json:"enabled"`
	Schedules []ServerBackupSchedule `json:"schedules,omitempty"`
}

// ServerBackupSchedule describes a backup schedule for a KVM server
type ServerBackupSchedule struct {
	Frequency            string `json:"frequency"`
	Numberofimagestokeep int    `json:"numberofimagestokeep"`
}

// ServerConsoleDetails details for connecting to sever web console.
type ServerConsoleDetails struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Password string `json:"password,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	URL      string `json:"url"`
}

// ServerTemplateDetails represents initialtemplate for a KVM server.
type ServerTemplateDetails struct {
	ID          string   `json:"id"`
	CurrentTags []string `json:"currenttags,omitempty"`
	Name        string   `json:"name"`
}

// ServerPlatformTemplates
type ServerPlatformTemplates struct {
	KVM    []ServerPlatformTemplateDetails `json:"KVM"`
	VMware []ServerPlatformTemplateDetails `json:"VMware"`
}

// ServerTemplateInstanceCost
type ServerTemplateInstanceCost struct {
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	Timeperiod string  `json:"timeperiod"`
}

// ServerTemplateLicenseCost
type ServerTemplateLicenseCost struct {
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	Timeperiod string  `json:"timeperiod"`
}

// ServerPlatformTemplateDetails represents a supported template.
type ServerPlatformTemplateDetails struct {
	ID              string                     `json:"id"`
	InstanceCost    ServerTemplateInstanceCost `json:"instancecost"`
	LicenseCost     ServerTemplateLicenseCost  `json:"licensecost"`
	Name            string                     `json:"name"`
	MinDiskSize     int                        `json:"minimumdisksize"`
	MinMemSize      int                        `json:"minimummemorysize"`
	OS              string                     `json:"operatingsystem"`
	Platform        string                     `json:"platform"`
	BootstrapMethod string                     `json:"bootstrapmethod"`
}

// ServerIP is a simple representation of the IP address used in a server.
type ServerIP struct {
	Address string `json:"ipaddress"`
	Version int    `json:"version,omitempty"`
}

// CreateServerParams is used when creating a new server
type CreateServerParams struct {
	Backup            []ServerBackupSchedule `json:"backupschedules,omitempty"`
	Bandwidth         int                    `json:"bandwidth"`
	CampaignCode      string                 `json:"campaigncode,omitempty"`
	CloudConfig       string                 `json:"cloudconfig,omitempty"`
	CloudConfigParams map[string]any         `json:"cloudconfigparams,omitempty"`
	CPU               int                    `json:"cpucores"`
	DataCenter        string                 `json:"datacenter"`
	Description       string                 `json:"description,omitempty"`
	Hostname          string                 `json:"hostname"`
	IPv4              string                 `json:"ip"`
	IPv6              string                 `json:"ipv6"`
	Memory            int                    `json:"memorysize"`
	Password          string                 `json:"rootpassword,omitempty"`
	Platform          string                 `json:"platform"`
	PublicKey         string                 `json:"sshkey,omitempty"`
	Storage           int                    `json:"disksize"`
	Template          string                 `json:"templatename"`
	Users             []User                 `json:"users,omitempty"`
}

// EditServerParams is used when editing an existing server
type EditServerParams struct {
	Backup      []ServerBackupSchedule `json:"backupschedules,omitempty"`
	Bandwidth   int                    `json:"bandwidth,omitempty"`
	CPU         int                    `json:"cpucores,omitempty"`
	Description string                 `json:"description,omitempty"`
	Hostname    string                 `json:"hostname,omitempty"`
	Memory      int                    `json:"memorysize,omitempty"`
	Storage     int                    `json:"disksize,omitempty"`
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
		Platform:   "KVM",
		Storage:    50,
		Template:   "Debian 11 (Bullseye)",
	}
	mergo.Merge(&p, defaults)
	return p
}

// WithUser populates the Users parameter of CreateServerParams for platforms with user support eg. KVM
// Existing parameters will not be overwritten.
func (p CreateServerParams) WithUser(username string, publicKeys []string, password string) CreateServerParams {

	p.Users = append(p.Users, User{
		username,
		publicKeys,
		password,
	})
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

// PreviewCloudConfigParams
type PreviewCloudConfigParams struct {
	CloudConfig       string         `json:"cloudconfig"`
	CloudConfigParams map[string]any `json:"cloudconfigparams,omitempty"`
	Users             []User         `json:"users,omitempty"`
}

// PreviewContext
type PreviewContext struct {
	Params map[string]any `json:"params,omitempty"`
	Users  []User         `json:"users"`
}

// CloudConfigPreview is returned when calling PreviewCloudConfig
type CloudConfigPreview struct {
	Preview string         `json:"preview"`
	Context PreviewContext `json:"context"`
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

// Console returns connection details for server web console
func (s *ServerService) Console(context context.Context, serverID string) (*ServerConsoleDetails, error) {
	data := struct {
		Response struct {
			Console ServerConsoleDetails
		}
	}{}
	err := s.client.post(context, "server/console", &data, struct {
		ServerID string `json:"serverid"`
	}{serverID})
	return &data.Response.Console, err
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

// NetworkAdapters returns a list of NetworkAdapters for `serverID`
func (s *ServerService) NetworkAdapters(context context.Context, serverID string) (*[]NetworkAdapter, error) {
	data := struct {
		Response struct {
			NetworkAdapters []NetworkAdapter
		}
	}{}
	err := s.client.post(context, "server/networkadapters", &data, struct {
		ServerID string `json:"serverid"`
	}{serverID})
	return &data.Response.NetworkAdapters, err
}

// PreviewCloudConfig preview a cloud config mustache template.
func (s *ServerService) PreviewCloudConfig(context context.Context, params PreviewCloudConfigParams) (*CloudConfigPreview, error) {
	data := struct {
		Response struct {
			Cloudconfig CloudConfigPreview
		}
	}{}
	err := s.client.post(context, "server/previewcloudconfig", &data, struct {
		PreviewCloudConfigParams
	}{params})
	return &data.Response.Cloudconfig, err
}

// ListISOs returns a list of ISO files available for `serverID`
func (s *ServerService) ListISOs(context context.Context, serverID string) (*[]string, error) {
	data := struct {
		Response struct {
			IsoFiles []string
		}
	}{}
	err := s.client.post(context, "server/listiso", &data, struct {
		ServerID string `json:"serverid"`
	}{serverID})
	return &data.Response.IsoFiles, err
}

// MountISO mounts the isoFile to the server.
func (s *ServerService) MountISO(context context.Context, serverID string, isoFile string) (*ServerDetails, error) {
	data := struct {
		Response struct {
			Server ServerDetails
		}
	}{}
	err := s.client.post(context, "server/mountiso", &data, struct {
		ServerID string `json:"serverid"`
		ISOFile  string `json:"isofile"`
	}{serverID, isoFile})
	return &data.Response.Server, err
}

// Templates lists all supported templates per platform
func (s *ServerService) Templates(context context.Context) (*ServerPlatformTemplates, error) {
	data := struct {
		Response struct {
			Templates ServerPlatformTemplates
		}
	}{}
	err := s.client.post(context, "server/templates", &data, nil)
	return &data.Response.Templates, err
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

	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	sections := []string{
		adjectives[r.Intn(len(adjectives))],
		nouns[r.Intn(len(nouns))],
		strconv.Itoa(100 + r.Intn(899)),
	}

	return strings.Join(sections, "-")
}
