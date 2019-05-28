# Lab 8 - Project Layout

- Copy the CRUD puppy from upstream master `07_errors/USERNAME`
- Create directory `08_project/USERNAME` containing

```
├── README.md
├── go.mod
├── go.sum
├── pkg
│   └── USERNAME-puppy
│       ├── types.go
│       ├── types_test.go
│       ├── errors.go
│       ├── errors_test.go
│       └── store
│           ├── storer.go
│           └── .... store files and tests, e.g. mapstore.go
├── cmd
│   └── USERNAME-puppy-server
│       └── main.go
└── vendor
```

Add project introduction and how to build, run & test it to `README.md`
