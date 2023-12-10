package db

type ip_log struct {
	Timestamp              string
	Hostname               string
	KernelTimestamp        string
	Prefix                 string
	InputInterface         string
	OutputInterface        string
	MACAddress             string
	Source                 string
	Destination            string
	Length                 int
	ToS                    int
	Precedence             int
	TTL                    int
	ID                     string
	CongestionExperienced  string
	DoNotFragment          string
	MoreFragmentsFollowing string
	Frag                   int
	IPOptions              string
	Protocol               string
	Type                   string
	Code                   string
	SourcePort             string
	DestinationPort        string
	Sequence               string
	AckSequence            string
	WindowSize             string
	Res                    string
	Urgent                 string
	Ack                    string
	Push                   string
	Reset                  string
	Syn                    string
	Fin                    string
	Urgp                   string
	TCPOption              string
}

type DashData struct {
	Traf_blocked   int
	Traf_total_in  int
	Traf_total_out int
	Traf_bandwidth int
	Traf_geo       int
	Traf_port      int
	Traf_generic   int
	Traf_tcp       int
	Traf_udp       int
	Traf_icmp      int
}
