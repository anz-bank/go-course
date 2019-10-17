# Puppy REST API Web Server

This is a REST web server for hosting puppies, the available routes for this server are:

    GET    /api/puppy/{id}
    POST   /api/puppy/          Payload: Puppy JSON without ID
    PUT    /api/puppy/{id}      Payload: Puppy JSON without ID
    DELETE /api/puppy/{id}

## How to run the puppy server?

`go run ./pkg/cmd/puppy-server/main.go --data fileName.json`

For any help with usage please use flag `--help`

#### Optional flags:

1. `--store` (long flag), `-s` (short flag)

- `map` using native map structure (default option if not specified)
- `sync`using [sync.Map](https://golang.org/pkg/sync/#Map) structure

2. `--port` (long flag), `-p` (short flag)

# Lost Puppy Web Server

This server runs on a single endpoint of:

    POST /api/lostpuppy/    Payload: { id: PUPPY_ID } 

## How to run the lost puppy server?

`go run ./pkg/cmd/lostpuppy-service/main.go`

## How to build this project?

First thing first, you must install go and golangci-lint:

1. Install [go 1.12](https://golang.org/doc/install)
2. Install [golangci-lint](https://github.com/golangci/golangci-lint#install)

To build the project, run the following command while you are in the root of the project folder:

`go build ./...`

## How to test this project?

To test this project, follow the steps below:

1. Run test cases:`go test ./...`
2. Lint it: `golangci-lint run`
3. Ensure 100% test coverage `go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out`
