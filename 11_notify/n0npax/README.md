# n0npax's puppy

This project is source completed lab11 from [go course](https://github.com/anz-bank/go-course/)

## Prerequisites

-   Install `go 1.12` according to [official installation instruction](https://golang.org/doc/install)
-   Clone this project outside your `$GOPATH` to enable [Go Modules](https://github.com/golang/go/wiki/Modules)
-   Install `golangci-lint` according to [instructions](https://github.com/golangci/golangci-lint#local-installation)

## Build, execute, test, lint

Build and install this project with
```bash
go install ./...
```
Run apps with:
```bash
go run cmd/lostpuppy-service/main.go
go run cmd/puppy-server/main.go
```
Or build with:
```bash
go build -o puppy-server cmd/puppy-server/main.go
go build -o lost-puppy-svc cmd/lostpuppy-service/main.go
```
Help and parameters description
```bash
./puppy-server --help
./lostpuppy-svc --help
```
Test it with
```bash
go test -race -coverprofile=coverage.out -covermode=atomic ./...
```
Lint it with
```bash
goimports -w .
golangci-lint run
gosec .
```
Review coverage with
```bash
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
```
## API

### Lost puppy service
Following endpoints are supported:

        POST   /api/lostpuppy/      Payload: id

Example request:
```bash
export HOST=localhost PORT=8182
curl -X POST http://${HOST}:${PORT}/api/lostpuppy/ -d '{"id":42}'
```
### Puppy server
Following endpoints are supported:

	GET    /api/puppy/{id}
	POST   /api/puppy/          Payload: Puppy JSON with or without ID
	PUT    /api/puppy/{id}      Payload: Puppy JSON with or without ID
	DELETE /api/puppy/{id}

	Puppy ID in Json will be ignored if provided as payload

Example request:
```bash
export HOST=localhost PORT=8181
curl -X POST http://${HOST}:${PORT}/api/puppy/ -d '{"value":1410,"breed":"Type: D","colour":"White"}'
curl -X PUT http://${HOST}:${PORT}/api/puppy/42 -d '{"value":71,"breed":"Type: G","colour":"Red"}'
curl -X DELETE http://${HOST}:${PORT}/api/puppy/42
curl http://${HOST}:${PORT}/api/puppy/42
```
