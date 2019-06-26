# n0npax's puppy

This project is source completed lab10 from [go course](https://github.com/anz-bank/go-course/)

## Prerequisites

-   Install `go 1.12` according to [official installation instruction](https://golang.org/doc/install)
-   Clone this project outside your `$GOPATH` to enable [Go Modules](https://github.com/golang/go/wiki/Modules)
-   Install `golangci-lint` according to [instructions](https://github.com/golangci/golangci-lint#local-installation)

## Build, execute, test, lint

Build and install this project with

    go install ./...

Build and execute its binary with

    go build -o lab10 cmd/puppy-server/main.go
    ./lab10

Help and parameters description

    ./lab10 --help

Test it with

    go test -race -coverprofile=coverage.out -covermode=atomic ./...

Lint it with

    golangci-lint run

Review coverage with

    go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## API

Following endpoints are supported:

	GET    /api/puppy/{id}
	POST   /api/puppy/          Payload: Puppy JSON with or without ID
	PUT    /api/puppy/{id}      Payload: Puppy JSON with or without ID
	DELETE /api/puppy/{id}

	Puppy ID in Json will be ignored if provided as payload

Example request:

	curl -X POST http://${IP_OR_FQDN}:${PORT}/api/puppy/ -d '{"id":0,"value":4,"breed":"Type: D","colour":"Red"}'
