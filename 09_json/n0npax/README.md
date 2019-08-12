# n0npax's puppy

This project is source completed lab09 from [go course](https://github.com/anz-bank/go-course/)

## Prerequisites

-   Install `go 1.12` according to [official installation instruction](https://golang.org/doc/install)
-   Clone this project outside your `$GOPATH` to enable [Go Modules](https://github.com/golang/go/wiki/Modules)
-   Install `golangci-lint` according to [instructions](https://github.com/golangci/golangci-lint#local-installation)

## Build, execute, test, lint

Build and install this project with

    go install ./...

Build and execute its binary with

    go build -o lab09 cmd/puppy-server/main.go
    ./lab09

Help and parameters description

    ./lab09 --help

Test it with

    go test -race -coverprofile=coverage.out -covermode=atomic ./...

Lint it with

    golangci-lint run

Review coverage with

    go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
