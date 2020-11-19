package nginxvts

import "github.com/netdata/go.d.plugin/agent/module"

type (
	// Charts is an alias for module.Charts
	Charts = module.Charts
	Chart  = module.Chart
	// Dims is an alias for module.Dims
	Dims = module.Dims
)

var nginxVtsMainCharts = Charts{
	{
		ID:    "connections",
		Title: "Active Client Connections Including Waiting Connections",
		Units: "connections",
		Fam:   "main",
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

var nginxVtsServerZonesChart = Chart{
	ID:    "%s",
	Title: "Number of responses, later",
	Units: "int64",
	Fam:   "serverzones",
	Ctx:   "nginxvts.servername",
	Dims: Dims{
		{ID: "serverzones_%s_1xx", Name: "HTTP 1xx"},
		{ID: "serverzones_%s_2xx", Name: "HTTP 2xx"},
		{ID: "serverzones_%s_3xx", Name: "HTTP 3xx"},
		{ID: "serverzones_%s_4xx", Name: "HTTP 4xx"},
		{ID: "serverzones_%s_5xx", Name: "HTTP 5xx"},
	},
}
