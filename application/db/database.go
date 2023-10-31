package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func db_() {
	db, err := sql.Open("sqlite3", "./bad.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
}

// checks for new logs to import into db
func update_db() {
	db, err := sql.Open("sqlite3", "./bad.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO data_logs() VALUES()")

	for log := logs {
		if _, err := stmt.Exec(); err != nil {
			log.Fatal(err)
		}
	}
}

// rotate logs
func rotate_logs() {

	db, err := sql.Open("sqlite3", "./bad.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rotateLogsSQL := `SELECT * FROM data_logs`
	rows, err := db.Query(rotateLogsSQL)

	t := time.Now()

	lgn := t + "_database"

	csvFile, err := os.Create("lgn")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	w.WriteAll(rows)

	db.Query(`DELETE * FROM data_logs`)
}

func add_rules() {

}

func add_logs() {

}
