package db

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"test_code/binding"
	"time"

	_ "github.com/mattn/go-sqlite3"
	iptables "github.com/moznion/go-iptables-logs-parser"
)

func Set_DB_PATH() string {

	cpath, _ := os.Getwd()
	log.Println(cpath)
	// set db path to use test or main db
	// db/data/core/proxy.db
	path := "../proxy.db"

	return path

}

func Insert_db_port(fd binding.PortDTO) (err error) {

	db, err := sql.Open("sqlite3", Set_DB_PATH()) // opens database file

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Port]: ", err)
		return err
	}

	defer db.Close() // close db later

	stmt, err := db.Prepare("INSERT INTO port_rules (Port, Target) VALUES(?, ?)")

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Port]: ", err)
		return err
	}

	// error stops program based on stmt execution
	if _, err := stmt.Exec(fd.Port, fd.Target); err != nil {
		log.Println("ERROR [DB Port]: ", err)
		return err
	}

	log.Println("Inserted Port Rule")
	return err
}

func Insert_db_firewall(ftr binding.FirewallRequestDTO) (err error) {
	db, err := sql.Open("sqlite3", Set_DB_PATH()) // opens database file

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Firewall]: ", err)
		return err
	}

	defer db.Close() // close db later

	stmt, err := db.Prepare("INSERT INTO firewall_rules (Chain, Protocol, Source, Destination, Port, Target) VALUES(?, ?, ?, ?, ?, ?)")

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Firewall]: ", err)
		return err
	}

	// error stops program
	if _, err := stmt.Exec(ftr.Table, ftr.Protocol, ftr.Source, ftr.Destination, ftr.Port, ftr.Target); err != nil {
		log.Println("ERROR [DB Firewall]: ", err)
		return err
	}

	log.Println("Inserted Firewall Rule")
	return err
}

func Insert_db_geo(gr binding.GeoDTO) (err error) {

	// opens database file
	db, err := sql.Open("sqlite3", Set_DB_PATH())

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Geo]: ", err)
	}

	// close db later
	defer db.Close()

	// prepare insert rule
	stmt, err := db.Prepare("INSERT INTO geo_rules (Location) VALUES(?)")

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Geo]: ", err)
	}

	// error stops program
	if _, err := stmt.Exec(gr.Location); err != nil {
		log.Println("ERROR [DB Geo]: ", err)
	}

	log.Println("Inserted Geo Rule")

	return err
}

func Insert_db_iprate(ipr binding.IPRateDTO) (err error) {
	db, err := sql.Open("sqlite3", Set_DB_PATH()) // opens database file

	// error stops program
	if err != nil {
		log.Println("ERROR [DB IPRate]: ", err)
	}

	defer db.Close() // close db later

	stmt, err := db.Prepare("INSERT INTO iprate_rules (Source, Conn, Protocol, TimeFrame) VALUES(?, ?, ?, ?)")

	// error stops program
	if err != nil {
		log.Println("ERROR [DB IPRate]: ", err)
	}

	// error stops program
	if _, err := stmt.Exec(ipr.Source, ipr.Conn, ipr.Protocol, ipr.TimeFrame); err != nil {
		log.Println("ERROR [DB IPRate]: ", err)
	}

	log.Println("Inserted Rate Rule")

	return err
}

func Insert_db_bandwidth(br binding.BandwidthDTO) (err error) {
	db, err := sql.Open("sqlite3", Set_DB_PATH()) // opens database file

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Bandwidth]: ", err)
		return err
	}

	defer db.Close() // close db later

	// prep to insert bandwidth rules
	stmt, err := db.Prepare("INSERT INTO bandwidth_rules (Source, Protocol, Time, Type, Bamount) VALUES(?, ?, ?, ?, ?)")

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Bandwidth]: ", err)
		return err
	}

	// error stops program
	if _, err := stmt.Exec(br.Source, br.Protocol, br.Time, br.Type, br.Bamount); err != nil {
		log.Println("ERROR [DB Bandwidth]: ", err)
		return err
	}

	log.Println("Inserted Bandwidth Rule")

	return err
}

