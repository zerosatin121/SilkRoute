package models

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strings"
    "sync"
    "time"
)

type CertificateEntry struct {
    NameValue string `json:"name_value"`
}

func GetCRTSubdomains(domain string) ([]string, error) {
    url := fmt.Sprintf("https://crt.sh/json?q=%s", domain)

    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    resp, err := client.Get(url)
    if err != nil {
        return nil, fmt.Errorf("crt http error: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("crt read error: %w", err)
    }

    var entries []CertificateEntry
    if err := json.Unmarshal(body, &entries); err != nil {
        return nil, fmt.Errorf("crt unmarshal error: %w", err)
    }

    uniqueSubs := sync.Map{}
    var wg sync.WaitGroup

    for _, entry := range entries {
        wg.Add(1)
        go func(e CertificateEntry) {
            defer wg.Done()
            for _, line := range strings.Split(e.NameValue, "\n") {
                line = strings.ToLower(strings.TrimSpace(line))
                if line != "" && !strings.HasPrefix(line, "*.") {
                    uniqueSubs.Store(line, struct{}{})
                }
            }
        }(entry)
    }
    wg.Wait()

    var result []string
    uniqueSubs.Range(func(key, _ any) bool {
        result = append(result, key.(string))
        return true
    })

    return result, nil
}
