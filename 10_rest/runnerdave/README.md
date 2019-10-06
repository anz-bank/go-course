# Puppy server

Storage of Puppies with two different storage methods.

## Pre-requisites

You need to have [Go](https://golang.org/doc/install) installed

## Build and install

    go install ./...

### Run installed version

    puppy_server

(assuming **\$GOPATH/bin/** is in your \$PATH)

## Build executable

    go build -o puppy-server cmd/puppy_server/main.go

### Run executable

There are three flags with for the executable:

- -d,--data: if a json file wants to be given as input to prefill the database, pass it as a value to this flag, if not set then a default file is used.

  ./puppy-server -d puppydata/data.json

- -s, --store: either _map_ or _sync_ if not set then defaults to _sync_, can be used with the -d, --data flag to choose the storage method.

  ./puppy-server -d puppydata/data.json -s map

- -p, --port: if set then an **API** server is started in the given port

  ./puppy-server -p 3000 -s map

## Test

    go test ./...

## API endpoints

### Create puppies

Endpoint: {host:port}/api/puppy/
Method: POST
Example:

    curl -X POST http://localhost:3000/api/puppy/ \
      -d '{
        "breed": "CUCHICUCHI",
        "color": "blue",
        "value": 39170.65
      }'

### Retrieve puppies

Endpoint: {host:port}/api/puppy/{id}
Method: GET
Example:

    curl -X GET http://localhost:3000/api/puppy/1

### Modify existing puppies

Endpoint: {host:port}/api/puppy/{id}
Method: PUT
Example:

    curl -X PUT http://localhost:3000/api/puppy/1 \
      -d '{
        "breed": "CUCHICUCHI",
        "color": "blue",
        "value": 173.98
      }'

### Delete puppies

Endpoint: {host:port}/api/puppy/{id}
Method: DELETE
Example:

    curl -X DELETE http://localhost:3000/api/puppy/2
