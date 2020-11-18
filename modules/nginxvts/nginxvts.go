package nginxvts

import (
	"fmt"
	"time"

	"github.com/netdata/go.d.plugin/pkg/web"

	"github.com/netdata/go.d.plugin/agent/module"
)

const (
	// defaultURL         = "http://127.0.0.1/status/format/json"
	defaultURL         = "http://10.0.0.110:8888/status/format/json"
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
	// charts    *module.Charts
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
	fmt.Println("++++++++++++++++++++++ Init()")
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

	//Init charts
	// charts, err := nv.initCharts()
	if err != nil {
		nv.Errorf("charts init: %v", err)
		return false
	}
	// nv.charts = charts

	return true
}

// Check makes check.
func (nv *NginxVTS) Check() bool {
	fmt.Println("++++++++++++++++++++++ Check()")
	fmt.Println(len(nv.Collect()))
	return len(nv.Collect()) > 0
}

// Charts creates Charts.
func (nv *NginxVTS) Charts() *Charts {
	fmt.Println("++++++++++++++++++++++ Charts()")
	return nginxVtsMainCharts.Copy()
	// return nv.charts
}

// Collect collects metrics.
func (nv *NginxVTS) Collect() map[string]int64 {
	fmt.Println("++++++++++++++++++++++ Collect()")
	mx, err := nv.collect()
	fmt.Println(mx)

	if err != nil {
		nv.Error(err)
		return nil
	}

	return mx
}
