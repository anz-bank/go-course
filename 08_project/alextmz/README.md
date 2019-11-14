# Alex 'Puppy Project'

This is part of the [Go](https://golang.org/) [Course](https://github.com/anz-bank/go-course) done by [ANZ Bank](https://www.anz.com.au) during 2019.

It contains code to do basic CRUD operations (Create, Read, Update, Delete) on a "Puppy" type defined by the user, with 2 in-memory storage implementations to back it using native Go maps and [sync.map](https://golang.org/pkg/sync/).

As it is, this code is not very useful to you, who is reading this now, as it does not 'do' anything. Its purpose is to teach Golang coding skills and good development practices. 

The project is organized as follows:
- `cmd/puppy-server` contains the code for the executable. It is very crude, just barely enough to demonstrate the internal working of the package.
- `pkg/puppy` contains the type definitions, interface and error values and tests for the Puppy package - very little working code here too.
- `pkg/puppy/store` is where the action is, and contains the store backings and tests.

## Prerequisites
-   Install [`go`](https://golang.org/doc/install) and alternatively [`golangci-lint`](https://github.com/golangci/golangci-lint#local-installation) if you want to run tests or lint
-   Clone this project outside your `$GOPATH` to enable [Go Modules](https://github.com/golang/go/wiki/Modules)

## Build, install, execute
---
### Short version
For the anxious, you can just run the main executable quickly doing a

    go run ./cmd/puppy-server/main.go
    
As promised, the output is not interesting at all.

### Long version
Alternatively, you can build, install and run from your `$GOPATH` with

    go install ./...
    puppy-server

Or yet build and run from the the same directory as this `README` with

    go build -o puppy-server cmd/puppy-server/main.go
    ./puppy-server

#### Lint, test, coverage

You can be sure the code adheres to (at least some) good practices by running the linter (alternatively, using -v): 

    golangci-lint run
    
You can also run the built-in tests with 

    go test ./...

And review the test coverage using the nice Go builtin tool with:

    go test -coverprofile=cover.out ./... && go tool cover -html=cover.out
