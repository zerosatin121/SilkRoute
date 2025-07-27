package main

import(
	"fmt"
	"os"
	"SilkRoute/models"
)


func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: recon-tool <domain>")
        return
    }

    domain := os.Args[1]

    // 🔎 Fetch from crt.sh
  subs, err := models.GetAllSubdomains(domain)
if err != nil {
    fmt.Printf("❌ Error: %v\n", err)
    return
}

fmt.Printf("✅ Found %d unique subdomains:\n", len(subs))
for _, sub := range subs {
    fmt.Println("-", sub)
}

}
