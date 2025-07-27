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

}