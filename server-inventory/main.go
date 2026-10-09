package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Inventory struct {
	Hostname     string `json:"hostname"`
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
	CPUs         int    `json:"logical_cpus"`
	CPUModel     string `json:"cpu_model"`
	RAMBytes     uint64 `json:"ram_bytes"`
}

func main() {
	// /proc contains host information on Linux. Fail clearly on other systems.
	if runtime.GOOS != "linux" {
		fmt.Fprintln(os.Stderr, "This example requires Linux /proc files")
		os.Exit(1)
	}
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	result := Inventory{Hostname: hostname, OS: runtime.GOOS, Architecture: runtime.GOARCH, CPUs: runtime.NumCPU(), CPUModel: "unknown"}
	cpu, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, line := range strings.Split(string(cpu), "\n") {
		key, value, ok := strings.Cut(line, ":")
		if ok && strings.TrimSpace(key) == "model name" {
			result.CPUModel = strings.TrimSpace(value)
			break
		}
	}
	memory, err := os.Open("/proc/meminfo")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer memory.Close()
	scanner := bufio.NewScanner(memory)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) >= 2 && fields[0] == "MemTotal:" {
			kb, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			result.RAMBytes = kb * 1024
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if result.RAMBytes == 0 {
		fmt.Fprintln(os.Stderr, "MemTotal was not found")
		os.Exit(1)
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(result); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
