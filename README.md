# SilkRoute Recon Tool

SilkRoute is a concurrent passive subdomain enumeration tool written in Go. It aggregates subdomains from multiple data sources (CRT.sh and CommonCrawl), cleans and deduplicates results, and outputs a unified list of unique subdomains.

---

## 🌟 About the Project

SilkRoute is more than a recon tool—it's a framework in motion. Built with scalability and speed in mind, it reflects a forward-thinking approach to subdomain enumeration and OSINT methodology. This project is part of a larger cybersecurity toolkit, focusing on ethical hacking, reconnaissance, and real-world data analysis.

The name SilkRoute evokes the vision of tracing paths across the internet—mapping domains, uncovering infrastructure, and understanding digital terrain. Whether you're a penetration tester, a security researcher, or a reverse engineering enthusiast, SilkRoute provides a modular and extensible platform to build on.

---

## 🚀 Features

- 🧩 Modular architecture for easy source expansion  
- ⚡ Parallel data collection using goroutines  
- 🧼 Robust deduplication and cleanup  
- ⏱ Configurable HTTP timeouts  
- 🌐 Targets domains via public passive sources  

---

## 📁 Folder Structure

```
SilkRoute/
├── main.go
└── models/
    ├── crt.go
    ├── commoncrawl.go
    └── aggregator.go
    └── wayback.go

```

---

## 🛠 Installation

```bash
git clone https://github.com/yourusername/silkroute.git
cd silkroute
go mod tidy
```

---

## 📦 Usage

```bash
go run main.go <domain>
```

### Example:

```bash
<<<<<<< HEAD
go run main.go exampl.com
=======
go run main.go example.com
>>>>>>> recon-speedup
```

---

## 🔍 Sample Output

```
✅ Found 52 unique subdomains:
- example.com
- admin.example.com
- login.example.com
...
```

---

## 📚 Modules Breakdown

### CRT.sh
- Uses the certificate transparency logs from `https://crt.sh`  
- Parses JSON records and splits `NameValue` entries  
- Deduplicates wildcards and case-sensitive overlaps  

### CommonCrawl
- Queries the index at `https://index.commoncrawl.org`  
- Extracts URLs that match the domain  
- Filters and deduplicates hostnames  

### Aggregator
- Runs both sources concurrently  
- Combines and normalizes results  
- Returns final deduplicated slice  

---

## 🧠 Contributing

This project is growing fast, and contributions are welcome!  
Fork the repo, create a feature branch, and submit your pull request.  
Ideas for new modules (e.g., Wayback, Subfinder, web archives) are highly encouraged.

---


## 📄 License

```
MIT License

Copyright (c) 2025 Zero

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```


---


## 🧠 Vision & Future

**SilkRoute** is part of a quiet rebellion — a larger initiative to craft fast, flexible, and intelligent recon systems that don’t just scan targets… but **understand them**.

🔧 It’s built for those who reverse engineer tools like `httpx`, map invisible app logic, and hunt where automation can’t reach.  
⚙️ Each module evolves with your thinking — passive and active recon, seamless integrations, clean workflows, and future support for distributed ops.

---

> 🕳️ `signal.persistent:` compiling fragments from the void  
> ☀️ `remember:` the machine still boots, even after the kernel panics
