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

	// nv.addMainCharts(ms, collected)
	// nv.addServerZonesCharts(ms, collected)
	// nv.addUpstreamZonesCharts(ms, collected)
	nv.addFilterZonesCharts(ms, collected)

	// fmt.Printf("\n\n\n%+v\n\n", collected)
	// tmp := stm.ToMap(ms)
	tmp := stm.ToMap(collected)

	// return stm.ToMap()
	// fmt.Printf("\n\n\n%+v\n\n\n", tmp)
	return tmp, nil
}

func (nv *NginxVts) addMainCharts(stat *vtsStatus, collected map[string]interface{}) {
	collected["loadmsec"] = stat.LoadMsec
	collected["nowmsec"] = stat.NowMsec
	collected["connections"] = stat.Connections
}

func (nv *NginxVts) addServerZonesCharts(stat *vtsStatus, collected map[string]interface{}) {
	if !stat.hasServerZones() {
		return
	}
	collected["serverzones"] = stat.ServerZones

	for server := range stat.ServerZones {
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

func (nv *NginxVts) addUpstreamZonesCharts(stat *vtsStatus, collected map[string]interface{}) {
	if !stat.hasUpstreamZones() {
		return
	}

	upstreamMap := make(map[string]Upstream)

	for upstreamKey, upstreamList := range stat.UpstreamZones {
		for _, upstream := range upstreamList {
			mergedKey := fmt.Sprintf("%s:%s", upstreamKey, upstream.Server)
			upstreamMap[mergedKey] = upstream

			charts := nginxVtsUpstreamZonesCharts.Copy()
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

func (nv *NginxVts) addFilterZonesCharts(stat *vtsStatus, collected map[string]interface{}) {
	if !stat.hasFilterZones() {
		return
	}

	filterMap := make(map[string]Server)

	for filter, serverMap := range stat.FilterZones {
		for group, upstream := range serverMap {
			mergedKey := fmt.Sprintf("%s%s", filter, group)
			filterMap[mergedKey] = upstream

			charts := nginxVtsFilterZonesCharts.Copy()
			for _, chart := range *charts {
				chart.ID = fmt.Sprintf(chart.ID, mergedKey)
				chart.Fam = filter
				for _, dim := range chart.Dims {
					dim.ID = fmt.Sprintf(dim.ID, mergedKey)
				}
			}
			_ = nv.charts.Add(*charts...)
		}
	}
	collected["filterzones"] = filterMap
}
