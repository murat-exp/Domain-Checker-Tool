# Domain Checker Tool

**A high-performance Go-based tool for checking the availability and responsiveness of domains, utilizing both HTTP requests and browser automation for comprehensive analysis.**

---

## Features

- **Concurrent Domain Checking**: Utilizes Go's goroutines to check multiple domains in parallel.
- **DNS Resolution**: Verifies if the domain can be resolved via DNS.
- **HTTP Request Checking**: Sends HTTP requests to domains and checks for specified successful status codes.
- **Browser Automation**: Uses `chromedp` to perform browser-based checks, handling JavaScript-rendered content.
- **Redirection Handling**: Detects and logs domain redirections.
- **Customizable Parameters**: Allows setting of successful status codes, timeout durations, and concurrency levels.
- **Detailed Logging**: Outputs active and inactive domains to separate files for easy analysis.

---

## Installation

### Prerequisites

- **Go**: Version **1.16** or later. [Download Go](https://golang.org/dl/)
- **Chromium or Google Chrome**: Required for browser automation with `chromedp`.

### Steps

1. **Clone the Repository**

   ```bash
   git clone https://github.com/murat-exp/Domain-Checker-Tool.git
   cd domain-checker
   ```

2. **Initialize Go Modules**

   ```bash
   go mod init domain_checker
   ```

3. **Install Dependencies**

   ```bash
   go get -u github.com/chromedp/chromedp@latest
   ```

4. **Install Chromium or Chrome**

   For Debian/Ubuntu:

   ```bash
   sudo apt update
   sudo apt install chromium
   ```

---

## Usage

```bash
go run domain_checker.go [domain_list_file] [status_codes]
```

- **`domain_list_file`**: Path to the file containing the list of domains (one per line).
- **`status_codes`** (Optional): Comma-separated HTTP status codes considered successful (default: `200`).

### Examples

- **Default status code `200`:**

   ```bash
   go run domain_checker.go domains.txt
   ```

- **Custom status codes `200`, `301`, `302`:**

   ```bash
   go run domain_checker.go domains.txt 200,301,302
   ```

---

## Example Output

### Console Output

```bash
Active: http://example.com (Status Code: 200)
Active: https://test.com (Status Code: 200) (Redirected to: newdomain.com)
```

### Generated Files

- **`active_domains.txt`**: List of active domains with optional redirection info.
- **`inactive_domains.txt`**: List of inactive or unreachable domains.

---

## Configuration

Customize the tool by modifying variables in `domain_checker.go`:

- **Successful Status Codes**

   ```go
   successStatusCodes = []int{200, 301, 302}
   ```

- **Timeout Duration**

   ```go
   timeout = 10 * time.Second
   ```

- **Retry Count**

   ```go
   retryCount = 2
   ```

- **Max Concurrent Checks**

   ```go
   maxConcurrentChecks = 100
   ```

---

## How It Works

1. **Domain Reading**: Reads domains from a file, ensuring each is non-empty.

2. **DNS Resolution**: Checks if each domain is resolvable; unresolvable domains are marked inactive.

3. **HTTP Requests**:
   - Sends GET requests over `http://` and `https://`.
   - Follows redirects up to 10 levels.
   - Validates response status codes.

4. **Browser Automation**:
   - Uses `chromedp` for headless browser checks.
   - Navigates to the final URL.
   - Can be extended to perform content verification.

5. **Concurrency Control**:
   - Employs goroutines and a semaphore pattern.
   - Limits concurrent checks to prevent resource exhaustion.

6. **Logging**:
   - Active domains are logged to `active_domains.txt`.
   - Inactive domains are logged to `inactive_domains.txt`.
   - Console output provides real-time feedback.

---

## Dependencies

- **Go Modules**: For dependency management.

   ```bash
   go mod init domain_checker
   ```

- **Chromedp**: For headless browser automation.

   ```bash
   go get -u github.com/chromedp/chromedp@latest
   ```

- **Chromium or Google Chrome**: Required by `chromedp`.

---

## Troubleshooting

1. **No Output or Errors**:
   - Ensure Go modules are initialized (`go mod init domain_checker`).
   - Install dependencies (`go get -u github.com/chromedp/chromedp`).

2. **Chromium Not Found**:
   - Install Chromium or Chrome:

      ```bash
      sudo apt update
      sudo apt install chromium
      ```

3. **Go Version Issues**:
   - Verify Go version:

      ```bash
      go version
      ```

   - Update Go if necessary.

4. **Permission Errors**:
   - Run commands without `sudo` unless necessary.
   - Check file permissions for `active_domains.txt` and `inactive_domains.txt`.

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Submit a pull request with a detailed description of your changes.

---

## Acknowledgments

Special thanks to:

- **[Chromedp](https://github.com/chromedp/chromedp)**: For browser automation capabilities.
- The **Go Community**: For providing a robust ecosystem for high-performance applications.

  ---

  checker.png
