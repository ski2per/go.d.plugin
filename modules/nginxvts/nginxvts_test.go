package nginxvts

import (
	"testing"

	"github.com/netdata/go.d.plugin/agent/module"
	"github.com/netdata/go.d.plugin/pkg/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Cleanup(t *testing.T) { New().Cleanup() }

func Test_New(t *testing.T) {
	job := New()

	assert.Implements(t, (*module.Module)(nil), job)
	assert.Equal(t, defaultURL, job.URL)
	assert.Equal(t, defaultHTTPTimeout, job.Timeout.Duration)
}

func Test_Init(t *testing.T) {
	job0 := New()
	job1 := &NginxVts{
		Config: Config{
			HTTP: web.HTTP{
				Request: web.Request{
					URL: "",
				},
			},
		},
	}

	require.True(t, job0.Init())
	require.False(t, job1.Init())
}
