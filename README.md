# Go Scripts

Eight small DevOps tools written in beginner-friendly Go. Each folder is an independent project with its own `go.mod`, `main.go`, and README. All Go code uses the standard library. Docker and Kubernetes tools call their installed CLIs.

| Project | What it does |
| --- | --- |
| [health-check](health-check/) | Health-check multiple services |
| [docker-monitor](docker-monitor/) | Docker container monitor |
| [log-analyzer](log-analyzer/) | Log analyzer |
| [disk-monitor](disk-monitor/) | Disk-space monitor |
| [pod-checker](pod-checker/) | Kubernetes pod checker |
| [server-inventory](server-inventory/) | Server inventory |
| [deploy-cli](deploy-cli/) | CI/CD deployment CLI |
| [api-agent](api-agent/) | Infrastructure/API monitoring agent |

## Getting started

Install Go 1.22 or newer, clone this repository, and enter a project folder:

```sh
git clone https://github.com/seppi-s/Go-Scripts.git
cd Go-Scripts/health-check
go run . -urls http://localhost:8080/health
```

Each project README explains requirements and flags. Replace sample URLs and configuration with your own. These are learning examples with terminal alerts; no email or notification service is configured.

## Separate repositories

Each project is self-contained and can be copied into its own Git repository. Update the module path in its `go.mod` to match the new repository URL.

## Verify all projects

```sh
for project in */; do
  (cd "$project" && go test ./... && go vet ./... && go build -o /tmp/go-script-check .) || exit 1
done
```
