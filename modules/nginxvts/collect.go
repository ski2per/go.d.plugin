package nginxvts

import (
	"errors"
	"fmt"

	"github.com/netdata/go.d.plugin/pkg/stm"
)

func (nv *Nginxvts) collect() (map[string]int64, error) {
	ms, err := nv.apiClient.getVtsStatus()
	// fmt.Printf("%+v\n", ms)
	if err != nil {
		return make(map[string]int64), errors.New("get vts status error")
	}

	// collected := make(map[string]int64)
	// nv.collectServerZonesStatus(collected, ms)

	nv.addServerZonesCharts(ms)

	// fmt.Printf("%+v\n\n", ms)
	tmp := stm.ToMap(ms)

	// return stm.ToMap()
	// fmt.Printf("\n\n\n%+v\n\n\n", tmp)
	return tmp, nil
}

// func (Nginxvts) collectMainStatus(collected map[string]int64, ms *vtsStatus) { }

// func (Nginxvts) collectServerZonesStatus(collected map[string]int64, ms *vtsStatus) { }

func (nv *Nginxvts) addServerZonesCharts(ms *vtsStatus) {
	if !ms.hasServerZones() {
		return
	}

	for server := range ms.ServerZones {
		charts := nginxVtsServerZonesCharts.Copy()
		for _, chart := range *charts {
			chart.ID = fmt.Sprintf(chart.ID, replaceString(server))
			for _, dim := range chart.Dims {
				dim.ID = fmt.Sprintf(dim.ID, replaceString(server))
				// fmt.Println(dim.ID)
			}
		}

		// fmt.Printf("\n\n\n%+v\n\n\n", *charts)
		_ = nv.charts.Add(*charts...)
	}
}

func replaceString(id string) string {
	switch id {
	case "*":
		return "all"
	case "_":
		return "default"
	default:
		return id
	}
}
