package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Container struct {
	Names  string
	State  string
	Status string
}

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "Docker command timeout")
	flag.Parse()
	if *timeout <= 0 {
		fmt.Fprintln(os.Stderr, "Timeout must be positive")
		os.Exit(2)
	}
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()
	// The Docker CLI handles local socket access and your configured Docker context.
	output, err := exec.CommandContext(ctx, "docker", "ps", "--all", "--format", "{{json .}}").Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot query Docker:", err)
		os.Exit(1)
	}
	if strings.TrimSpace(string(output)) == "" {
		fmt.Println("No containers found.")
		return
	}
	unhealthy := false
	for _, line := range strings.Split(strings.TrimSpace(string(output)), "\n") {
		var container Container
		if err := json.Unmarshal([]byte(line), &container); err != nil {
			fmt.Fprintln(os.Stderr, "Invalid Docker output:", err)
			os.Exit(1)
		}
		fmt.Printf("%s: %s (%s)\n", container.Names, container.State, container.Status)
		if container.State != "running" || strings.Contains(container.Status, "(unhealthy)") {
			unhealthy = true
		}
	}
	if unhealthy {
		os.Exit(1)
	}
}
