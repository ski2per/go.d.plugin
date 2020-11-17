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

func init() {
	creator := module.Creator{
		Create: func() module.Module { return New() },
	}

	module.Register("nginxvts", creator)
}

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

// Cleanup makes cleanup.
func (NginxVTS) Cleanup() {}

// Init makes initialization.
func (nv *NginxVTS) Init() bool {
	if nv.URL == "" {
		nv.Error("URL not set")
		return false
	}

	client, err := web.NewHTTPClient(nv.Client)
	if err != nil {
		nv.Error(err)
		return false
	}

	nv.apiClient = newAPIClient(client, nv.Request)

	nv.Debugf("using URL %s", nv.URL)
	nv.Debugf("using timeout: %s", nv.Timeout.Duration)

	return true
}

// Check makes check.
func (nv *NginxVTS) Check() bool { return len(nv.Collect()) > 0 }

// Charts creates Charts.
func (NginxVTS) Charts() *Charts { return charts.Copy() }

// Collect collects metrics.
func (nv *NginxVTS) Collect() map[string]int64 {
	mx, err := nv.collect()

	if err != nil {
		nv.Error(err)
		return nil
	}

	return mx
}
