# Server inventory

A small beginner Go project using only the standard library.

## Requirements

Go 1.22 or newer. Linux with /proc mounted. RAM is total host RAM, not Go process memory or container memory limits. CPU count is the logical CPU count available to Go.

## Run

From this folder:

```sh
go run .
go run . -help
```

## Build

```sh
go build -o server-inventory .
```

## Learn from the code

Start at `main()`. Flags collect user input, then the program reads data, checks it, and prints a result. Errors are checked explicitly.

Exit code 0 means success, 1 means a failed check or operation, and 2 means invalid usage. For the long-running agent, unhealthy checks are logged and monitoring continues.
