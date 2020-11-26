package nginxvts

import "github.com/netdata/go.d.plugin/agent/module"

type (
	// Charts is an alias for module.Charts
	Charts = module.Charts
	// Chart is an alias for module.Chart
	Chart = module.Chart
	// Dims is an alias for module.Dims
	Dims = module.Dims
)

var nginxVtsMainCharts = Charts{
	{
		ID:    "times",
		Title: "Nginx running times",
		Units: "milliseconds",
		Fam:   "main",
		Ctx:   "nginxvts.main",
		Dims: Dims{
			{ID: "loadmsec", Name: "Start time"},
			{ID: "nowmsec", Name: "Up time"},
		},
	},
	{
		ID:    "connections",
		Title: "Nginx Connections",
		Units: "number",
		Fam:   "main",
		Ctx:   "nginxvts.main",
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

// var nginxVtsSharedZonesChart = Charts{
// }

var nginxVtsServerZonesCharts = Charts{
	{
		ID:    "requests_%s",
		Title: "Total number of client requests",
		Units: "number",
		Fam:   "serverzones",
		Ctx:   "nginxvts.serverzones.requests",
		Dims: Dims{
			{ID: "serverzones_%s_requestcounter", Name: "Requests"},
		},
	},
	{
		ID:    "responses_%s",
		Title: "ServerZones response code",
		Units: "number",
		Fam:   "serverzones",
		Ctx:   "nginxvts.serverzones.responses",
		Dims: Dims{
			{ID: "serverzones_%s_responses_1xx", Name: "1xx"},
			{ID: "serverzones_%s_responses_2xx", Name: "2xx"},
			{ID: "serverzones_%s_responses_3xx", Name: "3xx"},
			{ID: "serverzones_%s_responses_4xx", Name: "4xx"},
			{ID: "serverzones_%s_responses_5xx", Name: "5xx"},
		},
	},
	{
		ID:    "cache_%s",
		Title: "ServerZones cache",
		Units: "number",
		Fam:   "serverzones",
		Ctx:   "nginxvts.serverzones.cache",
		Dims: Dims{
			{ID: "serverzones_%s_responses_miss", Name: "Miss"},
			{ID: "serverzones_%s_responses_bypass", Name: "Bypass"},
			{ID: "serverzones_%s_responses_expired", Name: "Expired"},
			{ID: "serverzones_%s_responses_stale", Name: "Stale"},
			{ID: "serverzones_%s_responses_updating", Name: "Updating"},
			{ID: "serverzones_%s_responses_revalidated", Name: "revalidated"},
			{ID: "serverzones_%s_responses_hit", Name: "Hit"},
			{ID: "serverzones_%s_responses_scarce", Name: "Scarce"},
		},
	},
	{
		ID:    "io_%s",
		Title: "ServerZones IO",
		Units: "bytes",
		Fam:   "serverzones",
		Ctx:   "nginxvts.serverzones.io",
		Dims: Dims{
			{ID: "serverzones_%s_inbytes", Name: "inbytes"},
			{ID: "serverzones_%s_outbytes", Name: "outbytes"},
		},
	},
}

var nginxVtsUpstreamZonesCharts = Charts{
	{
		ID:    "requests_%s",
		Title: "Total number of client connections forwarded to this server",
		Units: "number",
		Fam:   "upstreamzones",
		Ctx:   "nginxvts.upstreamzones.requests",
		Dims: Dims{
			{ID: "upstreamzones_%s_requestcounter", Name: "Requests"},
		},
	},
	{
		ID:    "responses_%s",
		Title: "Upstreamzones response code",
		Units: "number",
		Fam:   "upstreamzones",
		Ctx:   "nginxvts.upstreamzones.responses",
		Dims: Dims{
			{ID: "upstreamzones_%s_responses_1xx", Name: "1xx"},
			{ID: "upstreamzones_%s_responses_2xx", Name: "2xx"},
			{ID: "upstreamzones_%s_responses_3xx", Name: "3xx"},
			{ID: "upstreamzones_%s_responses_4xx", Name: "4xx"},
			{ID: "upstreamzones_%s_responses_5xx", Name: "5xx"},
		},
	},
	{
		ID:    "io_%s",
		Title: "UpstreamZones IO",
		Units: "bytes",
		Fam:   "upstreamzones",
		Ctx:   "nginxvts.upstreamzones.io",
		Dims: Dims{
			{ID: "upstreamzones_%s_inbytes", Name: "inbytes"},
			{ID: "upstreamzones_%s_outbytes", Name: "outbytes"},
		},
	},
}
