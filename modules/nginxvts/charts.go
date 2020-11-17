package nginxvts

import "github.com/netdata/go.d.plugin/agent/module"

type (
	// Charts is an alias for module.Charts
	Charts = module.Charts
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
