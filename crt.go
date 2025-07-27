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
}