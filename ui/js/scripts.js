/*!
    * Start Bootstrap - SB Admin v7.0.7 (https://startbootstrap.com/template/sb-admin)
    * Copyright 2013-2023 Start Bootstrap
    * Licensed under MIT (https://github.com/StartBootstrap/startbootstrap-sb-admin/blob/master/LICENSE)
    */
    // 
// Scripts
// 

window.addEventListener('DOMContentLoaded', event => {

    // Toggle the side navigation
    const sidebarToggle = document.body.querySelector('#sidebarToggle');
    if (sidebarToggle) {
        // Uncomment Below to persist sidebar toggle between refreshes
        // if (localStorage.getItem('sb|sidebar-toggle') === 'true') {
        //     document.body.classList.toggle('sb-sidenav-toggled');
        
        // }
        sidebarToggle.addEventListener('click', event => {
            event.preventDefault();
            document.body.classList.toggle('sb-sidenav-toggled');
            localStorage.setItem('sb|sidebar-toggle', document.body.classList.contains('sb-sidenav-toggled'));
        });
    }

});

async function update_dash() {
    dash_data = await fetch('/dashboard-data', {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
        },
    })
        .then(response => { return response.json() })
        .then(data => {
            console.log(JSON.parse(data));
            return JSON.parse(data)
        })

    let bandwidth_convrs_str = ""
    var bandwidth_convr = dash_data["Traf_bandwidth"] / 1000000000
    bandwidth_results = bandwidth_convrs_str.concat(bandwidth_convr.toString(), " GB")
    document.getElementById("traf_blocked").innerHTML = dash_data["Traf_blocked"]
    document.getElementById("traf_incoming").innerHTML = dash_data["Traf_total_in"]
    document.getElementById("traf_outgoing").innerHTML = dash_data["Traf_total_out"]
    document.getElementById("traf_bandwidth").innerHTML = bandwidth_results
    document.getElementById("traf_generic").innerHTML = dash_data["Traf_generic"]
    document.getElementById("traf_port").innerHTML = dash_data["Traf_port"]
    document.getElementById("traf_geo").innerHTML = dash_data["Traf_geo"]
}