func Update_log_table() (err error) {
	db, err := sql.Open("sqlite3", Set_DB_PATH()) // opens database file

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Log Table]: ", err)
		return err
	}

	defer db.Close() // close db later

	// prep to delete all records in traffic logs
	stmt, err := db.Prepare("DELETE FROM traffic_logs")

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Log Table]: ", err)
		return err
	}

	// error stops program
	if _, err := stmt.Exec(); err != nil {
		log.Println("ERROR [DB Log Table]: ", err)
		return err
	}

	// prep insert of traffic data
	stmt2, err := db.Prepare("INSERT INTO traffic_logs (Timestamp, Hostname, KernelTimestamp, Prefix, InputInterface, OutputInterface, MACAddress, Source, Destination,Length,ToS,Precedence,TTL,ID,CongestionExperienced,DoNotFragment,MoreFragmentsFollowing,Frag,IPOptions,Protocol,Type,Code,SourcePort,DestinationPort,Sequence,AckSequence,WindowSize,Res,Urgent,Ack,Push,Reset,Syn,Fin,Urgp,TCPOption) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Log Table]: ", err)
		return err
	}

	// open file
	f, err := os.Open("/var/log/syslog")
	if err != nil {
		log.Println("ERROR [DB Log Table]: ", err)
		return err
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	var iprec_check bool

	// loop through all lines in the custom iptables logs
	for scanner.Scan() {

		iprec_check = strings.Contains(scanner.Text(), "IPTABLES")

		if iprec_check {
			// parse line of IP table log
			parsedLog, err := iptables.Parse(scanner.Text())

			// error stops program
			if err != nil {
				break
			}

			// insert into db traffic log table
			_, err = stmt2.Exec(parsedLog.Timestamp, parsedLog.Hostname, parsedLog.KernelTimestamp, parsedLog.Prefix, parsedLog.InputInterface, parsedLog.OutputInterface, parsedLog.MACAddress, parsedLog.Source, parsedLog.Destination, parsedLog.Length, parsedLog.ToS, parsedLog.Precedence, parsedLog.TTL, parsedLog.ID, parsedLog.CongestionExperienced, parsedLog.DoNotFragment, parsedLog.MoreFragmentsFollowing, parsedLog.Frag, parsedLog.IPOptions, parsedLog.Protocol, parsedLog.Type, parsedLog.Code, parsedLog.SourcePort, parsedLog.DestinationPort, parsedLog.Sequence, parsedLog.AckSequence, parsedLog.WindowSize, parsedLog.Res, parsedLog.Urgent, parsedLog.Ack, parsedLog.Push, parsedLog.Reset, parsedLog.Syn, parsedLog.Fin, parsedLog.Urgp, parsedLog.TCPOption)

			// checks if db traffic was inserted
			if err != nil {
				log.Println("ERROR [DB Log Table]: ", err)
				return err
			}

		} else {
			continue
		}

	}

	/*
		// clear syslogs of all IP Table related messages
		sed_app := "sed"
		sed_inputs := []string{"-i", "/IPTABLES/d", "/var/log/syslog"}
		cmd := exec.Command(sed_app, sed_inputs...)

		// error stops program
		if err := cmd.Run(); err != nil {
			log.Println("ERROR [DB Update Logs] ", err.Error())
			return err
		}
	*/

	// error stops program
	if err := scanner.Err(); err != nil {
		log.Println("ERROR [DB Update Logs]: ", err)
		return err
	}

	log.Println("Updated Log Table")
	return err
}

