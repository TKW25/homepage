# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Stack

Go + [HTMX](https://htmx.org/) + [Templ](https://templ.guide/) (server-side HTML templating) + Tailwind CSS. The module name is `app` (`go.mod`). Node.js tooling is available for frontend assets.

## Commands

```bash
# Live-reload development server (uses air)
air

# Build
go build -o ./tmp/main .

# Run (after build)
./tmp/main

# Generate templ files (required before build when .templ files change)
templ generate

# Lint
golangci-lint run

# Test
go test ./...

# Single test
go test ./path/to/package -run TestFunctionName
```

## Development Environment

This project uses a devcontainer (`.devcontainer/`). The container includes: `air` (live reload), `templ`, `golangci-lint`, `dlv` (debugger), and Node.js 22. Ports 8080 (HTTP) and 8443 (HTTPS) are forwarded.

A PostgreSQL service is pre-configured but commented out in `.devcontainer/docker-compose.yml` — uncomment the `db` service and `depends_on` when a database is needed.

## Templ Workflow

`.templ` files are compiled to `.go` files by `templ generate`. The generated `*_templ.go` files are committed. Always run `templ generate` before `go build` when `.templ` files have changed. `air` watches `.templ` files and handles this automatically during development.
