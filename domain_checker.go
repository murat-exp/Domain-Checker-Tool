// domain_checker.go

package main

import (
    "bufio"
    "context"
    "crypto/tls"
    "fmt"
    "io"
    "net"
    "net/http"
    "os"
    "strings"
    "sync"
    "time"

    "github.com/chromedp/chromedp"
)

// Configuration variables
var (
    // Successful HTTP status codes
    successStatusCodes = []int{200}

    // Timeout for HTTP requests and browser navigation
    timeout = 10 * time.Second

    // Number of retry attempts for HTTP requests
    retryCount = 2

    // Maximum number of concurrent domain checks
    maxConcurrentChecks = 100
)

func main() {
    // Check command-line arguments
    if len(os.Args) < 2 {
        fmt.Println("Usage: domain_checker domain_list_file [status_codes]")
        os.Exit(1)
    }

    domainFile := os.Args[1]

    // Parse successful status codes if provided
    if len(os.Args) >= 3 {
        codes := strings.Split(os.Args[2], ",")
        successStatusCodes = []int{}
        for _, codeStr := range codes {
            var code int
            fmt.Sscanf(strings.TrimSpace(codeStr), "%d", &code)
            successStatusCodes = append(successStatusCodes, code)
        }
    }

    // Read domains from file
    domains, err := readDomainsFromFile(domainFile)
    if err != nil {
        fmt.Println("Failed to read domain list:", err)
        os.Exit(1)
    }

    // Start checking domains
    checkAllDomains(domains)
}

// Reads domains from the specified file
func readDomainsFromFile(fileName string) ([]string, error) {
    file, err := os.Open(fileName)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var domains []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        domain := strings.TrimSpace(scanner.Text())
        if domain != "" {
            domains = append(domains, domain)
        }
    }
    return domains, scanner.Err()
}

// Checks all domains concurrently
func checkAllDomains(domains []string) {
    // Semaphore to limit concurrency
    sem := make(chan struct{}, maxConcurrentChecks)
    var wg sync.WaitGroup

    for _, domain := range domains {
        wg.Add(1)
        go func(domain string) {
            defer wg.Done()
            sem <- struct{}{}
            checkDomain(domain)
            <-sem
        }(domain)
    }
    wg.Wait()
}

// Checks a single domain
func checkDomain(domain string) {
    // DNS resolution
    resolvable := isDomainResolvable(domain)
    if !resolvable {
        appendToFile("inactive_domains.txt", domain+"\n")
        return
    }

    protocols := []string{"http://", "https://"}
    isActive := false

    for _, protocol := range protocols {
        url := protocol + domain
        response, err := tryRequest(url)
        if err == nil && contains(successStatusCodes, response.StatusCode) {
            finalURL := response.Request.URL.String()
            originalHostname := domain
            finalHostname := response.Request.URL.Hostname()
            redirectInfo := ""
            if originalHostname != finalHostname {
                redirectInfo = fmt.Sprintf(" (Redirected to: %s)", finalHostname)
            }

            // Browser check
            browserResult := browserCheck(finalURL)
            if browserResult {
                fmt.Printf("Active: %s (Status Code: %d)%s\n", url, response.StatusCode, redirectInfo)
                appendToFile("active_domains.txt", fmt.Sprintf("%s%s\n", url, redirectInfo))
            } else {
                appendToFile("inactive_domains.txt", domain+"\n")
            }
            isActive = true
            break
        }
    }

    if !isActive {
        appendToFile("inactive_domains.txt", domain+"\n")
    }
}

// Checks if the domain is resolvable via DNS
func isDomainResolvable(domain string) bool {
    _, err := net.LookupHost(domain)
    return err == nil
}

// Sends an HTTP GET request with retries
func tryRequest(url string) (*http.Response, error) {
    for attempt := 0; attempt < retryCount; attempt++ {
        client := &http.Client{
            Timeout: timeout,
            Transport: &http.Transport{
                TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
            },
            CheckRedirect: func(req *http.Request, via []*http.Request) error {
                if len(via) >= 10 {
                    return http.ErrUseLastResponse
                }
                return nil
            },
        }

        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            continue
        }
        req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; DomainChecker/1.0)")

        resp, err := client.Do(req)
        if err == nil {
            return resp, nil
        }
    }
    return nil, fmt.Errorf("request failed")
}

// Performs a browser-based check using chromedp
func browserCheck(url string) bool {
    opts := append(chromedp.DefaultExecAllocatorOptions[:],
        chromedp.Flag("headless", true),
        chromedp.Flag("ignore-certificate-errors", true),
        chromedp.UserAgent("Mozilla/5.0 (compatible; DomainChecker/1.0)"),
    )

    allocCtx, cancelAlloc := chromedp.NewExecAllocator(context.Background(), opts...)
    defer cancelAlloc()

    ctx, cancel := chromedp.NewContext(allocCtx)
    defer cancel()

    ctx, cancelTimeout := context.WithTimeout(ctx, timeout)
    defer cancelTimeout()

    var title string
    err := chromedp.Run(ctx,
        chromedp.Navigate(url),
        chromedp.Title(&title),
    )
    if err != nil {
        return false
    }
    // Additional page checks can be added here
    return true
}

// Appends text to a file
func appendToFile(filename, text string) {
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return
    }
    defer file.Close()
    if _, err := io.WriteString(file, text); err != nil {
        return
    }
}

// Checks if an integer slice contains a specific value
func contains(slice []int, item int) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}
