package handlers

import (

)

func addFirewallRule(c echo.Context) (err error) {
	fw := new(FirewallRequest)
	if err := c.Bind(fw); err != nil {
		return c.String(http.StatusBadRequest, "bad post")
	}

	nfw := FirewallRequestDTO{
		Table: fw.Table
		Protocol: fw.Protocol
		Source: fw.Source
		Destination: fw.Destination
		Port: fw.Port
		Target: fw.Target
	}

	businesslogic(nfw)

	return c.JSON(http.StatusOk, fw)
}