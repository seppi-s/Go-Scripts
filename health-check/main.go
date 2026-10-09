package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	urls := flag.String("urls", "", "Comma-separated health endpoint URLs")
	timeout := flag.Duration("timeout", 5*time.Second, "Timeout for each request")
	flag.Parse()
	if strings.TrimSpace(*urls) == "" || *timeout <= 0 {
		fmt.Fprintln(os.Stderr, "Provide -urls and a positive -timeout")
		os.Exit(2)
	}
	client := &http.Client{Timeout: *timeout}
	unhealthy := false
	// Check one service at a time to keep the example easy to follow.
	for _, address := range strings.Split(*urls, ",") {
		address = strings.TrimSpace(address)
		response, err := client.Get(address)
		if err != nil {
			fmt.Printf("UNHEALTHY %s: %v\n", address, err)
			unhealthy = true
			continue
		}
		response.Body.Close()
		if response.StatusCode < 200 || response.StatusCode >= 300 {
			fmt.Printf("UNHEALTHY %s: %s\n", address, response.Status)
			unhealthy = true
		} else {
			fmt.Printf("HEALTHY %s: %s\n", address, response.Status)
		}
	}
	if unhealthy {
		os.Exit(1)
	}
}
