package bindings

type IPRateReq {
	Protocol string `form:"Protocol"`
	Source string `form:"Source"`
	Time string `form:"Time"`
	Timeframe string `form:"Timeframe"`
}

type IPRateDTO {
	Protocol string
	Source string
	Time string
	Timeframe string
}