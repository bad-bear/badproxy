package firewall

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"test_code/binding"
	"test_code/db"

	_ "github.com/mattn/go-sqlite3"
)

func AddFrwlPort(ptr binding.PortDTO) (err error) {

	app := "iptables" // iptables command

	label := "'[IPTABLES PORT]: '" // firewall log record

	// output logging string array
	args_LOG_OUTPUT := []string{"-A", "OUTPUT", "-p", "tcp", "--dport", ptr.Port, "-j", "LOG", "--log-prefix", label, "--log-level", "4"}

	cmd := exec.Command(app, args_LOG_OUTPUT...)

	// error stops program
	if stdO, err := cmd.Output(); err != nil {
		log.Println("ERROR [FW Port 1]: ", string(stdO))
	}

	args_LOG_INPUT := []string{"-A", "INPUT", "-p", "tcp", "--sport", ptr.Port, "-j", "LOG", "--log-prefix", label, "--log-level", "4"}

	cmd = exec.Command(app, args_LOG_INPUT...)

	// error stops program
	if stdO, err := cmd.Output(); err != nil {
		log.Println("ERROR [FW Port 2]: ", string(stdO))
	}

	args_OUTPUT := []string{"-A", "OUTPUT", "-p", "tcp", "--dport", ptr.Port, "-j", ptr.Target}

	cmd = exec.Command(app, args_OUTPUT...)

	// error stops program
	if stdO, err := cmd.Output(); err != nil {
		log.Println("ERROR [FW Port 3]: ", string(stdO))
	}

	args_INPUT := []string{"-A", "INPUT", "-p", "tcp", "--sport", ptr.Port, "-j", ptr.Target}

	cmd = exec.Command(app, args_INPUT...)

	// error stops program
	if stdO, err := cmd.Output(); err != nil {
		log.Println("ERROR [FW Port 4]: ", string(stdO))
		err = nil // none issue when ran appropriately
	}

	log.Println("Port Rule Added to Firewall")
	return err
}

