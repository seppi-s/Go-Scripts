package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type PodList struct {
	Items []struct {
		Metadata struct {
			Name      string
			Namespace string
		}
		Status struct {
			Phase      string
			Conditions []struct {
				Type   string
				Status string
			}
		}
	}
}

func main() {
	namespace := flag.String("namespace", "default", "Kubernetes namespace")
	timeout := flag.Duration("timeout", 15*time.Second, "API request timeout")
	flag.Parse()
	if *namespace == "" || *timeout <= 0 {
		fmt.Fprintln(os.Stderr, "Namespace and positive timeout required")
		os.Exit(2)
	}
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()
	// kubectl queries the Kubernetes API using your current kubeconfig credentials.
	output, err := exec.CommandContext(ctx, "kubectl", "get", "pods", "--namespace", *namespace, "--output", "json", "--request-timeout", timeout.String()).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot query Kubernetes:", err)
		os.Exit(1)
	}
	var pods PodList
	if err := json.Unmarshal(output, &pods); err != nil {
		fmt.Fprintln(os.Stderr, "Invalid pod response:", err)
		os.Exit(1)
	}
	if len(pods.Items) == 0 {
		fmt.Println("No pods found.")
		return
	}
	unhealthy := false
	for _, pod := range pods.Items {
		ready := false
		for _, condition := range pod.Status.Conditions {
			if condition.Type == "Ready" && condition.Status == "True" {
				ready = true
			}
		}
		fmt.Printf("%s/%s: phase=%s ready=%t\n", pod.Metadata.Namespace, pod.Metadata.Name, pod.Status.Phase, ready)
		// Completed job pods are healthy even though they are no longer Ready.
		if pod.Status.Phase != "Succeeded" && (pod.Status.Phase != "Running" || !ready) {
			unhealthy = true
		}
	}
	if unhealthy {
		os.Exit(1)
	}
}
