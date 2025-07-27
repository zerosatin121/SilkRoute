package models

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "time"
)

type CRTResult struct {
    NameValue string `json:"name_value"`
}

func GetCRTSubdomains(domain string) ([]string, error) {
    client := &http.Client{
        Timeout: 60 * time.Second, // bumped up from 10s for reliability
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

    contentType := resp.Header.Get("Content-Type")
    if contentType != "application/json" {
        body, _ := io.ReadAll(resp.Body)
        log.Printf("CRT response is not JSON (got %s):\n%s", contentType, body)
        return nil, fmt.Errorf("unexpected content type from crt.sh")
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("CRT body read error: %v", err)
    }

    var results []CRTResult
    if err := json.Unmarshal(body, &results); err != nil {
        log.Printf("CRT unmarshal failed. Raw response:\n%s", body)
        return nil, fmt.Errorf("CRT unmarshal error: %v", err)
    }

    subdomains := make(map[string]bool)
    for _, r := range results {
        for _, sub := range splitAndTrim(r.NameValue) {
            subdomains[sub] = true
        }
    }

    uniqueSubs := make([]string, 0, len(subdomains))
    for sub := range subdomains {
        uniqueSubs = append(uniqueSubs, sub)
    }

    return uniqueSubs, nil
}

func splitAndTrim(input string) []string {
    // Implement logic to clean up and deduplicate names
    // This could be as simple or complex as needed
    return []string{input}
}
