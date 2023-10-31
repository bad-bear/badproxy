package handlers

func addIPRateRule(c echo.Context) (err error) {
	r := new(IPRateReq)

	if err := c.Bind(r); err != nil {
		return c.String(http.StatusBadRequest, "bad post")
	}

	nr := IPRateDTO {
		Protocol: r.Protocol
		Source: r.Source
		Time: r.Time
		Timeframe: r.Timeframe
	}

	businesslogic(nr)

	return c.JSON(http.StatusOk, r)
}