func Update_dashboard_values() (err error) {
	db, err := sql.Open("sqlite3", Set_DB_PATH()) // opens database file

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
	}

	defer db.Close() // close db later

	// prep statement for updating Geo Blocked
	str := "UPDATE dashboard_data SET traffic_geo = (SELECT COUNT(Prefix) FROM traffic_logs WHERE Prefix LIKE " + "'%GEO%'" + ")"
	stmt, err := db.Prepare(str)

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	// error stops program && execute db code
	if _, err := stmt.Exec(); err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	// prep statement for updating Bandwidth Blocked
	str = "UPDATE dashboard_data SET traffic_port = (SELECT COUNT(Prefix) FROM traffic_logs WHERE Prefix LIKE " + "'%PORT%'" + ")"
	stmt, err = db.Prepare(str)

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	// error stops program && execute db code
	if _, err := stmt.Exec(); err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	// prep statement for updating Generic Blocked
	str = "UPDATE dashboard_data SET traffic_generic = (SELECT COUNT(Prefix) FROM traffic_logs WHERE Prefix LIKE " + "'%Generic%'" + ")"
	stmt, err = db.Prepare(str)

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	// error stops program && execute db code
	if _, err := stmt.Exec(); err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	// prep statement for updating Bandwidth
	stmt, err = db.Prepare("UPDATE dashboard_data SET traffic_bandwidth = (SELECT SUM(length) FROM traffic_logs)")

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	// error stops program && execute db code
	if _, err := stmt.Exec(); err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	// prep statement for updating tcp
	stmt, err = db.Prepare("UPDATE dashboard_data SET traffic_tcp = (SELECT SUM(length) FROM traffic_logs WHERE Protocol = 'TCP' or Protocol = 'tcp')")

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	// error stops program && execute db code
	if _, err := stmt.Exec(); err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	// prep statement for updating icmp
	stmt, err = db.Prepare("UPDATE dashboard_data SET traffic_icmp = (SELECT SUM(length) FROM traffic_logs WHERE Protocol = 'ICMP' or Protocol = 'icmp')")

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	// error stops program && execute db code
	if _, err := stmt.Exec(); err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	// prep statement for updating udp
	stmt, err = db.Prepare("UPDATE dashboard_data SET traffic_tcp = (SELECT SUM(length) FROM traffic_logs WHERE Protocol = 'UDP' or Protocol = 'udp')")

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	// error stops program && execute db code
	if _, err := stmt.Exec(); err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	empty := "''"
	str = "UPDATE dashboard_data SET traffic_total_in = (SELECT COUNT(*) FROM traffic_logs WHERE InputInterface != +" + empty + ")"

	// prep statement for updating incoming traffic value
	stmt, err = db.Prepare(str)

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	// error stops program && execute db code
	if _, err := stmt.Exec(); err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	empty = "''"
	str = "UPDATE dashboard_data SET traffic_total_out = (SELECT COUNT(*) FROM traffic_logs WHERE OutputInterface != +" + empty + ")"

	// prep statement for updating outgoing traffic value
	stmt, err = db.Prepare(str)

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	// error stops program && execute db code
	if _, err := stmt.Exec(); err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
		return err
	}

	log.Println("Updated Dashboard Values")
	return err
}

func Rotate_log_table() (err error) {

	// update logs
	if err := Update_log_table(); err != nil {
		log.Println("ERROR [DB Update Log]: ", err)
	}

	// create data holder with headers
	export_csv := [][]string{{"Timestamp", "Hostname", "KernelTimestamp", "Prefix", "InputInterface", "OutputInterface", "MACAddress", "Source", "Destination", "Length", "ToS", "Precedence", "TTL", "ID", "CongestionExperienced", "DoNotFragment", "MoreFragmentsFollowing", "Frag,IPOptions", "Protocol", "Type", "Code", "SourcePort", "DestinationPort", "Sequence", "AckSequence", "WindowSize", "Res", "Urgent", "Ack", "Push", "Reset", "Syn", "Fin", "Urgp", "TCPOption"}}

	db, err := sql.Open("sqlite3", Set_DB_PATH()) // opens database file

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Rotate Logs]: ", err)
		return err
	}

	defer db.Close() // close db later

	// prep to get all data from traffic logs
	stmt, err := db.Query("SELECT * FROM traffic_logs")

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Rotate Logs]: ", err)
		return err
	}

	// loop through all select data
	for stmt.Next() {

		var ilog ip_log // declare struct variable to collect data

		// scan values into the struct variable
		err = stmt.Scan(&ilog.Timestamp, &ilog.Hostname, &ilog.KernelTimestamp, &ilog.Prefix, &ilog.InputInterface, &ilog.OutputInterface, &ilog.MACAddress, &ilog.Source, &ilog.Destination, &ilog.Length, &ilog.ToS, &ilog.Precedence, &ilog.TTL, &ilog.ID, &ilog.CongestionExperienced, &ilog.DoNotFragment, &ilog.MoreFragmentsFollowing, &ilog.Frag, &ilog.IPOptions, &ilog.Protocol, &ilog.Type, &ilog.Code, &ilog.SourcePort, &ilog.DestinationPort, &ilog.Sequence, &ilog.AckSequence, &ilog.WindowSize, &ilog.Res, &ilog.Urgent, &ilog.Ack, &ilog.Push, &ilog.Reset, &ilog.Syn, &ilog.Fin, &ilog.Urgp, &ilog.TCPOption)

		// error stops program
		if err != nil {
			log.Println("ERROR [DB Rotate Logs]: ", err)
			return err
		}

		// a string array to append for the csv
		a_str := []string{ilog.Timestamp, ilog.Hostname, ilog.KernelTimestamp, ilog.Prefix, ilog.InputInterface, ilog.OutputInterface, ilog.MACAddress, ilog.Source, ilog.Destination, strconv.Itoa(ilog.Length), strconv.Itoa(ilog.ToS), strconv.Itoa(ilog.Precedence), strconv.Itoa(ilog.TTL), ilog.ID, ilog.CongestionExperienced, ilog.DoNotFragment, ilog.MoreFragmentsFollowing, strconv.Itoa(ilog.Frag), ilog.IPOptions, ilog.Protocol, ilog.Type, ilog.Code, ilog.SourcePort, ilog.DestinationPort, ilog.Sequence, ilog.AckSequence, ilog.WindowSize, ilog.Res, ilog.Urgent, ilog.Ack, ilog.Push, ilog.Reset, ilog.Syn, ilog.Fin, ilog.Urgp, ilog.TCPOption}

		// append record into the array
		export_csv = append(export_csv, a_str)
	}

	// get the current time
	curr_time := time.Now().UTC()

	//label for the archive
	file_name := curr_time.Format("2006-01-02") + "_logs.csv"

	// create csv file
	file, err := os.Create(file_name)

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Rotate Logs]: ", err)
		return err
	}

	defer file.Close() // close file later

	// error stops program
	w := csv.NewWriter(file)

	defer w.Flush() // write changes later

	// write values in the csv to the file
	if err := w.WriteAll(export_csv); err != nil {
		log.Println("ERROR [DB Rotate Logs]: ", err)
		return err
	}

	// prep to delete old traffic log data
	stmt2, err := db.Prepare("DELETE FROM traffic_logs")

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Rotate Logs]: ", err)
		return err
	}

	// error stops program && execute db code
	if _, err := stmt2.Exec(); err != nil {
		log.Println("ERROR [DB Rotate Logs]: ", err)
		return err
	}

	/*
		if err := os.Truncate("/var/log/fw.log", 0); err != nil {
			log.Printf("Failed to truncate: %v", err)
		}
	*/

	log.Println("Logs Rotated: ", file_name)
	return err
}

