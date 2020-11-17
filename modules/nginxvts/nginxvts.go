package nginxvts

import (
	"time"

	"github.com/netdata/go.d.plugin/pkg/web"

	"github.com/netdata/go.d.plugin/agent/module"
)

const (
	defaultURL         = "http://127.0.0.1/status"
	defaultHTTPTimeout = time.Second
)

// Config is the Nginx module configuration.
type Config struct {
	web.HTTP `yaml:",inline"`
}

// Nginx nginx module.
type NginxVTS struct {
	module.Base
	Config `yaml:",inline"`

	apiClient *apiClient
}

// New creates Nginx with default values.
func New() *NginxVTS {
	config := Config{
		HTTP: web.HTTP{
			Request: web.Request{
				URL: defaultURL,
			},
			Client: web.Client{
				Timeout: web.Duration{Duration: defaultHTTPTimeout},
			},
		},
	}

	return &NginxVTS{Config: config}
}

func init()

// Cleanup makes cleanup.
func (NginxVTS) Cleanup() {}

// Init makes initialization.
func (nvts *NginxVTS) Init() bool {
	if nvts.URL == "" {
		nvts.Error("URL not set")
		return false
	}

	// client, err := web.NewHTTPClient(n.Client)
	if err != nil {
		nvts.Error(err)
		return false
	}

	// n.apiClient = newAPIClient(client, n.Request)

	// n.Debugf("using URL %s", n.URL)
	// n.Debugf("using timeout: %s", n.Timeout.Duration)

	return true
}

// Check makes check.
func (nvts *NginxVTS) Check() bool { return len(nvts.Collect()) > 0 }

// Charts creates Charts.
func (NginxVTS) Charts() *Charts { return charts.Copy() }

// Collect collects metrics.
func (nvts *NginxVTS) Collect() map[string]int64 {
	mx, err := nvts.collect()

	if err != nil {
		nvts.Error(err)
		return nil
	}

	return mx
}
