package nginxvts

import "github.com/netdata/go.d.plugin/agent/module"

type (
	// Charts is an alias for module.Charts
	Charts = module.Charts
	// Dims is an alias for module.Dims
	Dims = module.Dims
)

var nginxVtsMainCharts = Charts{
	{
		ID:    "vtsconn",
		Title: "Active Client Connections Including Waiting Connections",
		Units: "connections",
		Fam:   "connections",
		Ctx:   "nginxvts.connections",
		Dims: Dims{
			{ID: "connections_active", Name: "Active connections"},
			{ID: "connections_reading", Name: "Reading connections"},
			{ID: "connections_writing", Name: "Writing connections"},
			{ID: "connections_waiting", Name: "Waiting connections"},
			{ID: "connections_accepted", Name: "Accepted connections"},
			{ID: "connections_handled", Name: "Handled connections"},
			{ID: "connections_requests", Name: "Requests connections"},
		},
	},
}