func Add_frwl_firewall(ftr binding.FirewallRequestDTO) (err error) {

	app := "iptables" // iptables command

	label := "'[IPTABLES GENERIC]: '" // firewall log record

	// argument values for INPUT based tables
	args_INPUT := []string{"-A", "INPUT"}
	args_LOG_INPUT := []string{"-A", "INPUT"}

	// argument values for OUTPUT based tables
	args_OUTPUT := []string{"-A", "OUTPUT"}
	args_LOG_OUTPUT := []string{"-A", "OUTPUT"}

	// checks if protocol exist and adds to command
	if ftr.Protocol != "" {
		args_INPUT = append(args_INPUT, "-p", ftr.Protocol)
		args_LOG_INPUT = append(args_LOG_INPUT, "-p", ftr.Protocol)

		args_OUTPUT = append(args_OUTPUT, "-p", ftr.Protocol)
		args_LOG_OUTPUT = append(args_LOG_OUTPUT, "-p", ftr.Protocol)
	}

	// checks if Source exist and adds to command
	if ftr.Source != "" {
		args_INPUT = append(args_INPUT, "-s", ftr.Source)
		args_LOG_INPUT = append(args_LOG_INPUT, "-s", ftr.Source)

		args_OUTPUT = append(args_OUTPUT, "-s", ftr.Source)
		args_LOG_OUTPUT = append(args_LOG_OUTPUT, "-s", ftr.Source)
	}

	// checks if destination exist and adds to command
	if ftr.Destination != "" {
		args_INPUT = append(args_INPUT, "-d", ftr.Destination)
		args_LOG_INPUT = append(args_LOG_INPUT, "-d", ftr.Destination)

		args_OUTPUT = append(args_OUTPUT, "-d", ftr.Destination)
		args_LOG_OUTPUT = append(args_LOG_OUTPUT, "-d", ftr.Destination)
	}

	// checks if port exist and adds to command
	if ftr.Port != "" {
		if ftr.Table == "INPUT" {
			args_INPUT = append(args_INPUT, "--sport", ftr.Port)
			args_LOG_INPUT = append(args_LOG_INPUT, "--sport", ftr.Port)
		} else if ftr.Table == "OUTPUT" {
			args_OUTPUT = append(args_OUTPUT, "--dport", ftr.Port)
			args_LOG_OUTPUT = append(args_LOG_OUTPUT, "--dport", ftr.Port)
		} else if ftr.Table == "BOTH" {
			args_INPUT = append(args_INPUT, "--sport", ftr.Port)
			args_LOG_INPUT = append(args_LOG_INPUT, "--sport", ftr.Port)

			args_OUTPUT = append(args_OUTPUT, "--dport", ftr.Port)
			args_LOG_OUTPUT = append(args_LOG_OUTPUT, "--dport", ftr.Port)
		} else {
			log.Println("ERROR [FW Firewall]: ", "Missing Input Value")
			//return err
		}

	}

	//log.Println(ftr)
	args_INPUT = append(args_INPUT, "-j", ftr.Target)
	args_LOG_INPUT = append(args_LOG_INPUT, "-j", "LOG", "--log-prefix", label, "--log-level", "4")

	args_OUTPUT = append(args_OUTPUT, "-j", ftr.Target)
	args_LOG_OUTPUT = append(args_LOG_OUTPUT, "-j", "LOG", "--log-prefix", label, "--log-level", "4")

	//log.Println(args_OUTPUT, args_LOG_OUTPUT)
	//log.Println(args_INPUT, args_LOG_INPUT)

	// checks if target is INPUT and runs command
	if ftr.Table == "INPUT" {
		// create command to log input
		cmd := exec.Command(app, args_LOG_INPUT...)

		// error stops program
		if stdO, err := cmd.Output(); err != nil {
			log.Println("ERROR [FW Firewall ILogs]: ", string(stdO))
			//return err
		}

		// create command for IPtable
		cmd = exec.Command(app, args_INPUT...)

		// error stops program
		if stdO, err := cmd.Output(); err != nil {
			log.Println("ERROR [FW Firewall In]: ", string(stdO))
			//return err
		}

	}

	// checks if target is OUTPUT and runs command
	if ftr.Table == "OUTPUT" {
		// create command to log output
		cmd := exec.Command(app, args_LOG_OUTPUT...)

		// error stops program
		if stdO, err := cmd.Output(); err != nil {
			log.Println("ERROR [FW Firewall OLogs]: ", string(stdO))
			return err
		}

		// create command for IPtable
		cmd = exec.Command(app, args_OUTPUT...)

		// error stops program
		if stdO, err := cmd.Output(); err != nil {
			log.Println("ERROR [FW Firewall Out]: ", string(stdO))
			//return err
		}

	}

	// checks if target is BOTH and runs command
	if ftr.Table == "BOTH" {
		// create command to log output
		cmd := exec.Command(app, args_LOG_OUTPUT...)

		// error stops program
		if stdO, err := cmd.Output(); err != nil {
			log.Println("ERROR [FW Firewall OLogs]: ", string(stdO))
			//return err
		}

		// create command for IPtable
		cmd = exec.Command(app, args_OUTPUT...)

		// error stops program
		if stdO, err := cmd.Output(); err != nil {
			log.Println("ERROR [FW Firewall Out]: ", string(stdO))
			//return err
		}

		// create command to log input
		cmd = exec.Command(app, args_LOG_INPUT...)

		// error stops program
		if stdO, err := cmd.Output(); err != nil {
			log.Println("ERROR [FW Firewall ILogs]: ", string(stdO))
			//return err
		}

		// create command for IPtable
		cmd = exec.Command(app, args_INPUT...)

		// error stops program
		if stdO, err := cmd.Output(); err != nil {
			log.Println("ERROR [FW Firewall In]: ", string(stdO))
			//return err
			err = nil // none issue when ran appropriately
		}
	}

	log.Println("Firewall Rule Added to Firewall")
	return err
}

func Add_frwl_geo(gtr binding.GeoDTO) (err error) {

	app := "iptables" // iptables command

	label := "[IPTABLES GEO]: " // firewall log record

	db, err := sql.Open("sqlite3", db.Set_DB_PATH()) // opens database file

	// error stops program
	if err != nil {
		log.Println("ERROR [FW Geo]: ", err)
		// return err
	}

	defer db.Close() // close db later // close db later

	// prep statement for db execution
	stmt := fmt.Sprintf("SELECT network FROM geolite_city_blocks_ip WHERE geoname_id IN (SELECT geoname_id FROM geolite_city_locations_en WHERE subdivision_1_name = '%s')", gtr.Location)

	rows, err := db.Query(stmt)

	if err != nil {
		log.Println("ERROR [FW Geo]: ", err)
		// return err
	}
	defer rows.Close()

	count := 0
	source := ""

	args_LOG_INPUT := []string{"-A", "INPUT", "-j", "LOG", "--log-prefix", label, "--log-level", "4"}

	cmd := exec.Command(app, args_LOG_INPUT...)

	// error stops program after attempting to run
	if stdO, err := cmd.Output(); err != nil {
		log.Println("ERROR [FW Geo]: ", string(stdO))
		// return err
	}

	for rows.Next() {
		count = count + 1

		var network string

		// error stops program
		if err = rows.Scan(&network); err != nil {
			log.Println("ERROR [FW Geo]: ", err)
			// return err
		}

		source = source + "," + network

		if count == 500 {

			source = source[1:]
			//fmt.Println(source)

			args_INPUT := []string{"-A", "INPUT", "-s", source, "-j", "DROP"}
			cmd = exec.Command(app, args_INPUT...)

			// error stops program
			if err = cmd.Run(); err != nil {
				log.Println("ERROR [FW Geo]: ", err)
				// return err
			}
			source = ""
			count = 0
		}

	}

	log.Println("Geo Rule Added to Firewall")
	return err
}

