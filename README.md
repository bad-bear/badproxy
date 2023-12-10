# badproxy
Proxy Application for CSCE 5585
Other documents are under the _doc folder


Bad Proxy is a system designed on Debian-based systems, also this application will require for there to be IP-tables and golang installed on the system for this application to work. This application runs by running IP-tables commands to manage a system's proxy. 

# Requirements
- linux (tested on ubuntu mate)
- rsyslog (application checks syslogs to get IP logs)
- golang 1.20+

## Database
The database needs to be unziped before use. GitHub size restrictions kept me from uploading the database so the database was zipped to compress to the file size limit for upload

# Installation

1. To setup, begin by downloading the repo:
	`link to gitrepo`

2. Go into the repo and find the setup script and run:
	`cd _setup/`
	`sudo ./setup`

3. After this is completed run the following commands:
	`export PATH=$PATH:/usr/local/go/bin`

4. Go to the main folder and run the following command:
`go build`

5. Run the following command to start the program:
`sudo ./<name of file>`




# Web Interface
The web interface currently doesn't have any authentication features though in that will be an improvement in the future.

## Access WebUI
To access the web interface go to the following:
[Proxy Interface](http://localhost:8080)

## Dashboard
The first page the application will open to is the home page which will show traffic data that is being collected from the web server. This will be where you will see the following information:

-  incoming/outgoing
-   Number of requests blocked by each rule.

## Logs
On the Menu Bar when on the Dashboard page there is a button 'Archive Data'. Clicking on this button will cause the server to rotate logs and store in the /db/data/archive folder 

## Firewall Rules
The links under this relate to creating block rules for the proxy. The blocks are the following:
- Firewall
- Port
- Geo
- IPRate
- Bandwidth
Each of these are forms with labels to show what are required or not for the form to accept a request. 

### Firewall Rules
The firewall page is where you can submit a general firewall request. The firewall section has the following fields:

Table       
Protocol    
Source      
Destination
Port        
Target

This should be used to input requests one by one. Keep in mind that based on the type of rule it will need to have the respective options filled out. Hitting submit will send a request to the API to run.

Optional Fields: Source, Destination, Port


### Port Rules
The Port Rule will add rules specific to Ports. The fields for this page are the following:

Port
Target

This will block on both input and output of a given service. Hitting submit will send a request to the API to run.

### Geo Rules
This page will be used for Geo based blocking. For this the blocking is centered around the states in the United States. The following fields:

Location

After selecting a state, hitting submit will send a request to the API to run.

### IP Rate
This page is used for submitting IP rate limiting requests to the server. The intention is to limit the total number of connections within a given time Period. The following fields:

Source
Connections
Protocol
TimeFrame


### Bandwidth
This page is dedicated to setting the Bandwidth rate for any given connection to the server. The following are the fields for this application:

Protocol
Source
TimeFrame
Type
Amount of Throughput
