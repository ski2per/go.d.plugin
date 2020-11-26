package nginxvts

import (
	"fmt"

	"github.com/netdata/go.d.plugin/pkg/stm"
)

func (nv *NginxVts) collect() (map[string]int64, error) {
	collected := make(map[string]interface{})
	ms, err := nv.apiClient.getVtsStatus()
	// fmt.Printf("%+v\n", ms)
	if err != nil {
		// return make(map[string]int64), errors.New("get vts status error")
		// return nil, errors.New("get vts status error")
		return nil, err
	}

	nv.addMainCharts(ms, collected)
	nv.addServerZonesCharts(ms, collected)
	nv.addUpstreamZonesCharts(ms, collected)

	// fmt.Printf("%+v\n\n", collected)
	// tmp := stm.ToMap(ms)
	tmp := stm.ToMap(collected)

	// return stm.ToMap()
	// fmt.Printf("\n\n\n%+v\n\n\n", tmp)
	return tmp, nil
}

func (nv *NginxVts) addMainCharts(ms *vtsStatus, collected map[string]interface{}) {
	collected["loadmsec"] = ms.LoadMsec
	collected["nowmsec"] = ms.NowMsec
	collected["connections"] = ms.Connections
}

func (nv *NginxVts) addServerZonesCharts(ms *vtsStatus, collected map[string]interface{}) {
	if !ms.hasServerZones() {
		return
	}
	collected["serverzones"] = ms.ServerZones

	for server := range ms.ServerZones {
		charts := nginxVtsServerZonesCharts.Copy()
		for _, chart := range *charts {
			chart.ID = fmt.Sprintf(chart.ID, server)
			chart.Fam = server
			for _, dim := range chart.Dims {
				dim.ID = fmt.Sprintf(dim.ID, server)
			}
		}

		// fmt.Printf("\n\n\n%+v\n\n\n", *charts)
		_ = nv.charts.Add(*charts...)
	}
}

func (nv *NginxVts) addUpstreamZonesCharts(ms *vtsStatus, collected map[string]interface{}) {
	if !ms.hasUpstreamZones() {
		return
	}

	upstreamMap := make(map[string]Upstream)
	for upstreamKey, upstreamList := range ms.UpstreamZones {
		for _, upstream := range upstreamList {
			charts := nginxVtsUpstreamZonesCharts.Copy()
			mergedKey := fmt.Sprintf("%s:%s", upstreamKey, upstream.Server)

			upstreamMap[mergedKey] = upstream
			for _, chart := range *charts {
				chart.ID = fmt.Sprintf(chart.ID, mergedKey)
				chart.Fam = upstream.Server
				for _, dim := range chart.Dims {
					dim.ID = fmt.Sprintf(dim.ID, mergedKey)
				}
			}
			_ = nv.charts.Add(*charts...)
		}
	}
	collected["upstreamzones"] = upstreamMap
}
