package binding

// port request struct for incoming form data
type PortReq struct {
	Port   string `form:"Port"`
	Target string `form:"Target"`
}

// port transfer struct for security
type PortDTO struct {
	Port   string
	Target string
}

// firewall request struct for incoming form data
type FirewallRequest struct {
	Table       string `form:"IP_Table" valid:"Table"`
	Protocol    string `form:"Protocol" valid:"Protocol"`
	Source      string `form:"Source" valid:"Source,optional"`
	Destination string `form:"Destination" valid:"Destination,optional"`
	Port        string `form:"Port_Blocking" valid:"Port_Blocking,optional"`
	Target      string `form:"Target" valid:"Target"`
}

// firewall transfer struct for security
type FirewallRequestDTO struct {
	Table       string
	Protocol    string
	Source      string
	Destination string
	Port        string
	Target      string
}

// geo request struct for incoming form data
type Geo struct {
	Location string `form:"Location" valid:"Location"`
}

// geo transfer struct for security
type GeoDTO struct {
	Location string
}

// ip rate request struct for incoming form data
type IPRate struct {
	Source    string `form:"Source" valid:"Source"`
	Conn      string `form:"Connections" valid:"Connections"`
	Protocol  string `form:"Protocol" valid:"Protocol"`
	TimeFrame string `form:"TimeFrame" valid:"TimeFrame"`
}

// ip rate transfer struct for security
type IPRateDTO struct {
	Source    string
	Conn      string
	Protocol  string
	TimeFrame string
}

// bandwidth request struct for incoming form data
type Bandwidth struct {
	Protocol string `form:"Protocol" valid:"Protocol"`
	Source   string `form:"Source" valid:"Source"`
	Time     string `form:"Time" valid:"Time"`
	Type     string `form:"Type" valid:"Type"`
	Bamount  string `form:"Bamount" valid:"Bamount"`
}

// bandwidth transfer struct for security
type BandwidthDTO struct {
	Protocol string
	Source   string
	Time     string
	Type     string
	Bamount  string
}
