# Table of contents

- [Introduction](#introduction)
- [API documentation](#api-documentation)
  * [Object specification](#object-specification)
  * [API requests](#api-requests)
    + [```POST /api/puppy/```](#---post--api-puppy----)
    + [```GET /api/puppy/{id}```](#---get--api-puppy--id----)
    + [```PUT /api/puppy/{id}```](#---put--api-puppy--id----)
    + [```DELETE /api/puppy/{id}```](#---delete--api-puppy--id----)
- [Running it](#running-it)
  * [Prerequisites](#prerequisites)
  * [Build, install, execute](#build--install--execute)
    + [Short version](#short-version)
    + [Long version](#long-version)
  * [Lint, test, coverage](#lint--test--coverage)

# Introduction

This is part of the [Go](https://golang.org/) [Course](https://github.com/anz-bank/go-course) done by [ANZ Bank](https://www.anz.com.au) during 2019.

It contains code to create a basic webserver that serves an REST API over HTTP allowing POST/PUT/GET/DELETE operations on simple objects that can be backed by different storage methods.

The project is organized as follows:
- `cmd/puppy-server` contains the code for the server executable itself.
- `pkg/puppy` contains the type definitions, interfaces and error values for the package
- `pkg/puppy/store` contains the bulk of the code, 2 separate store backends based on the native Golang map and on sync.map.

# API documentation
## Object specification
The object Puppy is represented on the API requests by a JSON containing pre-defined fields, of which all are optional when sent by the client, and aways contain at least a valid `id` when sent by the server. Any alien field is ignored.

In case a pre-defined field is ommitted, it defaults to "" (empty string) for strings, and 0 (zero) for numbers.

Below, `{id}` means the object identifier on the URL called, and `id` means the object identifier on the JSON. Both values are, or should be, the same.

The JSON field `id` is special: its value is always supplied by the server. Any value passed on it by the client on any request is either ignored or causes an error, depending on the request type.

The URL field `{id}` should be a non-zero, positive integer. 

Valid JSON fields are:
> - **id**: Numeric positive integer.
> - **breed**: String
> - **colour**: String
> - **value**: Numeric positive integer

Example valid JSON:
 
```json
{
    "id": 290,
    "breed": "Chihuahua",
    "colour": "Cream",
    "value": 300
}
```
## API requests
### ```POST /api/puppy/```
Creates an object. 

**Input**: Valid JSON on body. If `id` is supplied on the JSON, it is an error for it to be different of 0 (zero).

**Output** is one of:

  Type | Header   | Body   | Meaning
 ------| -------- | ------ | -------
 Valid | `201 Created` | Puppy JSON | Object created successfully. Returned JSON contains the full object, including the `id` value assigned by the API that can be used to `GET` it. Currently `id` values start at 1 and increment by 1 for each new object created; however, do not rely on this as it may change without warning.
 Error | `400 Bad Request` | `400 Bad Request` | `id` value is invalid and/or invalid JSON. No object was created.
 
### ```GET /api/puppy/{id}```
Returns a JSON object specified by `{id}`, representing a valid object existing in storage.

**Input**: `{id}` on URL only. Request body is ignored.

**Output** is one of:

  Type | Header   | Body   | Meaning
 ------| -------- | ------ | -------
 Valid | `200 OK` | Valid JSON  | Object read successfully. Returned JSON contains the full object.
 Error | `404 Not Found` | `404 Not Found` | Object `{id}` not found.
 Error | `400 Bad Request` | `400 Bad Request` | `{id}` value invalid and/or invalid JSON.

### ```PUT /api/puppy/{id}```
Updates an existing object identified by `{id}`.

**Input**: `{id}` on URL, valid JSON on body. JSON field `id` is ignored if supplied.

**Output** is one of:

  Type | Header   | Body   | Meaning
 ------| -------- | ------ | -------
 Valid | `200 OK` | `200 OK` | Object updated successfully.
 Error | `404 Not Found` | `404 Not Found` | Object `{id}` not found. No object was updated.
 Error | `400 Bad Request`| `400 Bad Request` | `{id}` value invalid and/or invalid JSON. No object was updated.
 
### ```DELETE /api/puppy/{id}```
Deletes an existing object identified by `{id}`. 

**Input**: `{id}` on URL. Request body is ignored.

**Output** is one of:

  Type | Header   | Body   | Meaning
 ------| -------- | ------ | -------
 Valid | `200 OK`| `200 OK` | Object deleted successfully.
 Error | `404 Not Found` | `404 Not Found` | Object `{id}` not found. No object was deleted.
 Error | `400 Bad Request` | `400 Bad Request` | `{id}` value invalid and/or invalid JSON. No object was deleted.


# Running it

## Prerequisites
-   Install [`go`](https://golang.org/doc/install) and alternatively [`golangci-lint`](https://github.com/golangci/golangci-lint#local-installation) if you want to run tests or lint
-   Clone this project outside your `$GOPATH` to enable [Go Modules](https://github.com/golang/go/wiki/Modules)

All directory paths mentioned are relative to the root of the project.

## Build, install, execute

### Short version
For the anxious, you can just run the main executable quickly doing

    go run ./cmd/puppy-server/main.go

### Long version
Alternatively, you can build, install and run from your `$GOPATH` with

    go install ./...
    puppy-server

Or yet build and run from the project's root directory with 

    go build -o puppy-server cmd/puppy-server/main.go
    ./puppy-server

## Lint, test, coverage

You can be sure the code adheres to (at least some) good practices by running the linter (alternatively, using -v): 

    golangci-lint run
    
You can also run the built-in tests with 

    go test ./...

And review the test coverage using the nice Go builtin tool with:

    go test -coverprofile=cover.out ./... && go tool cover -html=cover.out
