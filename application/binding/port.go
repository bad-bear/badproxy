package bindings

type PortReq {
	Port string `form:"Port"`
	Target string `form:"Target"`
}

type PortDTO {
	Port string
	Target string
}