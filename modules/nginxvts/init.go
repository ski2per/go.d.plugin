package nginxvts

import "github.com/netdata/go.d.plugin/agent/module"

func (nv *Nginxvts) initCharts() (*Charts, error) {
	charts := module.Charts{}
	if err := charts.Add(*nginxVtsMainCharts.Copy()...); err != nil {
		return nil, err
	}

	return &charts, nil
}