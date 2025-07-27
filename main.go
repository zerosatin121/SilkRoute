package main

import(
	"fmt"
	"os"
)

func main(){
	if len(os.Args)<2{
		fmt.Println("ussage : recon-tool <domain>")
		return 
	} 

	domain := os.Args[1]

	subs , err := GetSubdomains(domain)
	if err!= nil{
		fmt.Println("❌ Error: %v\n", err)
		return
	}

	fmt.Printf("✅ Found %d unique subdomains:\n", len(subs))

	for _,sub := range subs{
		fmt.Println("-",sub)
	}
}