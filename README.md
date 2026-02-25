# Sample Go Application — Task Manager

A simple, thread-safe task manager library demonstrating Go best practices including unit testing, table-driven tests, concurrency safety, and linting.

## Project Structure

```
sample-go-app/
├── main.go                      # CLI entry point
├── go.mod                       # Go module definition
├── .golangci.yml                # golangci-lint configuration
├── README.md
└── taskmanager/
    ├── task.go                  # Task, Priority, Status types
    ├── task_test.go             # Tests for task types
    ├── manager.go               # Manager with CRUD operations
    └── manager_test.go          # Tests for Manager
```

## Running

```bash
go run .
```

## Testing

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run with race detector
go test -race ./...

# Run with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Recommended Go Linters

### 1. [golangci-lint](https://golangci-lint.run/) (Meta-linter — **Start Here**)

The most popular Go linter aggregator. Runs 50+ linters in parallel. This project includes a `.golangci.yml` config.

```bash
# Install
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
# Or via Homebrew
brew install golangci-lint

# Run
golangci-lint run ./...
```

### 2. [staticcheck](https://staticcheck.dev/)

Advanced static analysis with checks for bugs, performance, and style.

```bash
go install honnef.co/go/tools/cmd/staticcheck@latest
staticcheck ./...
```

### 3. [go vet](https://pkg.go.dev/cmd/vet)

Built into the Go toolchain. Catches common mistakes like printf format mismatches and unreachable code.

```bash
go vet ./...
```

### 4. [errcheck](https://github.com/kisielk/errcheck)

Ensures error return values are always checked.

```bash
go install github.com/kisielk/errcheck@latest
errcheck ./...
```

### 5. [gosec](https://github.com/securego/gosec)

Security-focused linter that finds common vulnerabilities (SQL injection, hardcoded credentials, etc.).

```bash
go install github.com/securego/gosec/v2/cmd/gosec@latest
gosec ./...
```

### 6. [revive](https://github.com/mgechev/revive)

Drop-in replacement for the deprecated `golint`, with configurable rules.

```bash
go install github.com/mgechev/revive@latest
revive ./...
```

### 7. [gofumpt](https://github.com/mvdan/gofumpt)

Stricter version of `gofmt` that enforces additional formatting rules.

```bash
go install mvdan.cc/gofumpt@latest
gofumpt -w .
```

### 8. [gocritic](https://github.com/go-critic/go-critic)

Opinionated meta-linter with checks for style, performance, and common bugs.

```bash
go install github.com/go-critic/go-critic/cmd/gocritic@latest
gocritic check ./...
```

## Quick Start for Linting

The easiest approach is to just use **golangci-lint** which bundles most of the above:

```bash
brew install golangci-lint
golangci-lint run ./...
```

The `.golangci.yml` in this project already enables a curated set of linters.
