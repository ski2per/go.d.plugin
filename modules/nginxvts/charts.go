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
		ID:    "times",
		Title: "Nginx running times",
		Units: "milliseconds",
		Fam:   "main",
		Ctx:   "nginxvts.main",
		Dims: Dims{
			{ID: "loadmsec", Name: "Starting time"},
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
			{ID: "connections_active", Name: "Active"},
			{ID: "connections_reading", Name: "Reading"},
			{ID: "connections_writing", Name: "Writing"},
			{ID: "connections_waiting", Name: "Waiting"},
			{ID: "connections_accepted", Name: "Accepted"},
			{ID: "connections_handled", Name: "Handled"},
			{ID: "connections_requests", Name: "Requests"},
		},
	},
}

var nginxVtsSharedZonesChart = Charts{
	{
		ID:    "size",
		Title: "Shared memory size",
		Units: "bytes",
		Fam:   "sharedzones",
		Ctx:   "nginxvts.sharedzones.size",
		Dims: Dims{
			{ID: "sharedzones_maxsize", Name: "Max size"},
			{ID: "sharedzones_usedsize", Name: "Used size"},
		},
	},
	{
		ID:    "node",
		Title: "Number of node using shared memory",
		Units: "number",
		Fam:   "sharedzones",
		Ctx:   "nginxvts.sharedzones.node",
		Dims: Dims{
			{ID: "sharedzones_usednode", Name: "Used node"},
		},
	},
}

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
		Title: "Response code of Serverzones",
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
		ID:    "io_%s",
		Title: "ServerZones IO",
		Units: "bytes",
		Fam:   "serverzones",
		Ctx:   "nginxvts.serverzones.io",
		Dims: Dims{
			{ID: "serverzones_%s_inbytes", Name: "Inbytes"},
			{ID: "serverzones_%s_outbytes", Name: "Outbytes"},
		},
	},
	{
		ID:    "cache_%s",
		Title: "Serverzones cache",
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
		Title: "Response code of Upstreamzones",
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
		Title: "Upstreamzones IO",
		Units: "bytes",
		Fam:   "upstreamzones",
		Ctx:   "nginxvts.upstreamzones.io",
		Dims: Dims{
			{ID: "upstreamzones_%s_inbytes", Name: "inbytes"},
			{ID: "upstreamzones_%s_outbytes", Name: "outbytes"},
		},
	},
}

var nginxVtsFilterZonesCharts = Charts{
	{
		ID:    "requests_%s",
		Title: "Total number of client requests",
		Units: "number",
		Fam:   "filterzones",
		Ctx:   "nginxvts.filterzones.requests",
		Dims: Dims{
			{ID: "filterzones_%s_requestcounter", Name: "Requests"},
		},
	},
	{
		ID:    "responses_%s",
		Title: "Response code of Filterzones",
		Units: "number",
		Fam:   "filterzones",
		Ctx:   "nginxvts.filterzones.responses",
		Dims: Dims{
			{ID: "filterzones_%s_responses_1xx", Name: "1xx"},
			{ID: "filterzones_%s_responses_2xx", Name: "2xx"},
			{ID: "filterzones_%s_responses_3xx", Name: "3xx"},
			{ID: "filterzones_%s_responses_4xx", Name: "4xx"},
			{ID: "filterzones_%s_responses_5xx", Name: "5xx"},
		},
	},
	{
		ID:    "io_%s",
		Title: "Filterzones IO",
		Units: "bytes",
		Fam:   "filterzones",
		Ctx:   "nginxvts.filterzones.io",
		Dims: Dims{
			{ID: "filterzones_%s_inbytes", Name: "inbytes"},
			{ID: "filterzones_%s_outbytes", Name: "outbytes"},
		},
	},
	{
		ID:    "cache_%s",
		Title: "Filterzones cache",
		Units: "number",
		Fam:   "filterzones",
		Ctx:   "nginxvts.filterzones.cache",
		Dims: Dims{
			{ID: "filterzones_%s_responses_miss", Name: "Miss"},
			{ID: "filterzones_%s_responses_bypass", Name: "Bypass"},
			{ID: "filterzones_%s_responses_expired", Name: "Expired"},
			{ID: "filterzones_%s_responses_stale", Name: "Stale"},
			{ID: "filterzones_%s_responses_updating", Name: "Updating"},
			{ID: "filterzones_%s_responses_revalidated", Name: "revalidated"},
			{ID: "filterzones_%s_responses_hit", Name: "Hit"},
			{ID: "filterzones_%s_responses_scarce", Name: "Scarce"},
		},
	},
}

var nginxVtsCacheZonesCharts = Charts{
	{
		ID:    "size_%s",
		Title: "Cache size",
		Units: "bytes",
		Fam:   "cachezones",
		Ctx:   "nginxvts.cachezones.requests",
		Dims: Dims{
			{ID: "cachezones_%s_maxsize", Name: "Max size"},
			{ID: "cachezones_%s_usedsize", Name: "Used size"},
		},
	},
	{
		ID:    "io_%s",
		Title: "Cache IO",
		Units: "bytes",
		Fam:   "cachezones",
		Ctx:   "nginxvts.cachezones.io",
		Dims: Dims{
			{ID: "cachezones_%s_inbytes", Name: "inbytes"},
			{ID: "cachezones_%s_outbytes", Name: "outbytes"},
		},
	},
	{
		ID:    "responses_%s",
		Title: "Cache responses",
		Units: "number",
		Fam:   "cachezones",
		Ctx:   "nginxvts.cachezones.cache",
		Dims: Dims{
			{ID: "cachezones_%s_responses_miss", Name: "Miss"},
			{ID: "cachezones_%s_responses_bypass", Name: "Bypass"},
			{ID: "cachezones_%s_responses_expired", Name: "Expired"},
			{ID: "cachezones_%s_responses_stale", Name: "Stale"},
			{ID: "cachezones_%s_responses_updating", Name: "Updating"},
			{ID: "cachezones_%s_responses_revalidated", Name: "revalidated"},
			{ID: "cachezones_%s_responses_hit", Name: "Hit"},
			{ID: "cachezones_%s_responses_scarce", Name: "Scarce"},
		},
	},
}
