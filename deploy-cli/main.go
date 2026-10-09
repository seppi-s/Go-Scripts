package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	manifest := flag.String("file", "deployment.yaml", "Local Kubernetes manifest file")
	apply := flag.Bool("apply", false, "Actually deploy (default is a server-side dry run)")
	timeout := flag.Duration("timeout", 30*time.Second, "Deployment timeout")
	flag.Parse()
	if *timeout <= 0 {
		fmt.Fprintln(os.Stderr, "Timeout must be positive")
		os.Exit(2)
	}
	info, err := os.Stat(*manifest)
	if err != nil || !info.Mode().IsRegular() {
		fmt.Fprintln(os.Stderr, "Provide an existing manifest file with -file")
		os.Exit(2)
	}
	args := []string{"apply", "-f", *manifest, "--request-timeout", timeout.String()}
	if !*apply {
		args = append(args, "--dry-run=server")
		fmt.Println("Validating deployment with a server-side dry run.")
	} else {
		fmt.Println("Applying deployment to the current Kubernetes context.")
	}
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()
	// Pass arguments directly, without a shell. kubectl uses your current context.
	command := exec.CommandContext(ctx, "kubectl", args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if err := command.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Deployment failed:", err)
		os.Exit(1)
	}
}