func Add_frwl_iprate(iptr binding.IPRateDTO) (err error) {
	/*
		iptables -p all -s x.x.x.x -m connlimit --connlimit-above x -j DROP

	*/

	app := "iptables" // iptables command

	label := "'[IPTABLES IPRATE]: '" // firewall log record

	// output logging string array
	args_LOG_INPUT := []string{"-A", "INPUT", "-p", "all", "-s", iptr.Source, "-m", "connlimit", "--connlimit-above", iptr.Conn, "-j", "LOG", "--log-prefix", label, "--log-level", "4"}
	cmd := exec.Command(app, args_LOG_INPUT...)

	// error stops program after attempting to run
	if stdO, err := cmd.Output(); err != nil {
		log.Println("ERROR [FW IPrate Logs]: ", string(stdO))
		// return err
	}

	args_INPUT := []string{"-A", "INPUT", "-p", "all", "-s", iptr.Source, "-m", "connlimit", "--connlimit-above", iptr.Conn, "-j", "DROP"}
	cmd = exec.Command(app, args_INPUT...)

	// error stops program after attempting to run
	if stdO, err := cmd.Output(); err != nil {
		log.Println("ERROR [FW IPrate]: ", string(stdO))
		// return err
	}

	log.Println("IP Rate Rule Added to Firewall")
	return err
}

func Add_fwrl_bandwidth(bndw binding.BandwidthDTO) (err error) {

	/*
		// iptables -A INPUT -p x -s x.x.x.x --sport x -m hashlimit --hashlimit-mode srcip,srcport  --hashlimit-above 512kb/s -j REJECT
		//log.Println(iptr)
	*/
	app := "iptables" // iptables command

	label := "'[IPTABLES Bandwidth]: '" // firewall log record

	args_INPUT := []string{"-A", "INPUT"}
	args_LOG_INPUT := []string{"-A", "INPUT"}

	bandwidth_rate := "" + string(bndw.Bamount) + string(bndw.Type) + "/" + bndw.Time

	log.Println(bandwidth_rate)

	if bndw.Protocol != "" {
		args_LOG_INPUT = append(args_LOG_INPUT, "-p", bndw.Protocol)
		args_INPUT = append(args_INPUT, "-p", bndw.Protocol)
	} else {
		args_LOG_INPUT = append(args_LOG_INPUT, "-p", "all")
		args_INPUT = append(args_INPUT, "-p", "all")
	}

	if bndw.Source != "" {
		args_LOG_INPUT = append(args_LOG_INPUT, "-s", bndw.Source)
		args_INPUT = append(args_INPUT, "-s", bndw.Source)
	}

	args_LOG_INPUT = append(args_LOG_INPUT, "-m", "hashlimit", "--hashlimit-name", "Bandwidth Rate", "--hashlimit-mode", "srcip", "--hashlimit-above", bandwidth_rate, "-j", "LOG", "--log-prefix", label, "--log-level", "4")
	args_INPUT = append(args_INPUT, "-m", "hashlimit", "--hashlimit-name", "Bandwidth Rate", "--hashlimit-mode", "srcip", "--hashlimit-above", bandwidth_rate, "-j", "REJECT")

	//log.Println(args_LOG_INPUT)
	// output logging string array
	cmd := exec.Command(app, args_LOG_INPUT...)

	// error stops program after attempting to run
	if stdO, err := cmd.Output(); err != nil {
		log.Println("ERROR [FW Bandwidth Logs]: ", string(stdO))
		// return err
	}

	//log.Println(args_INPUT)
	// arguments for limiting bandwidth
	cmd = exec.Command(app, args_INPUT...)

	// error stops program after attempting to run
	if stdO, err := cmd.Output(); err != nil {
		log.Println("ERROR [FW Bandwidth]: ", string(stdO))
		// return err
	}

	log.Println("Bandwidth Rule Added")
	return err
}
