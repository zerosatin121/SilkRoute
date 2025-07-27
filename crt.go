package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
	uniqueSubs := make(map[string] struct{})
 
	for _,entry := range entries{
		for _,line := range strings.Split(entry.NameValue, "\n"){
			 line = strings.TrimSpace(line)

            // Only add valid, non-empty lines to the map
            if line != "" {
                uniqueSubs[line] = struct{}{}
            }

		}
	

	}

	var result []string
	for sub := range uniqueSubs{
		result = append(result,sub)

	}

	return result , nil

}