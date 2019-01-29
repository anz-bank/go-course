# Go Training

[![Build Status](https://travis-ci.com/anz-bank/go-training.svg?branch=master)](https://travis-ci.com/anz-bank/go-training)
[![Coverage](https://codecov.io/gh/anz-bank/go-training/branch/master/graph/badge.svg)](https://codecov.io/gh/anz-bank/go-training)
[![GolangCI](https://golangci.com/badges/github.com/anz-bank/go-training.svg)](https://golangci.com/r/github.com/anz-bank/go-training)

This project is a playground for hands-on golang training.

## Prerequisites

-   Install `go 1.11` according to [official installation instruction](https://golang.org/doc/install)
-   Clone this project outside your `$GOPATH` to enable [Go Modules](https://github.com/golang/go/wiki/Modules)
-   Install `golangci-lint` according to [instructions](https://github.com/golangci/golangci-lint#local-installation)

## Build, execute, test, lint

Build and install this project with

    go install ./...

Execute its binary with

    0_hello_world

Test it with

    go test ./...

Lint it with

    golangci-lint run
