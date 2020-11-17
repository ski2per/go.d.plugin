package nginxvts

import (
	"errors"
	"fmt"
)

func (nv *NginxVTS) collect() (map[string]int64, error) {
	fmt.Println("hello nginxvts")
	resp, err := nv.apiClient.getVtsStatus()
	fmt.Println(resp.Body)
	if err != nil {
		fmt.Println("Shit error")
	} else {
		fmt.Println(resp)
	}

	return make(map[string]int64), errors.New("mockup")
}
