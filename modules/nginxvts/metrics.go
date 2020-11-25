package nginxvts

type vtsStatus struct {
	HostName     string //`json:"hostName"`
	NginxVersion string //`json:"nginxVersion"`
	LoadMsec     int64  //`json:"loadMsec"`
	NowMsec      int64  //`json:"nowMsec"`
	Connections  struct {
		Active   int64 `stm:"active"`
		Reading  int64 `stm:"reading"`
		Writing  int64 `stm:"writing"`
		Waiting  int64 `stm:"waiting"`
		Accepted int64 `stm:"accepted"`
		Handled  int64 `stm:"handled"`
		Request  int64 `stm:"requests"`
	} `stm:"connections"`
	SharedZones struct {
		Name     string
		MaxSize  int64 `stm:"maxsize"`
		UsedSize int64 `stm:"usedsize"`
		UsedNode int64 `stm:"usednode"`
	} `stm:"sharedzones" json:"sharedZones"`
	ServerZones   map[string]Server              `stm:"serverzones" json:"serverZones"`
	UpstreamZones map[string][]Upstream          `json:"upstreamZones"`
	FilterZones   map[string]map[string]Upstream `json:"filterZones"`
	CacheZones    map[string]Cache               `json:"cacheZones"`
}

func (v vtsStatus) hasServerZones() bool   { return v.ServerZones != nil }
func (v vtsStatus) hasUpstreamZones() bool { return v.UpstreamZones != nil }
func (v vtsStatus) hasFilterZones() bool   { return v.FilterZones != nil }
func (v vtsStatus) hasCacheZones() bool    { return v.CacheZones != nil }

type Server struct {
	RequestCounter int64 `stm:"requestcounter"`
	InBytes        int64 `stm:"inbytes"`
	OutBytes       int64 `stm:"outbytes"`
	RequestMsec    int64
	Responses      struct {
		OneXx       int64 `stm:"1xx" json:"1xx"`
		TwoXx       int64 `stm:"2xx" json:"2xx"`
		ThreeXx     int64 `stm:"3xx" json:"3xx"`
		FourXx      int64 `stm:"4xx" json:"4xx"`
		FiveXx      int64 `stm:"5xx" json:"5xx"`
		Miss        uint64
		Bypass      uint64
		Expired     uint64
		Stale       uint64
		Updating    uint64
		Revalidated uint64
		Hit         uint64
		Scarce      uint64
	} `stm:"responses" json:"responses"`
	OverCounts struct {
		MaxIntegerSize float64
		RequestCounter uint64
		InBytes        uint64
		OutBytes       uint64
		OneXx          uint64
		TwoXx          uint64
		ThreeXx        uint64
		FourXx         uint64
		FiveXx         uint64
		Miss           uint64
		Bypass         uint64
		Expired        uint64
		Stale          uint64
		Updating       uint64
		Revalidated    uint64
		Hit            uint64
		Scarce         uint64
	}
}

type Upstream struct {
	Server         string `json:"server"`
	RequestCounter uint64 `json:"requestCounter"`
	InBytes        uint64 `json:"inBytes"`
	OutBytes       uint64 `json:"outBytes"`
	Responses      struct {
		OneXx   uint64 `json:"1xx"`
		TwoXx   uint64 `json:"2xx"`
		ThreeXx uint64 `json:"3xx"`
		FourXx  uint64 `json:"4xx"`
		FiveXx  uint64 `json:"5xx"`
	} `json:"responses"`
	ResponseMsec uint64 `json:"responseMsec"`
	RequestMsec  uint64 `json:"requestMsec"`
	Weight       uint64 `json:"weight"`
	MaxFails     uint64 `json:"maxFails"`
	FailTimeout  uint64 `json:"failTimeout"`
	Backup       bool   `json:"backup"`
	Down         bool   `json:"down"`
	OverCounts   struct {
		MaxIntegerSize float64 `json:"maxIntegerSize"`
		RequestCounter uint64  `json:"requestCounter"`
		InBytes        uint64  `json:"inBytes"`
		OutBytes       uint64  `json:"outBytes"`
		OneXx          uint64  `json:"1xx"`
		TwoXx          uint64  `json:"2xx"`
		ThreeXx        uint64  `json:"3xx"`
		FourXx         uint64  `json:"4xx"`
		FiveXx         uint64  `json:"5xx"`
	} `json:"overCounts"`
}

type Cache struct {
	MaxSize   uint64 `json:"maxSize"`
	UsedSize  uint64 `json:"usedSize"`
	InBytes   uint64 `json:"inBytes"`
	OutBytes  uint64 `json:"outBytes"`
	Responses struct {
		Miss        uint64 `json:"miss"`
		Bypass      uint64 `json:"bypass"`
		Expired     uint64 `json:"expired"`
		Stale       uint64 `json:"stale"`
		Updating    uint64 `json:"updating"`
		Revalidated uint64 `json:"revalidated"`
		Hit         uint64 `json:"hit"`
		Scarce      uint64 `json:"scarce"`
	} `json:"responses"`
	OverCounts struct {
		MaxIntegerSize float64 `json:"maxIntegerSize"`
		InBytes        uint64  `json:"inBytes"`
		OutBytes       uint64  `json:"outBytes"`
		Miss           uint64  `json:"miss"`
		Bypass         uint64  `json:"bypass"`
		Expired        uint64  `json:"expired"`
		Stale          uint64  `json:"stale"`
		Updating       uint64  `json:"updating"`
		Revalidated    uint64  `json:"revalidated"`
		Hit            uint64  `json:"hit"`
		Scarce         uint64  `json:"scarce"`
	} `json:"overCounts"`
}
