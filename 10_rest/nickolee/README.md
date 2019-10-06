# Project: Final Step in Incrementally Building Towards a Puppy REST API

This project is part of the [go course](https://github.com/anz-bank/go-course/) run by ANZ Bank.

It builds you up incrementally with the ultimate end goal of having a fully functioning Rest API.

And finally after a deep and delightful learning journey we have finally arrived at Lab 10!

It aims to impart several key learnings such as but not limited to:

- Working with custom types in go using structs
- Understanding and working with interfaces in go
- Implementing multiple implementations of an interface
- Working with various tools in the Go toolchain such as `golangci-lint` and `go test -coverprofile=coverage.out` to check for 100% test coverage or fail CI
- Go standard project layout
- Best practices with Git. Proper commit messages and PRs. Git rebasing when necessary to keep commit history clean etc.
- Creating custom error types to suit your project needs
- Thread safety using sync.Mutex to prevent errors associated with data race conditions
- Test coverage using [test suites] (https://godoc.org/github.com/stretchr/testify/suite) as well as [table driven testing](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests) 
- Implement ability to convert between Go in memory objects to JSON representation and vice versa using `json.Marshal() and json.Unmarshal()`
- Implement ability to accept user input from CLI and read data from a file
- Load data from json text file and store in an instance of storer
- Learn how to create REST APIs to expose functionality over the network
- Learn how to test REST APIs with the net/http/httptest package
- Learning how to use the Go standard library as well as external third party libraries such as [chi](https://godoc.org/github.com/go-chi/chi), [kingpin](https://godoc.org/gopkg.in/alecthomas/kingpin.v2) and [testify](https://godoc.org/github.com/stretchr/testify)

## Prerequisites

-   Install `go 1.12` according to [official installation instruction](https://golang.org/doc/install)
-   Clone this project outside your `$GOPATH` to enable [Go Modules](https://github.com/golang/go/wiki/Modules)
-   Install `golangci-lint` according to [instructions](https://github.com/golangci/golangci-lint#local-installation)

## Build

To build the project, run the following command while you are in the root of the project folder:

`go build ./cmd/puppy-server`

## Run
`go run ./cmd/puppy-server/main.go --data puppy-data/puppies.json --port 7777 --store map`

Alternatively, simply run:
`go run ./cmd/puppy-server/main.go`

Which will simply revert to pre-specified default values.

Note that there are three flags which work with main.go:

1. **-d,--data:** if you wish to specify a .json file to load in seed data for the database, pass it as a value to this flag. Note that if no file is specified then you will get an empty Storer with no puppies inside.
2. **-s, --store:** either _map_ or _sync_. Note that if not set then the value defaults to _map_
3. **-p, --port:** if set then an **API** server is started on the user specified port. Must be a valid port number between 0 and 65535. If port is not set it will default to port 7777.

## Test

To test this project, follow the steps below:

1. Run test cases:`go test ./...`
2. Lint it: `golangci-lint run`
3. Ensure 100% test coverage `go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out`

## API Spec

### POST /api/puppy

Adds a new puppy to the store, returns the stored puppy with a generated id.

 Example:

    curl -X POST http://localhost:7777/api/puppy/ \
      -d '{
        "breed": "Iron Dog",
        "colour": "Red",
        "value": 9500
      }'

Note that a negative puppy value will not be accepted.

### GET /api/puppy/{id}

Retrieves a puppy from the store. Returns a stored puppy.

 Example:

    curl -X GET http://localhost:7777/api/puppy/1

### PUT /api/puppy/{id}

Updates an existing puppy in the store, returns the newly updated puppy.

 Example:

    curl -X PUT http://localhost:7000/api/puppy/1 \
      -d '{
        "breed": "Spider Dog not Snoop Dogg",
        "colour": "Friendly Neighbourhood Colors",
        "value": 5555
      }'

As with POST this will not accept a negative puppy value since we believe every dog has intrinsic value and is worth something!

### DELETE /api/puppy/{id}

Deletes a puppy of a given id from the store, returns a sucess message.

 Example:

    curl -X DELETE http://localhost:7777/api/puppy/1


  