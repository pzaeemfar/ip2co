package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/pzaeemfar/oip2co/geoip"
)

func parseInput(input string) string {
	if strings.Contains(input, "://") {
		u, err := url.Parse(input)
		if err == nil && u.Host != "" {
			host := u.Host
			if strings.Contains(host, ":") {
				host, _, _ = net.SplitHostPort(host)
			}
			return host
		}
	}
	return input
}

func main() {
	debug := flag.Bool("debug", false, "Enable debug output")
	jsonOut := flag.Bool("json", false, "Output results as JSON")
	flag.Parse()

	stat, _ := os.Stdin.Stat()
	var inputs []string
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" {
				inputs = append(inputs, line)
			}
		}
	} else {
		inputs = flag.Args()
		if len(inputs) == 0 {
			flag.Usage()
			return
		}
	}

	results := make(map[string]string)
	var mu sync.Mutex

	inputCh := make(chan string)
	var wg sync.WaitGroup
	workerCount := 50

	worker := func() {
		defer wg.Done()
		for input := range inputCh {
			host := parseInput(input)

			ip := net.ParseIP(host)
			if ip == nil {
				if *debug {
					fmt.Fprintf(os.Stderr, "Skipping domain (not IP): %s\n", input)
				}
				continue
			}

			country, err := geoip.GetCountry(ip.String(), *debug)
			if err != nil {
				if *debug {
					fmt.Fprintf(os.Stderr, "Lookup failed for IP %s: %v\n", ip, err)
				}
				mu.Lock()
				results[input] = fmt.Sprintf("Lookup failed: %v", err)
				mu.Unlock()
				continue
			}

			mu.Lock()
			results[input] = country
			mu.Unlock()

			if !*jsonOut {
				fmt.Printf("%s - %s\n", input, country)
			}
		}
	}

	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go worker()
	}

	for _, input := range inputs {
		inputCh <- input
	}
	close(inputCh)

	wg.Wait()

	if *jsonOut {
		out, err := json.Marshal(results)
		if err != nil {
			if *debug {
				fmt.Fprintf(os.Stderr, "JSON marshal error: %v\n", err)
			}
		} else {
			fmt.Println(string(out))
		}
	}
}
