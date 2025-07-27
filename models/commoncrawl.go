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

type CrawlEntry struct {
    URL string `json:"url"`
}

func GetCommonCrawlSubdomains(domain string) ([]string, error) {
    index := "CC-MAIN-2024-10-index"
    query := fmt.Sprintf("https://index.commoncrawl.org/%s?url=*.%s&output=json", index, domain)

    client := &http.Client{
        Timeout: 15 * time.Second,
    }

    resp, err := client.Get(query)
    if err != nil {
        return nil, fmt.Errorf("cc http error: %w", err)
    }
    defer resp.Body.Close()

    reader := bufio.NewReader(resp.Body)
    uniqueSubs := make(map[string]struct{})

    for {
        line, err := reader.ReadBytes('\n')
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, fmt.Errorf("cc read line error: %w", err)
        }

        var entry CrawlEntry
        if err := json.Unmarshal(line, &entry); err != nil {
            continue
        }

        url := entry.URL
        url = strings.TrimPrefix(url, "http://")
        url = strings.TrimPrefix(url, "https://")
        url = strings.Split(url, "/")[0]

        if url != "" && url != domain && !strings.HasPrefix(url, "*.") {
            uniqueSubs[strings.ToLower(url)] = struct{}{}
        }
    }

    var results []string
    for sub := range uniqueSubs {
        results = append(results, sub)
    }

    return results, nil
}
