# PuppyStore

`PuppyStore` provides CRUD operation for PuppyStore management.

# Installation

- Installations of Go v1.12+ according to [official instructions](https://golang.org/doc/install), including [GOPATH env variables](https://golang.org/doc/code.html#GOPATH).

# Build, install, test the PuppyStore with:

```bash
go build ./...
go install ./...
go test ./...
```

# To run Puppy Store:

```bash
go run ./cmd/puppy-server/main.go --data "pkg/puppy/store/testdata/many-puppies.json"

or the short version is :
go run ./cmd/puppy-server/main.go -d "pkg/puppy/store/testdata/many-puppies.json"
```

For linting:

```bash
golangci-lint run
```

For Test Coverage:

```bash
go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
```
