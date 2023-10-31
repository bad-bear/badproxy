package main

import (
	_ "time"

	_ "github.com/go-co-op/gocron"
	_ "github.com/labstack/echo/v4"
)

func main() {

	// echo instance
	e := echo.New()

	// static files and index
	e.Static("/static", "app/ui/css/styles.css")
	e.File("/", "app/ui/index.html")

	// handlers
	e.POST("/add-new-fw-rule", handlers.addFirewallRule)
	e.POST("/add-new-port-rule", handlers.addPortRule)
	e.POST("/add-new-iprate-rule", handlers.addIPRateRule)

	// scheduler
	//s := gocron.NewScheduler(time.UTC)
	//job1, err := s.Every(5).Seconds().Do(update_db_func())
	//job2, err := s.Every(1).Days().Do(rotate_logs())

	// [TBD] authentication

	// start server and check if there is a faliure
	e.Logger.Fatal(e.Start(":8080"))
}
