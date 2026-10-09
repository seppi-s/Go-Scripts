# Docker container monitor

A small beginner Go project using only the standard library.

## Requirements

Go 1.22 or newer. Install Docker CLI and start the Docker daemon. Uses your current Docker context. Reports all containers; stopped containers and Docker health status `unhealthy` produce exit code 1. Containers without health checks are judged by their running state.

## Run

From this folder:

```sh
go run .
go run . -help
```

## Build

```sh
go build -o docker-monitor .
```

## Learn from the code

Start at `main()`. Flags collect user input, then the program reads data, checks it, and prints a result. Errors are checked explicitly.

Exit code 0 means success, 1 means a failed check or operation, and 2 means invalid usage. For the long-running agent, unhealthy checks are logged and monitoring continues.
