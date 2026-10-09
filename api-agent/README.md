# Infrastructure/API monitoring agent

A small beginner Go project using only the standard library.

## Requirements

Go 1.22 or newer. Any OS supporting Go. Prints a check immediately and then repeatedly. Run the built executable in the background using `./api-agent -url http://localhost:8080/health > agent.log 2>&1 &` on Linux/macOS; stop it with SIGTERM. This is a learning example, not an installed system service.

## Run

From this folder:

```sh
go run . -url http://localhost:8080/health -interval 30s
go run . -help
```

## Build

```sh
go build -o api-agent .
```

## Learn from the code

Start at `main()`. Flags collect user input, then the program reads data, checks it, and prints a result. Errors are checked explicitly.

Exit code 0 means success, 1 means a failed check or operation, and 2 means invalid usage. For the long-running agent, unhealthy checks are logged and monitoring continues.