func Get_dashboard_data() (dd DashData, err error) {
	if err := Update_dashboard_values(); err != nil {
		log.Println("ERROR [DB Update Dashboard]: ", err)
	} // update values before getting them

	db, err := sql.Open("sqlite3", Set_DB_PATH()) // opens database file

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Dashboard]: ", err)
	}

	defer db.Close() // close db later

	// prep to get all traffic data which is a single preset record
	stmt := db.QueryRow("SELECT traffic_blocked, traffic_total_in, traffic_total_out, traffic_bandwidth, traffic_generic, traffic_port, traffic_geo, traffic_tcp, traffic_udp, traffic_icmp FROM dashboard_data WHERE id = 1")

	// scan values to data struct
	err = stmt.Scan(&dd.Traf_blocked, &dd.Traf_total_in, &dd.Traf_total_out, &dd.Traf_bandwidth, &dd.Traf_generic, &dd.Traf_port, &dd.Traf_geo, &dd.Traf_tcp, &dd.Traf_udp, &dd.Traf_icmp)

	// error stops program
	if err != nil {
		log.Println("NOTICE [DB Dashboard]: ", err)
		err = nil
	}

	return dd, err
}

/*
func Get_log_data() (all_logs []ip_log, err error) {

	if err := Update_log_table(); err != nil {
		log.Println("ERROR [DB Logs]: ", err)
	} //Update logs before archiving

	db, err := sql.Open("sqlite3", Set_DB_PATH()) // opens database file

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Logs]: ", err)
	}

	defer db.Close() // close db later

	stmt, err := db.Query("SELECT * FROM traffic_logs")

	// error stops program
	if err != nil {
		log.Println("ERROR [DB Logs]: ", err)
	}

	for stmt.Next() {
		var ilog ip_log
		err = stmt.Scan(&ilog.Timestamp, &ilog.Hostname, &ilog.KernelTimestamp, &ilog.Prefix, &ilog.InputInterface, &ilog.OutputInterface, &ilog.MACAddress, &ilog.Source, &ilog.Destination, &ilog.Length, &ilog.ToS, &ilog.Precedence, &ilog.TTL, &ilog.ID, &ilog.CongestionExperienced, &ilog.DoNotFragment, &ilog.MoreFragmentsFollowing, &ilog.Frag, &ilog.IPOptions, &ilog.Protocol, &ilog.Type, &ilog.Code, &ilog.SourcePort, &ilog.DestinationPort, &ilog.Sequence, &ilog.AckSequence, &ilog.WindowSize, &ilog.Res, &ilog.Urgent, &ilog.Ack, &ilog.Push, &ilog.Reset, &ilog.Syn, &ilog.Fin, &ilog.Urgp, &ilog.TCPOption)

		// error stops program
		if err != nil {
			log.Println("ERROR [DB Logs]: ", err)
		}

		all_logs = append(all_logs, ilog)
	}

	return all_logs, err
}
*/
