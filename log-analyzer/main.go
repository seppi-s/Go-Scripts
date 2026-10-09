package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	path := flag.String("file", "sample.log", "Log file to read")
	words := flag.String("keywords", "error,fatal,panic", "Comma-separated, case-insensitive keywords")
	flag.Parse()
	keywords := []string{}
	for _, word := range strings.Split(*words, ",") {
		word = strings.ToLower(strings.TrimSpace(word))
		if word != "" {
			keywords = append(keywords, word)
		}
	}
	if len(keywords) == 0 {
		fmt.Fprintln(os.Stderr, "Provide at least one keyword")
		os.Exit(2)
	}
	file, err := os.Open(*path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// Allow lines up to 1 MiB; report larger lines instead of silently skipping them.
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)
	lineNumber, matches := 0, 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		for _, word := range keywords {
			if strings.Contains(strings.ToLower(line), word) {
				fmt.Printf("%d: %s\n", lineNumber, line)
				matches++
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Read failed:", err)
		os.Exit(1)
	}
	fmt.Printf("Found %d matching lines out of %d.\n", matches, lineNumber)
	if matches > 0 {
		os.Exit(1)
	}
}
