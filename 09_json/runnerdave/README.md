# Puppy server

Storage of Puppies with two different storage methods.

## Pre-requisites

You need to have [Go](https://golang.org/doc/install) installed

## Build and install

```
go install ./...
```

### Run installed version

```
puppy_server
```
(assuming __$GOPATH/bin/__ is in your $PATH)

## Build executable

```
go build -o puppy-server cmd/puppy_server/main.go
```

### Run executable
```
./puppy-server
```

## Test

```
go test ./...
```
