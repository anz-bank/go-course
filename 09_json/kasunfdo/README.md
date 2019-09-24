# Puppy Store - An in-memory store for puppies

<!-- MarkdownTOC -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Build, execute, test, lint](#betl)
    - [Build](#build)
    - [Execute](#execute)
    - [Test](#test)
    - [Format and Lint](#lint)
    - [Review unit test coverage](#coverage)

<!-- /MarkdownTOC -->

## Overview

Puppy Store is a simple in-memory store for [Puppy](pkg/puppy/types.go) objects. Puppy Store 
is implemented with CRUD methods for creating, reading, updating, and deleting puppies in puppy store.

## Prerequisites

-   Install `go 1.12` according to [official installation instruction](https://golang.org/doc/install)
-   Clone this project outside your `$GOPATH` to enable [Go Modules](https://github.com/golang/go/wiki/Modules)
-   Install `goimports` according to [instructions](https://godoc.org/golang.org/x/tools/cmd/goimports)
-   Install `golangci-lint` according to [instructions](https://github.com/golangci/golangci-lint#local-installation)

## Build, execute, test, lint <a name="betl"></a>

#### Build

    go build -o puppystore cmd/puppy-server/main.go
    
#### Execute

    ./puppystore

#### Test

    go test ./...

#### Format and lint <a name="lint"></a>

    gofmt -w .
    goimports -w .
    golangci-lint run

#### Review unit test coverage <a name="coverage"></a>

    go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
