# Log analyzer

A small beginner Go project using only the standard library.

## Requirements

Go 1.22 or newer. Any OS supporting Go. Matches substrings without regard to letter case, once per matching line. Supports lines up to 1 MiB. Exit code 1 also indicates matching errors were found.

## Run

From this folder:

```sh
go run . -file sample.log
go run . -help
```

## Build

```sh
go build -o log-analyzer .
```

## Learn from the code

Start at `main()`. Flags collect user input, then the program reads data, checks it, and prints a result. Errors are checked explicitly.

Exit code 0 means success, 1 means a failed check or operation, and 2 means invalid usage. For the long-running agent, unhealthy checks are logged and monitoring continues.
