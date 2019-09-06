# Project: Building Incrementally Towards a Puppy Rest API

This project is part of the [go course](https://github.com/anz-bank/go-course/) run by ANZ Bank.

It builds you up incrementally with the ultimate end goal of having a fully functioning Rest API.

It aims to impart several key learnings such as but not limited to:

- Working with custom types in go using structs
- Understanding and working with interfaces in go
- Implementing multiple implementations of an interface
- Go standard project layout
- Best practices with Git. Proper commit messages and PRs. Git rebasing when necessary to keep commit history clean etc.
- Creating custom error types to suit your project needs
- Thread safety using sync.Mutex
- Test coverage using [test suites](https://godoc.org/github.com/stretchr/testify/suite)
- Implement ability to convert between Go in memory objects to JSON representation and vice versa
- Implement ability to accept user input from CLI and read data from a file
- Load data from json text file and store in an instance of storer

## Prerequisites

-   Install `go 1.12` according to [official installation instruction](https://golang.org/doc/install)
-   Clone this project outside your `$GOPATH` to enable [Go Modules](https://github.com/golang/go/wiki/Modules)
-   Install `golangci-lint` according to [instructions](https://github.com/golangci/golangci-lint#local-installation)

## How to build this project?

To build the project, run the following command while you are in the root of the project folder:

`go build ./...`

## How to run this project?
`go run ./cmd/puppy-server/main.go --data puppy-data/puppies.json`

## How to test this project?

To test this project, follow the steps below:

1. Run test cases:`go test ./...`
2. Lint it: `golangci-lint run`
3. Ensure 100% test coverage `go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out`