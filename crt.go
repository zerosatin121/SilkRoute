package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

type CertificateEntry struct {
	NameValue string `json:"name_value"`
}

func GetSubdomains(domain string)([]string, error){
	url := fmt.Sprintf("https://crt.sh/json?q=%s",domain)

	client := &http.Client{}

	resp, err := client.Get(url)
	if err !=nil{
		return nil,fmt.Errorf("http error: %w" , err)
	} 
	// close response body to avoid leakage.....
	defer resp.Body.Close()

	body , err := io.ReadAll(resp.Body)
	if err != nil {
		return nil ,fmt.Errorf("read  error :%w" , err)
	}

	var entries []CertificateEntry

	if err := json.Unmarshal(body, &entries); err != nil {
		return  nil, fmt.Errorf("unmarsha error : %w" ,err)
	}

	// remove duplicates 
	// uniqueSubs := make(map[string] struct{})
 
	// for _,entry := range entries{
	// 	for _,line := range strings.Split(entry.NameValue, "\n"){
	// 		 line = strings.TrimSpace(line)

    //         // Only add valid, non-empty lines to the map
    //         if line != "" {
    //             uniqueSubs[line] = struct{}{}
    //         }

	// 	}
	

	// }

	// var result []string
	// for sub := range uniqueSubs{
	// 	result = append(result,sub)

	// }

	// return result , nil

	uniqueSubs := sync.Map{}
	var wg  sync.WaitGroup

		for _, entry := range entries{
			wg.Add(1)
			go func (entry CertificateEntry){
				defer wg.Done()
				 for _, line := range strings.Split(entry.NameValue, "\n") {
                line = strings.TrimSpace(line)
                line = strings.ToLower(line) // normalize
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