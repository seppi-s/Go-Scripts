# Kubernetes pod checker

A small beginner Go project using only the standard library.

## Requirements

Go 1.22 or newer. Install kubectl and configure kubeconfig access to a cluster. Uses the current context and requires permission to list pods in the chosen namespace. Queries the API through kubectl. Running pods must be Ready; Succeeded job pods are accepted. No pods is an informational result.

## Run

From this folder:

```sh
go run . -namespace default
go run . -help
```

## Build

```sh
go build -o pod-checker .
```

## Learn from the code

Start at `main()`. Flags collect user input, then the program reads data, checks it, and prints a result. Errors are checked explicitly.

Exit code 0 means success, 1 means a failed check or operation, and 2 means invalid usage. For the long-running agent, unhealthy checks are logged and monitoring continues.
