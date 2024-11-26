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


#Installation
##Prerequisites 

- Go: Version 1.16 or later. Download Go
- Chromium or Google Chrome: Required for browser automation with chromedp.


##Steps
1. Clone the Repository

```bash
git clone https://github.com/yourusername/domain-checker.git
cd domain-checker
```

2. Initialize Go Modules

```bash
go mod init domain_checker
```

3. Install Dependencies

```bash
go get -u github.com/chromedp/chromedp@latest
```

4. Install Chromium or Chrome

For Debian/Ubuntu:

```bash
sudo apt update
sudo apt install chromium
```

