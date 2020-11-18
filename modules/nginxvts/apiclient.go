package nginxvts

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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

func (api *apiClient) getVtsStatus() (VTSData, error) {
	req, err := web.NewHTTPRequest(api.request)
	if err != nil {
		return VTSData{}, fmt.Errorf("error on creating request : %v", err)
	}
	resp, err := api.doRequestOK(req)
	defer closeBody(resp)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll failed", err)
	}

	var vtsData VTSData
	err = json.Unmarshal(data, &vtsData)
	if err != nil {
		log.Println("json.Unmarshal failed", err)
	}
	return vtsData, nil

}

func (api *apiClient) doRequestOK(req *http.Request) (*http.Response, error) {
	resp, err := api.httpClient.Do(req)

	if err != nil {
		return resp, fmt.Errorf("error on request : %v", err)
	}

	if resp.StatusCode != 200 {
		return resp, fmt.Errorf("%s returned HTTP status %d", req.URL, resp.StatusCode)
	}

	return resp, err
}

func closeBody(resp *http.Response) {
	if resp != nil && resp.Body != nil {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}
}
