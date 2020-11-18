package nginxvts

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/netdata/go.d.plugin/pkg/stm"
)

func (nv *NginxVTS) collect() (map[string]int64, error) {
	// collected := make(map[string]int64)

	data, err := nv.apiClient.getVtsStatus()

	var vtsData VTSData
	err = json.Unmarshal(data, &vtsData)
	if err != nil {
		log.Println("json.Unmarshal failed", err)
	}

	fmt.Println(vtsData.Connections)

	// tmp := stm.ToMap(vtsData.Connections)
	// fmt.Printf("%+v\n", tmp)
	// for k, v := range tmp {
	// 	fmt.Println(k, v)
	// }

	// for _, chart := range *nv.charts {
	// 	// collected[chart.ID] = stm.ToMap(vtsData.Connections)
	// 	chart.MarkNotCreated()
	// 	return stm.ToMap(vtsData.Connections), nil
	// }
	fmt.Println(nv.charts)
	fmt.Println(len(*nv.charts))

	return stm.ToMap(vtsData.Connections), nil
}
