package models

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strings"
)

// Entry represents a single Common Crawl JSON record
type Entry struct {
    URL string `json:"url"`
}

// GetSubdomains queries Common Crawl for domain-related URLs and returns subdomains
func GetCommonCrawlSubdomains(domain string) ([]string, error) {
    index := "CC-MAIN-2024-10-index" // You can rotate this manually or make it dynamic
    query := fmt.Sprintf("https://index.commoncrawl.org/%s?url=*.%s&output=json", index, domain)

    resp, err := http.Get(query)
    if err != nil {
        return nil, fmt.Errorf("http error: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("read error: %w", err)
    }

    lines := strings.Split(string(body), "\n")
    uniqueSubs := make(map[string]struct{})

    for _, line := range lines {
        line = strings.TrimSpace(line)
        if line == "" {
            continue
        }

        var entry Entry
        if err := json.Unmarshal([]byte(line), &entry); err != nil {
            continue // skip broken lines
        }

        url := entry.URL
        // Extract hostname only (without scheme and path)
        url = strings.TrimPrefix(url, "http://")
        url = strings.TrimPrefix(url, "https://")
        url = strings.Split(url, "/")[0]

        // Filter wildcards and base domain
        if url != "" && url != domain && !strings.HasPrefix(url, "*.") {
            uniqueSubs[url] = struct{}{}
        }
    }

    var results []string
    for sub := range uniqueSubs {
        results = append(results, sub)
    }

    return results, nil
}
