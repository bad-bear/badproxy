package main

import (
	"log"
	"net/http"
	"test_code/db"
	"test_code/handlers"

	"github.com/labstack/echo/v4"
)

func main() {

	// init new web server instance
	e := echo.New()

	//
	//govalidator.SetFieldsRequiredByDefault(true)

	// css/js
	e.File("/static/css", "ui/css/styles.css")
	e.File("/bzPvBe6ghJ", "ui/js/script.js")

	// routes
	e.File("/", "ui/html/index.html")
	//e.File("/logs", "ui/html/logs.html")
	e.File("/add-firewall-rule-port", "ui/html/add-firewall-rule-port.html")
	e.File("/add-firewall-rule-geo", "ui/html/add-firewall-rule-geo.html")
	e.File("/add-firewall-rule-firewall", "ui/html/add-firewall-rule-firewall.html")
	e.File("/add-rate-rule-bandwidth", "ui/html/add-rate-rule-bandwidth.html")
	e.File("/add-rate-rule-iprate", "ui/html/add-rate-rule-iprate.html")

	// form posts
	e.POST("/add-frwl-firewall", handlers.AddFirewallRule)
	e.POST("/add-frwl-port", handlers.AddPortRule)
	e.POST("/add-frwl-geo", handlers.AddGeoRule)
	e.POST("/add-rt-bandwidth", handlers.AddBandwidthRule)
	e.POST("/add-rt-iprate", handlers.AddIPrateRule)

	// data gets
	e.GET("/dashboard-data", handlers.GetDashboardData)
	//e.GET("/log-data", handlers.GetLogData)

	// reload
	e.GET("/archive-logs", handlers.RotateLogTable)

	/*
		// initialize database
		if err := db.Update_log_table(); err != nil {
			log.Println("ERROR [Log Table]: ", err)
		}
	*/

	if err := db.Update_dashboard_values(); err != nil {
		log.Println("ERROR [Dashboard]: ", err)
	}

	// start server with TLS

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Println("ERROR [Server]:", err)
	}

}
