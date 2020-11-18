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

	fmt.Printf("%+v\n\n", vtsData)
	tmp := stm.ToMap(vtsData)
	fmt.Println(tmp)
	// return stm.ToMap(vtsData), nil
	return tmp, nil
}
