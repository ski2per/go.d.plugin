package nginxvts

import "github.com/netdata/go.d.plugin/agent/module"

type (
	// Chart is an alias for module.Chart
	Chart = module.Chart
	// Charts is an alias for module.Charts
	Charts = module.Charts
	// Dim is an alias for module.Dim
	Dim = module.Dim
	// Dims is an alias for module.Dims
	Dims = module.Dims
)

var charts = Charts{
	{
		ID:    "connections",
		Title: "Active Client Connections Including Waiting Connections",
		Units: "connections",
		Fam:   "connections",
		Ctx:   "nginx.connections",
	},
}

var nginxInfoChart = Chart{
	ID:    "nginx_info",
	Title: "Nginx Info",
	Units: "string",
	Fam:   "nginx_vts",
	Ctx:   "vts.info",
}
