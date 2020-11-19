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

var nginxVtsServerZonesCharts = Charts{
	{
		ID:    "responses_%s",
		Title: "Response code status.",
		Units: "number",
		Fam:   "serverzones",
		Ctx:   "response_code",
		Dims: Dims{
			{ID: "serverzones_%s_responses_1xx", Name: "1xx"},
			{ID: "serverzones_%s_responses_2xx", Name: "2xx"},
			{ID: "serverzones_%s_responses_3xx", Name: "3xx"},
			{ID: "serverzones_%s_responses_4xx", Name: "4xx"},
			{ID: "serverzones_%s_responses_5xx", Name: "5xx"},
		},
	},
	{
		ID:    "status_%s",
		Title: "ServerZone status.",
		Units: "status",
		Fam:   "serverzones",
		Ctx:   "serverzone_status",
		Dims: Dims{
			{ID: "serverzones_%s_requestcounter", Name: "requestcounter"},
			{ID: "serverzones_%s_inbytes", Name: "inbytes"},
			{ID: "serverzones_%s_outbytes", Name: "outbytes"},
		},
	},
}
