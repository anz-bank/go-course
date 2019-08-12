# n0npax's puppy

This project is source completed lab11 from [go course](https://github.com/anz-bank/go-course/)

## Prerequisites

-   Install `go 1.12` according to [official installation instruction](https://golang.org/doc/install)
-   Clone this project outside your `$GOPATH` to enable [Go Modules](https://github.com/golang/go/wiki/Modules)
-   Install `golangci-lint` according to [instructions](https://github.com/golangci/golangci-lint#local-installation)

## Build, execute, test, lint

Build and install this project with

    go install ./...

Run apps with:

    go run cmd/lostpuppy-service/main.go
    go run cmd/puppy-server/main.go

Or build with:

    go build -o puppy-server cmd/puppy-server/main.go
    go build -o lost-puppy-svc cmd/lostpuppy-service/main.go

Help and parameters description

    ./puppy-server --help
    ./lostpuppy-svc --help

Test it with

    go test -race -coverprofile=coverage.out -covermode=atomic ./...

Lint it with

    goimports -w .
    golangci-lint run

Review coverage with

    go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## API

### Lost puppy service
Following endpoints are supported:

        POST   /api/lostpuppy/      Payload: id

Example request:

	curl -X POST http://${IP_OR_FQDN}:${PORT}/api/lostpuppy/ -d '{"id":42}'

### Puppy server
Following endpoints are supported:

	GET    /api/puppy/{id}
	POST   /api/puppy/          Payload: Puppy JSON with or without ID
	PUT    /api/puppy/{id}      Payload: Puppy JSON with or without ID
	DELETE /api/puppy/{id}

	Puppy ID in Json will be ignored if provided as payload

Example request:

	curl -X POST http://${IP_OR_FQDN}:${PORT}/api/puppy/ -d '{"id":0,"value":4,"breed":"Type: D","colour":"Red"}'
