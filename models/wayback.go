package models

import (
    "bufio"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strings"
    "time"
)

type WaybackEntry struct {
    URL string `json:"original"`
}

func GetWaybackSubdomains(domain string) ([]string, error) {
    client := &http.Client{
        Timeout: 60 * time.Second,
    }

    query := fmt.Sprintf("https://web.archive.org/cdx/search/cdx?url=*.%s/*&output=json&collapse=urlkey&fl=original&pageSize=100", domain)
    resp, err := client.Get(query)
    if err != nil {
        return nil, fmt.Errorf("wayback http error: %w", err)
    }
    defer resp.Body.Close()

    reader := bufio.NewReader(resp.Body)
    var results []string
    var lineNum int
    seen := make(map[string]struct{})

    for {
        line, err := reader.ReadBytes('\n')
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, fmt.Errorf("wayback read error: %w", err)
        }

        lineNum++
        if lineNum == 1 {
            continue // skip header
        }

        var entry []string
        if err := json.Unmarshal(line, &entry); err != nil || len(entry) == 0 {
            continue
        }

        url := entry[0]
        url = strings.TrimPrefix(url, "http://")
        url = strings.TrimPrefix(url, "https://")
        host := strings.Split(url, "/")[0]
        host = strings.ToLower(strings.TrimSpace(strings.TrimPrefix(host, "*.")))

        if host != "" && host != domain {
            seen[host] = struct{}{}
        }
    }

    for sub := range seen {
        results = append(results, sub)
    }

    return results, nil
}
