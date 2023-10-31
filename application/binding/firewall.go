package bindings

type FirewallRequest struct {
	Table string `form:"IP_Table"`
	Protocol string `form:"Protocol"`
	Source string `form:"Source"`
	Destination string `form:"Destination"`
	Port string `form:"Port_Blocking"`
	Target string `form:"Target"`
}

type FirewallRequestDTO struct {
	Table string
	Protocol string
	Source string
	Destination string
	Port string
	Target string
}