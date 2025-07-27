package models

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "strings"
    "time"
)

type CRTResult struct {
    NameValue string `json:"name_value"`
}

func GetCRTSubdomains(domain string) ([]string, error) {
    client := &http.Client{
        Timeout: 60 * time.Second,
    }

    url := fmt.Sprintf("https://crt.sh/?q=%s&output=json", domain)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("CRT request error: %v", err)
    }

    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("CRT response error: %v", err)
    }
    defer resp.Body.Close()

    if !strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
        body, _ := io.ReadAll(resp.Body)
        log.Printf("CRT non-JSON response:\n%s", body)
        return nil, fmt.Errorf("Unexpected content type from crt.sh")
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("CRT body read error: %v", err)
    }

    var results []CRTResult
    if err := json.Unmarshal(body, &results); err != nil {
        log.Printf("CRT unmarshal failed:\n%s", body)
        return nil, fmt.Errorf("CRT unmarshal error: %v", err)
    }

    seen := make(map[string]struct{})
    for _, r := range results {
        subs := splitAndClean(r.NameValue)
        for _, sub := range subs {
            seen[sub] = struct{}{}
        }
    }

    var unique []string
    for sub := range seen {
        unique = append(unique, sub)
    }

    return unique, nil
}

func splitAndClean(input string) []string {
    parts := strings.FieldsFunc(input, func(r rune) bool {
        return r == '\n' || r == ',' || r == ' '
    })

    var subs []string
    for _, part := range parts {
        sub := strings.ToLower(strings.TrimSpace(part))
        sub = strings.TrimPrefix(sub, "*.")
        if sub != "" {
            subs = append(subs, sub)
        }
    }
    return subs
}
