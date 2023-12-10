For this project, the idea that I had for this was to implement a client-server architecture. The application will run a small web server with a web interface that will allow for users to input firewall rules which will be applied to the server.

## General
The main guide and design for this application was based on the information by the book of Ben Hudson "Echo Quick Start Guide". The book goes through explaining echo by creating an MVC model which is why the application is based and designed in this manner.

## Application Model
The design for this application is based on the Model-View-Controller (MVC). This will be implemented by using a combination of GoLang and HTML for the backend and frontend respectively. The following table shows an in-depth of the application tech stack:

| Stack | Technology|
--
| Database | SQLite |
| Model | GoLang |
| View | HTML/JS/CSS |
| Controller | GoLang |
| Web Server | echo (GoLang) |

## Application User Interface (View)
For the user interface. The basis was referenced and copied from a free to download a Bootstrap Admin Dashboard from "Start Bootstrap" as a basis for the general view. Changes were made to better suit the application. The first page is to show all of the data that would be to be reflected in the dashboard. This was done so that an admin can see the data views when they first access the website

The navigation is to the left organized in a structured way so that all data can be viewed. The logs are first as it would ideally be the first place an admin would look. Following that are the rule forms for the firewall and following that are the rate rules. 

Each of the pages are forms that can be filled out with there being optional and required fields for each type of form. 

For the dashboard data, a function was created to run to collect this information by calling an API every 5 seconds. 

Additionally, because of timing and parallel processing failures to rotate logs is something to be created with an interaction by clicking on the archive button.

## Application Data (Model)
For this aspect of the design because I am using a relational database (SQLite). I created multiple models for each kind of form/request type:

Models
- Forms
	- Firewall
	- Port
	- Geo
	- Bandwidth
	- Connection Limit
- Backend
	- GeoMax IP Table
	- Traffic Logs

The Form models are related to each unique form to collect information and data to send over to the backend to process and funnel into the database. Their direction is going only one-way from the web UI to the backend. 

The Backend models used for data going to the dashboard and log views. This data is collected from the logs on the server and sent to the database on a regular basis.

## Application Logic (Controller)
When the application starts there are a few processes that run:

- update database logs
- update dashboard records
- web server

The update processes are done to load logs from the syslog file and update the values for when the dashboard calls for the data.

### IP Commands
All interaction with IP tables is done using the golang exec and require sudo permissions for success. This is because trying to use go-iptables, there is little documentation and when attempting to implement, I ran into more problems trying to get it implemented. 

### Requests
The server is listening to specific apis created for each form type based on the original design specification. There is a request going in the reverse to call on the data for the database. 

### Limitations
Because of the method of using golang exec and forms, there are a number of limitations that this can't do yet though with time can be improved on. For example the freedom of all possible fields for a iptables command are not enabled as functionality and practically for implementation was severely limiting.

### Testing Results
The results shown in the image while generally accurate there is a significant issue with it. This is because the application during the testing phase wasn't allowed to be ran with sudo permissions which limited a number of results that required sudo permissions. 

Also, the reason the image shows a failure is not because of the actual code, but because of a database lock as the functions called use a large amount of database transactions which caused failures when running the coverage reports all at one time for total coverage percentages.

Aside from this the main function doesn't have a test created because it holds the echo file and when ran will fail as there is nothing to stop it generally aside from the coverage timeout


Reference:
Echo Quick Start Guide By - Ben Huson (PacktPub)
Bootstrap 5 - getbootstrap.com
Start Bootstrap - https://startbootstrap.com/template/sb-admin