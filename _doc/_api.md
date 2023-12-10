# API Documentation
This document goes into detail about the API used for this application that connects the frontend to the backend of this proxy server.

### API List
- Add Firewall Rule
- Add Geo Rule
- Add Port Rule
- Add IP rate Rule
- Add Bandwidth Rule
- Archive Logs
- Get Dashboard Data

## General

## Add Firewall Rule
Entry Url: 'localhost:8080/add-frwl-firewall'

Method: POST

[Table] 
Field: "IP_Table"
Acceptable Inputs: "INPUT","OUTPUT"


[Protocol]
Field: "Protocol"
Acceptable Inputs: "all","icmp","tcp","udp"


[Source]
Field: "Source"
Acceptable Inputs: "8.8.8.8" or "8.8.8.0/24"

For source use either IP or cider format


[Destination]
Field: "Destination"
Acceptable Inputs: "8.8.8.8" or "8.8.8.0/24"

For source use either IP or cider format


[Port_Blocking]
Field: "Port_blocking"
Acceptable: "22"

This supports using an integer to represent the port


[Target]
Field: "Target"
Acceptable Inputs: "ACCEPT","DROP"


## Add Port Rule
Entry Url: 'localhost:8080/add-frwl-port'

Method: POST

[Port_Blocking]
Field: "Port_blocking"
Acceptable: "22"

This supports using an integer to represent the port


[Target]
Field: "Target"
Acceptable Inputs: "ACCEPT","DROP"


## Add Geo Rule
Entry Url: 'localhost:8080/add-frwl-geo'

Method: POST

[Location]
Field: "Location"
Acceptable: "Texas", "Arizona", etc.

This requires the name of the state


## Add IPRate Rule
Entry Url: 'localhost:8080/add-frwl-iprate'

Method: POST

[Source]
Field: "Source"
Acceptable Inputs: "8.8.8.8" or "8.8.8.0/24"

For source use either IP or cider format


[Protocol]
Field: "Protocol"
Acceptable Inputs: "all","icmp","tcp","udp"


[Connections]
Field: "TimeFrame"
Acceptable Inputs: "10", "100"

This value needs to be an integer to represent the amount of time


[Time]
Field: "TimeFrame"
Acceptable Inputs: "s", "m", "h", "d"

As a note the letters relate to different time frames
s - secs
m - minutes
h - hours
d - days


## Add IPRate Rule
Entry Url: 'localhost:8080/add-frwl-iprate'

Method: POST

[Source]
Field: "Source"
Acceptable Inputs: "8.8.8.8" or "8.8.8.0/24"

For source use either IP or cider format


[Protocol]
Field: "Protocol"
Acceptable Inputs: "all","icmp","tcp","udp"


[Throughput]
Field: "Bamount"
Acceptable Inputs: "10", "100"

This value needs to be an integer to represent the amount of throughput for the server to allow


[Type]
Field: "Type"
Acceptable Inputs: "p", "b", "kb", "mb"

As a note the letters relate to different byte rate sizes
p - packets
b - bytes
kb - kilobytes
mb - megabytes


[Time]
Field: "TimeFrame"
Acceptable Inputs: "s", "m", "h", "d"

As a note the letters relate to different time frames
s - secs
m - minutes
h - hours
d - days


## Get Dashboard Data
Entry Url: 'localhost:8080/dashboard-data'

Method: GET

This returns dashboard data and has no fields


## Get Archive Logs
Entry Url: 'localhost:8080/archive-logs'

Method: GET

This starts the data archival process
