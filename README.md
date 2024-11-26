# Domain-Checker-Tool
A high-performance Go-based tool for checking the availability and responsiveness of domains, utilizing both HTTP requests and browser automation for comprehensive analysis.

##Features

- Concurrent Domain Checking: Utilizes Go's goroutines to check multiple domains in parallel.
- DNS Resolution: Verifies if the domain can be resolved via DNS.
- HTTP Request Checking: Sends HTTP requests to domains and checks for specified successful status codes.
- Browser Automation: Uses chromedp to perform browser-based checks, handling JavaScript-rendered content.
- Redirection Handling: Detects and logs domain redirections.
- Customizable Parameters: Allows setting of successful status codes, timeout durations, and concurrency levels.
- Detailed Logging: Outputs active and inactive domains to separate files for easy analysis.
