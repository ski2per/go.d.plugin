package nginxvts

import (
	"net/http"

	"github.com/netdata/go.d.plugin/pkg/web"
)

type apiClient struct {
	httpClient *http.Client
	request    web.Request
}

func newAPIClient(client *http.Client, request web.Request) *apiClient {
	return &apiClient{
		httpClient: client,
		request:    request,
	}
}
