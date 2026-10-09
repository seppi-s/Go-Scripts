package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func check(ctx context.Context, client *http.Client, address string) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, address, nil)
	if err != nil {
		fmt.Println("Invalid request:", err)
		return
	}
	response, err := client.Do(request)
	stamp := time.Now().Format(time.RFC3339)
	if err != nil {
		fmt.Printf("%s UNHEALTHY: %v\n", stamp, err)
		return
	}
	response.Body.Close()
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		fmt.Printf("%s HEALTHY: %s\n", stamp, response.Status)
	} else {
		fmt.Printf("%s UNHEALTHY: %s\n", stamp, response.Status)
	}
}

func main() {
	address := flag.String("url", "", "Health endpoint URL")
	interval := flag.Duration("interval", 30*time.Second, "Delay between checks")
	timeout := flag.Duration("timeout", 5*time.Second, "HTTP timeout")
	flag.Parse()
	if *address == "" || *interval <= 0 || *timeout <= 0 {
		fmt.Fprintln(os.Stderr, "Provide -url and positive interval/timeout")
		os.Exit(2)
	}
	request, err := http.NewRequest(http.MethodGet, *address, nil)
	if err != nil || request.URL.Host == "" || (request.URL.Scheme != "http" && request.URL.Scheme != "https") {
		fmt.Fprintln(os.Stderr, "URL must use http or https and include a host")
		os.Exit(2)
	}
	// Ctrl+C or SIGTERM stops the loop and cancels an active HTTP request.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	client := &http.Client{Timeout: *timeout}
	check(ctx, client, *address)
	ticker := time.NewTicker(*interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Agent stopped.")
			return
		case <-ticker.C:
			check(ctx, client, *address)
		}
	}
}
