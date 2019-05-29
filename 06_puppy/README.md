# Lab 6 - CRUD puppy with interface

- Create an executable go program in directory `06_puppy/USERNAME` (see [hints](https://github.com/anz-bank/go-samplerest/blob/master/pkg/pet/types.go))
- Implement a `Puppy` struct containing `ID`, `Breed`, `Colour`, `Value`
- Create `Storer` interface with [crud](https://en.wikipedia.org/wiki/Create,\_read,\_update_and_delete) methods for `Puppy`
- Write a `MapStore` implementation of `Storer` backed by a `map`
- Write a `SyncStore` implementation of `Storer` backed by a [sync.Map](https://golang.org/pkg/sync/#Map)
- Keep all implementation files in the same folder and in package `main`
- Test against the `Storer` interface and run in [suite](https://godoc.org/github.com/stretchr/testify/suite) with both implementations
