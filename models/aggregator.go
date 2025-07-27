package models

import (
    "fmt"
    "strings"
    "sync"
)

// GetAllSubdomains merges results from CRT.sh and CommonCrawl concurrently
func GetAllSubdomains(domain string) ([]string, error) {
    var combined = sync.Map{}
    var wg sync.WaitGroup
    var errs []string
    var mu sync.Mutex

    sources := []func(string) ([]string, error){
        GetCRTSubdomains,
        GetCommonCrawlSubdomains,
    }

    for _, fn := range sources {
        wg.Add(1)
        go func(f func(string) ([]string, error)) {
            defer wg.Done()

            subs, err := f(domain)
            if err != nil {
                mu.Lock()
                errs = append(errs, err.Error())
                mu.Unlock()
                return
            }

            for _, sub := range subs {
                sub = strings.ToLower(strings.TrimSpace(sub))
                if sub != "" {
                    combined.Store(sub, struct{}{})
                }
            }
        }(fn)
    }

    wg.Wait()

    var result []string
    combined.Range(func(key, _ any) bool {
        result = append(result, key.(string))
        return true
    })

    if len(errs) > 0 {
        return result, fmt.Errorf("some errors occurred: %s", strings.Join(errs, "; "))
    }

    return result, nil
}
