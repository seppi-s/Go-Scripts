# Health-check multiple services

A small beginner Go project using only the standard library.

## Requirements

Go 1.22 or newer. Any OS supporting Go. A 2xx response is healthy; network errors and other response codes are unhealthy. Checks are sequential. An endpoint should return a meaningful health status.

## Run

From this folder:

```sh
go run . -urls http://localhost:8080/health,http://localhost:3000/health
go run . -help
```

## Build

```sh
go build -o health-check .
```

## Learn from the code

Start at `main()`. Flags collect user input, then the program reads data, checks it, and prints a result. Errors are checked explicitly.

Exit code 0 means success, 1 means a failed check or operation, and 2 means invalid usage. For the long-running agent, unhealthy checks are logged and monitoring continues.
