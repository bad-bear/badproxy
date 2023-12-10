package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"test_code/binding"
	"test_code/db"
	"test_code/firewall"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

func AddPortRule(c echo.Context) (err error) {
	// init bind form request to variable
	fr := new(binding.PortReq)

	if err := c.Bind(fr); err != nil {
		return c.String(http.StatusBadRequest, "bad post")
	}

	/*
		// check if values are filled in properly
		if res, err := govalidator.ValidateStruct(fr); err != nil {
			log.Println("ERROR [Bad Form Request Error]: ", err)
			log.Println("ERROR [Bad Form Request Result]: ", res)
			return c.String(http.StatusNotAcceptable, "Bad Form")
		}
	*/

	// check if port is an integer
	if err := govalidator.IsInt(fr.Port) && govalidator.IsNotNull(fr.Port); !err {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Port Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	// check if Target is either ACCEPT or DROP
	if fr.Target != "ACCEPT" && fr.Target != "DROP" {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Target Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	// bind port data to var
	pd := binding.PortDTO{
		Port:   fr.Port,
		Target: fr.Target,
	}

	if err := db.Insert_db_port(pd); err != nil {
		log.Println("ERROR [Add Port DB]: ", err)
	}

	if err := firewall.AddFrwlPort(pd); err != nil {
		log.Println("ERROR [Add Port FW]: ", err)
	}

	return c.JSON(http.StatusOK, "Port Request Completed")
}

func AddFirewallRule(c echo.Context) (err error) {
	fr := new(binding.FirewallRequest)

	if err := c.Bind(fr); err != nil {
		return c.String(http.StatusBadRequest, "bad post")
	}

	/*
		if res, err := govalidator.ValidateStruct(fr); err != nil {
			log.Println("ERROR [Bad Form Request Error]: ", err)
			log.Println("ERROR [Bad Form Request Result]: ", res)
			return c.String(http.StatusNotAcceptable, "Bad Form")
		}
	*/

	//log.Println(fr)
	// check if Table is either OUTPUT or INPUT or BOTH
	if fr.Table != "OUTPUT" && fr.Table != "INPUT" && fr.Table != "BOTH" {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Table Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	// check if Protocol is either tcp or udp or icmp or all
	if fr.Protocol != "tcp" && fr.Protocol != "udp" && fr.Protocol != "icmp" && fr.Protocol != "all" {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Protocol Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	// check if Source is an IP
	if err := govalidator.IsIP(fr.Source) || govalidator.IsCIDR(fr.Source) || (govalidator.IsNull(fr.Source) && govalidator.IsNotNull(fr.Destination)); !err {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Source Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	// check if Destination is an IP
	if err := govalidator.IsIP(fr.Destination) || govalidator.IsCIDR(fr.Destination) || (govalidator.IsNull(fr.Destination) && govalidator.IsNotNull(fr.Source)); !err {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Destination Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	// check if port is an integer
	if err := govalidator.IsInt(fr.Port); !err {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Port Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	// check if Target is either ACCEPT or DROP
	if fr.Target != "ACCEPT" && fr.Target != "DROP" {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Target Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	fwd := binding.FirewallRequestDTO{
		Table:       fr.Table,
		Protocol:    fr.Protocol,
		Source:      fr.Source,
		Destination: fr.Destination,
		Port:        fr.Port,
		Target:      fr.Target,
	}

	if err := db.Insert_db_firewall(fwd); err != nil {
		log.Println("ERROR [Add Port DB]: ", err)
	}

	if err := firewall.Add_frwl_firewall(fwd); err != nil {
		log.Println("ERROR [Add Port FW]: ", err)
	}

	return c.JSON(http.StatusOK, "Firewall Request Completed")
}

func AddGeoRule(c echo.Context) (err error) {
	fr := new(binding.Geo)

	if err := c.Bind(fr); err != nil {
		return c.String(http.StatusBadRequest, "bad post")
	}

	if err := (govalidator.IsNotNull(fr.Location)) || (fr.Location != ""); !err {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Location Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	gd := binding.GeoDTO{
		Location: fr.Location,
	}

	if err := db.Insert_db_geo(gd); err != nil {
		log.Println("ERROR [Add Geo DB]: ", err)
	}

	if err := firewall.Add_frwl_geo(gd); err != nil {
		log.Println("ERROR [Add Geo FW]: ", err)
	}

	return c.JSON(http.StatusOK, "Geo Request Completed")
}

func AddIPrateRule(c echo.Context) (err error) {
	fr := new(binding.IPRate)

	if err := c.Bind(fr); err != nil {
		return c.String(http.StatusBadRequest, "bad post")
	}

	log.Println(fr)

	/*
		if res, err := govalidator.ValidateStruct(fr); err != nil {
			log.Println("ERROR [Bad Form Request Error]: ", err)
			log.Println("ERROR [Bad Form Request Result]: ", res)
			return c.String(http.StatusNotAcceptable, "Bad Form")
		}
	*/
	// check if Connections is an integer
	if err := govalidator.IsInt(fr.Conn); !err {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Port Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	// check if Source is an IP
	if err := govalidator.IsIP(fr.Source) || govalidator.IsCIDR(fr.Source); !err {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Source Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	// check if Protocol is either tcp or udp or icmp or all
	if fr.Protocol != "tcp" && fr.Protocol != "udp" && fr.Protocol != "icmp" && fr.Protocol != "all" {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Protocol Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	// check if Protocol is either tcp or udp or icmp or all
	if fr.TimeFrame != "s" && fr.TimeFrame != "m" && fr.TimeFrame != "h" && fr.TimeFrame != "d" && fr.TimeFrame == "" {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect TimeFrame Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	// check if connections is a number

	if err := govalidator.IsInt(fr.Conn) && govalidator.IsNotNull(fr.Conn); !err {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Connections Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	iprd := binding.IPRateDTO{
		Conn:      fr.Conn,
		Source:    fr.Source,
		Protocol:  fr.Protocol,
		TimeFrame: fr.TimeFrame,
	}

	if err := db.Insert_db_iprate(iprd); err != nil {
		log.Println("ERROR [Add IP Rate DB]: ", err)
	}

	if err := firewall.Add_frwl_iprate(iprd); err != nil {
		log.Println("ERROR [Add IP Rate FW]: ", err)
	}

	return c.JSON(http.StatusOK, "IP Rate Request Completed")
}

func AddBandwidthRule(c echo.Context) (err error) {
	fr := new(binding.Bandwidth)

	if err := c.Bind(fr); err != nil {
		return c.String(http.StatusBadRequest, "bad post")
	}

	/*
		if res, err := govalidator.ValidateStruct(fr); err != nil {
			log.Println("ERROR [Bad Form Request Error]: ", err)
			log.Println("ERROR [Bad Form Request Result]: ", res)
			return c.String(http.StatusNotAcceptable, "Bad Form")
		}
	*/

	// check if Protocol is either tcp or udp or icmp or all
	if fr.Protocol != "tcp" && fr.Protocol != "udp" && fr.Protocol != "icmp" && fr.Protocol != "all" {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Protocol Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	// check if B amount is an integer
	if err := govalidator.IsInt(fr.Bamount) && govalidator.IsNotNull(fr.Bamount); !err {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Bandwidth Amount Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	// check if Source is an IP
	if err := govalidator.IsIP(fr.Source) || govalidator.IsCIDR(fr.Source); !err {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Source Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	// check if Time is either tcp or udp or icmp or all
	if fr.Type != "p" && fr.Type != "b" && fr.Type != "kb" && fr.Type != "mb" {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Data Type Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	// check if Protocol is either tcp or udp or icmp or all
	if fr.Time != "s" && fr.Time != "m" && fr.Time != "h" && fr.Time != "d" {
		log.Println("ERROR [Bad Form Request Error]: ", "Incorrect Time Format")
		return c.String(http.StatusNotAcceptable, "Bad Form Input")
	}

	// bind form request to data object
	bnd := binding.BandwidthDTO{
		Source:   fr.Source,
		Protocol: fr.Protocol,
		Time:     fr.Time,
		Type:     fr.Type,
		Bamount:  fr.Bamount,
	}

	// Add data to DB
	if err := db.Insert_db_bandwidth(bnd); err != nil {
		log.Println("ERROR [Add Port DB]: ", err)
	}

	// Add Request to FW Rules
	if err := firewall.Add_fwrl_bandwidth(bnd); err != nil {
		log.Println("ERROR [Add Bandwidth FW]: ", err)
	}

	// Return ok
	return c.JSON(http.StatusOK, "Bandwidth Request Completed")
}

func GetDashboardData(c echo.Context) (err error) {
	// collect data
	d, err := db.Get_dashboard_data()

	// if an issue occurred
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed Dashboard Data Request")
	}

	// convert to string
	data, _ := json.Marshal(d)
	//log.Println(string(data))
	return c.JSON(http.StatusOK, string(data))
}

/*
func GetLogData(c echo.Context) (err error) {
	d := db.Get_log_data()
	data, _ := json.Marshal(d)
	print(data[1])
	return c.JSON(http.StatusOK, string(data))
}
*/

func RotateLogTable(c echo.Context) (err error) {
	// rotate logs
	if err := db.Rotate_log_table(); err != nil {
		return c.JSON(http.StatusOK, "Successful Log request")
	}

	return c.JSON(http.StatusOK, "Logs Rotated")
}
