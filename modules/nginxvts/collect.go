package nginxvts

import (
	"errors"
	"fmt"
)

func (nv *NginxVTS) collect() (map[string]int64, error) {
	fmt.Println("hello nginxvts")
	return make(map[string]int64), errors.New("mockup")
}
