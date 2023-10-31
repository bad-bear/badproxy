package handlers

func addPortRule(c echo.Context) (err error) {
	p := new(PortReq)

	if err := c.Bind(p); err != nil {
		return c.String(http.StatusBadRequest, "bad post")
	}

	np := PortDTO {
		Port: p.Port
		Target: p.Target
	}

	businesslogic(np)

	return c.JSON(http.StatusOk, p)
}