# AetherDNS ⚡

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen)
![License](https://img.shields.io/badge/license-MIT-blue)

AetherDNS is a blazingly fast, concurrent DNS enumeration and subdomain discovery CLI built for penetration testers and bug bounty hunters. 

By leveraging Go's lightweight goroutines, AetherDNS can resolve thousands of subdomains per second with minimal CPU and memory overhead, making it ideal for running on constrained VPS environments.

## Features
* **Massively Concurrent:** Spin up custom thread pools to saturate your network link.
* **Standard DNS Resolution:** Uses native `net.LookupIP` for accurate, OS-level cached resolutions.
* **Zero Dependencies:** Built entirely with the Go standard library. No bloated `go.mod`.

## Installation

Ensure you have Go installed, then clone the repository and build the binary:

```bash
git clone [https://github.com/YourOrg/aether-dns.git](https://github.com/YourOrg/aether-dns.git)
cd aether-dns
go build -o aether cmd/aether/main.go
sudo mv aether /usr/local/bin/
