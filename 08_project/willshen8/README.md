# Puppy Storer - A database for puppies

This project gives you experience in the following areas in Golang:
1. struct
2. types
3. interfaces
4. implementing interfaces with different implementations
5. Concurrency and race-conditions
6. Error handling
7. Test coverage using [test suite](https://godoc.org/github.com/stretchr/testify/suite)

### Overall overview:
This simple program maintains a list of puppies. 

A puppy has the following attributes:
- ID
- Breed
- Colour 
- Value 

It has a `storer` interface which has the following CRUD methods for puppies:
- CreatePuppy
- ReadPuppy
- UpdatePuppy
- DeletePuppy

The storer interface has two different implementations:
- `MapStore` using native map structure
- `SyncStore` using [sync.Map](https://golang.org/pkg/sync/#Map) structure

## How to build this project?
First thing first, you must install go and golangci-lint:
1. Install [go 1.12](https://golang.org/doc/install)
2. Install [golangci-lint](https://github.com/golangci/golangci-lint#install)

To build the project, run the following command while you are in the root of the project folder:

`go build ./...`

## How to run this project?
`go run ./pkg/cmd/puppy-server/main.go`

## How to test this project?

To test this project, follow the steps below:

1. Run test cases:`go test ./...`
2. Lint it: `golangci-lint run`
3. Ensure 100% test coverage `go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out`

