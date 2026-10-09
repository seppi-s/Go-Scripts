# CI/CD deployment CLI

A small beginner Go project using only the standard library.

## Requirements

Go 1.22 or newer. Install kubectl and configure a cluster context. Default mode performs a server-side dry run, which still requires cluster access. Add `-apply` to perform the real deployment. Check `kubectl config current-context` first. This command applies a manifest; it does not wait for rollout completion.

## Run

From this folder:

```sh
go run . -file deployment.yaml
go run . -help
```

## Build

```sh
go build -o deploy-cli .
```

## Learn from the code

Start at `main()`. Flags collect user input, then the program reads data, checks it, and prints a result. Errors are checked explicitly.

Exit code 0 means success, 1 means a failed check or operation, and 2 means invalid usage. For the long-running agent, unhealthy checks are logged and monitoring continues.
