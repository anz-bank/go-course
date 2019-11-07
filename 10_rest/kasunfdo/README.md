# Puppy Store - An in-memory store for puppies

<!-- MarkdownTOC -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Build, execute, test, lint](#betl)
    - [Build](#build)
    - [Execute](#execute)
        - [Flags](#flags)
    - [Test](#test)
    - [Format and Lint](#lint)
    - [Review unit test coverage](#coverage)
- [API Documentation](#api)
    - [Create Puppy](#create)
    - [Read Puppy](#read)
    - [Update Puppy](#update)
    - [Delete Puppy](#delete)

<!-- /MarkdownTOC -->

## Overview

Puppy REST is a simple REST web application with [CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) 
endpoints for [Puppies](pkg/puppy/types.go). Puppy server uses an in-memory store which is designed to store 
[Puppy](pkg/puppy/types.go) objects and store is capable of creating, reading, updating, and deleting puppies in puppy store.

## Prerequisites

-   Install `go 1.12` according to [official installation instruction](https://golang.org/doc/install)
-   Clone this project outside your `$GOPATH` to enable [Go Modules](https://github.com/golang/go/wiki/Modules)
-   Install `goimports` according to [instructions](https://godoc.org/golang.org/x/tools/cmd/goimports)
-   Install `golangci-lint` according to [instructions](https://github.com/golangci/golangci-lint#local-installation)

## Build, execute, test, lint <a name="betl"></a>

#### Build

    go build -o puppystore cmd/puppy-server/main.go
    
#### Execute

    ./puppystore -d <path to puppy data file in json format>
    
##### Flags (Optional) <a name="flags"></a>

- `--port`, `-p`: Puppy server port
- `--store`, `-s`: Puppy store type (`map`/`sync`)

#### Test

    go test ./...

#### Format and lint <a name="lint"></a>

    goimports -w .
    golangci-lint run

#### Review unit test coverage <a name="coverage"></a>

    go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## API

#### Create puppy <a name="create"></a>

Add a puppy to the store.

URL : `/api/puppy/`

Method : `POST`

Data example :

```json
{
    "breed"  : "Labrador",
    "colour" : "White",
    "value"  : 1200
}
```

##### Response (Success): 

Code : `201 Created`

Content example :

```json
{
    "id"     : 1,
    "breed"  : "Labrador",
    "colour" : "White",
    "value"  : 1200
}
```

#### Read puppy <a name="read"></a>

Retrieve a puppy from the store.

URL : `/api/puppy/{id}`

Method : `GET`

##### Response (Success): 

Code : `200 OK`

Content example :

```json
{
    "id"     : 1,
    "breed"  : "Labrador",
    "colour" : "White",
    "value"  : 1200
}
```

#### Update puppy <a name="update"></a>

Update an existing puppy in the store.

URL : `/api/puppy/`

Method : `PUT`

Data example :

```json
{
    "id"     : 1,
    "breed"  : "Labrador",
    "colour" : "Brown",
    "value"  : 1200
}
```

##### Response (Success): 

Code : `200 OK`

Content example :

```json
    "puppy updated"
```

#### Delete puppy <a name="delete"></a>

Remove a puppy from the store.

URL : `/api/puppy/{id}`

Method : `DELETE`

##### Response (Success): 

Code : `200 OK`

Content example :

```json
    "puppy deleted"
```
