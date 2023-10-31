


# Bad_Proxy
A web proxy for website traffic. Developed using GoLang

Proxy Firewall Development with Dashboard
-----------------------------------------

Individual assignment

### Objective

To design and develop a proxy firewall with a dashboard using one of the given languages (Go, Rust, C++, C, Zig). The firewall will monitor and control incoming and outgoing network traffic based on specified criteria.

### Requirements

1.  **Basic Firewall Features**:
    -   Source/Destination IP blocking.
    -   Source/Destination port blocking.
    -   Protocol-based blocking (e.g., block all HTTPS traffic).
2.  **Rate Limiting**
    -   Limit the number of requests from a particular source IP within a specific time frame (e.g., max 100 requests per minute).
    -   Limit the total bandwidth used by a particular IP or service.
3.  **Geo-blocking**
    -   Block or allow traffic based on the geographic location of the source or destination IP. Utilize a GeoIP database for this purpose.
4.  **Logging**
    -   Log all blocked traffic with timestamps, source/destination IPs, and the reason for blocking.
    -   Implement a mechanism to regularly rotate and archive logs.
5.  **Dashboard & User Interface**
    -   Develop a web-based UI that allows an admin to set rules, view logs, and monitor the system.
    -   Display statistics such as:
        -   Total traffic (incoming/outgoing).
        -   Traffic breakdown by protocol, IP, and geography.
        -   Number of requests blocked by each rule.
        -   Bandwidth usage over time.

1.  **Test Coverage**:
    -   Implement unit and integration tests for both the firewall logic and the dashboard.
    -   Aim for at least 80% code coverage.
2.  **Static Analysis (Linting)**:
    -   Code should adhere to the best practices and standards for the chosen language.
    -   Tools like `golangci-lint` for Go, `clippy` for Rust, or equivalents for C++, C, or Zig should be utilized.
3.  **SSL Inspection (Optional)**:
    -   Implement a mechanism to inspect SSL/TLS encrypted traffic.

### Deliverables:

1.  **Source Code: **must be properly organized, commented, and structured.
2.  **Documentation**:
    -   Architecture and design decisions.
    -   User manual: setup, configuration, and operation of the firewall and dashboard.
    -   API documentation (if any).
3.  **Test Reports**: Coverage reports from testing.
4.  **Lint Reports**: Generated output or screenshots from linting tools.

### Evaluation Criteria:

1.  **Functionality**: Does the firewall and dashboard work as specified?
2.  **Code Quality**: Maintainability, clarity, and efficiency.
3.  **Robustness**: How well does it handle errors and edge cases?
4.  **Test Coverage**: Breadth and quality of tests.
5.  **Documentation Quality**: How well can someone else understand and use your work based on your documentation?
6.  **Security**: The firewall should not introduce new vulnerabilities.
7.  **Usability**: How intuitive is the dashboard and user interface?

### Implementation Tips:

1.  Choose a language you're comfortable with from the given list. Each has its own strengths and community support.
2.  Libraries:
    -   For Go: `echo` for a minimalist web framework or even standard lib for a simple web server, and `go-iptables` for interfacing with Linux IPTables.
    -   For Rust: rocket for web framework, `rust-iptables` for IPTables interaction.
    -   For C++/C: crow or `cpp-httplib` for creating a simple web server.
    -   Zig: Given its relative newness, you may have to rely more on its standard library and potentially interface with C libraries.
3.  For the GeoIP database, consider using the MaxMind GeoIP2 database and associated libraries in the chosen language.
4.  Keep security in mind when developing the dashboard. Ensure inputs are sanitized and authentication is enforced.

### Submission

Please submit a link to your Github repo with your code, screenshots, and documentation.