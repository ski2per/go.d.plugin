package nginxvts

import (
	"errors"

	"github.com/netdata/go.d.plugin/pkg/stm"
)

func (nv *NginxVTS) collect() (map[string]int64, error) {
	// collected := make(map[string]int64)

	data, err := nv.apiClient.getVtsStatus()
	if err != nil {
		return make(map[string]int64), errors.New("get vts status error")
	}

	// fmt.Printf("%+v\n\n", data)
	tmp := stm.ToMap(data)
	// fmt.Println(tmp)
	// return stm.ToMap(vtsData), nil
	return tmp, nil
}
