# Disk-space monitor

A small beginner Go project using only the standard library.

## Requirements

Go 1.22 or newer. Linux or macOS with `df` in PATH. Alerts are printed to the terminal. Use an absolute directory path.

## Run

From this folder:

```sh
go run . -path / -threshold 80
go run . -help
```

## Build

```sh
go build -o disk-monitor .
```

## Learn from the code

Start at `main()`. Flags collect user input, then the program reads data, checks it, and prints a result. Errors are checked explicitly.

Exit code 0 means success, 1 means a failed check or operation, and 2 means invalid usage. For the long-running agent, unhealthy checks are logged and monitoring continues.
