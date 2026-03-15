package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type Config struct {
	Domain      string
	Wordlist    string
	Concurrency int
	Timeout     int
}

func main() {
	banner := `
    ___       __  __               ____  _   __  _____
   /   | ___ / /_/ /_  ___  _____ / __ \/ | / / / ___/
  / /| |/ _ \ __/ __ \/ _ \/ ___// / / /  |/ /  \__ \ 
 / ___ /  __/ /_/ / / /  __/ /   / /_/ / /|  /  ___/ / 
/_/  |_\___/\__/_/ /_/\___/_/   /_____/_/ |_/  /____/  v1.2.0
	`
	fmt.Println(banner)

	cfg := Config{}
	flag.StringVar(&cfg.Domain, "d", "", "Target domain (e.g., example.com)")
	flag.StringVar(&cfg.Wordlist, "w", "", "Path to subdomain wordlist")
	flag.IntVar(&cfg.Concurrency, "t", 50, "Number of concurrent threads")
	flag.IntVar(&cfg.Timeout, "timeout", 5, "DNS resolution timeout in seconds")
	flag.Parse()

	if cfg.Domain == "" || cfg.Wordlist == "" {
		fmt.Println("[-] Error: Domain (-d) and Wordlist (-w) are required.")
		os.Exit(1)
	}

	file, err := os.Open(cfg.Wordlist)
	if err != nil {
		fmt.Printf("[-] Failed to open wordlist: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	jobs := make(chan string, cfg.Concurrency)
	results := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < cfg.Concurrency; i++ {
		wg.Add(1)
		go worker(cfg.Domain, jobs, results, &wg)
	}

	go func() {
		for res := range results {
			fmt.Println(res)
		}
	}()

	scanner := bufio.NewScanner(file)
	start := time.Now()
	for scanner.Scan() {
		jobs <- scanner.Text()
	}
	close(jobs)

	wg.Wait()
	close(results)
	fmt.Printf("\n[*] Enumeration completed in %v\n", time.Since(start))
}

func worker(baseDomain string, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for sub := range jobs {
		sub = strings.TrimSpace(sub)
		if sub == "" {
			continue
		}

		target := fmt.Sprintf("%s.%s", sub, baseDomain)
		ips, err := net.LookupIP(target)

		if err == nil && len(ips) > 0 {
			results <- fmt.Sprintf("[+] %-30s : %s", target, ips[0].String())
		}
	}
}
