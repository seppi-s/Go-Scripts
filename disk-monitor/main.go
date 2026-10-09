package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	path := flag.String("path", "/", "Directory on the filesystem to check")
	threshold := flag.Int("threshold", 80, "Alert when used space reaches this percentage")
	flag.Parse()
	if *threshold < 1 || *threshold > 100 {
		fmt.Fprintln(os.Stderr, "Threshold must be 1 through 100")
		os.Exit(2)
	}
	directory, err := filepath.Abs(*path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	absolute, err := os.Stat(directory)
	if err != nil || !absolute.IsDir() {
		fmt.Fprintln(os.Stderr, "Path must be an existing directory")
		os.Exit(2)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// df -P prints a portable table on Linux and macOS.
	output, err := exec.CommandContext(ctx, "df", "-P", directory).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot read disk usage:", err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) < 2 {
		fmt.Fprintln(os.Stderr, "Unexpected df output")
		os.Exit(1)
	}
	fields := strings.Fields(lines[len(lines)-1])
	if len(fields) < 6 {
		fmt.Fprintln(os.Stderr, "Unexpected df row")
		os.Exit(1)
	}
	used, err := strconv.Atoi(strings.TrimSuffix(fields[4], "%"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid disk percentage:", err)
		os.Exit(1)
	}
	fmt.Printf("Disk %s: %d%% used (threshold %d%%)\n", *path, used, *threshold)
	if used >= *threshold {
		fmt.Println("ALERT: disk space is low")
		os.Exit(1)
	}
}
