# Lab 1 - Fibonacci

- Create an executable go program in directory `01_fib/USERNAME`
- Write a function that prints the first n Fibonacci numbers

```
func fib(n int)
```

Call `fib(7)` in `main` to print

```
1
1
2
3
5
8
13
```

Bonus points: For negative n print Negafibonacci numbers.

# Lab 2 - Bubble sort

- Create an executable go program in directory `02_bubble/USERNAME`
- Write a function that returns a sorted copy of `int` slice `s` using [bubble sort](https://en.wikipedia.org/wiki/Bubble_sort):

```
func bubble(s []int) []int
```

Call `fmt.Println(bubble([]int{3, 2, 1, 5}))` in `main` to print:

```
[1 2 3 5]
```

- Bonus points: implement [Insertion sort](https://en.wikipedia.org/wiki/Insertion_sort)
- Extra bonus points: implement an _O(n_ _log(n))_ sorting algorithm

# Lab 3 - Letter frequency

- Create an executable go program in directory `03_letters/USERNAME`
- Write a function that returns a mapping of each letter to its frequency:

```
func letters(s string) map[rune]int
```

Write a function that returns a sorted slice of strings with elements `"{key}:{val}"`. Use package [sort](https://golang.org/pkg/sort/):

```
func sortLetters(m map[rune]int) []string
```

Call `fmt.Println(strings.Join(sortLetters(letters("aba")), "\n"))` in `main` to print:

```
a:2
b:1
```

Bonus points: comprehensive tests

# Lab 4 - Numeronym

- Create an executable go program in directory `04_numeronym/USERNAME`
- Write a function that returns a slice of numeronyms for its input strings:

```
func numeronyms(vals ...string) []string
```

Call `fmt.Println(numeronyms("accessibility", "Kubernetes", "abc"))` in `main` to print:

```
[a11y K8s abc]
```

# Lab 5 - Stringer

- Create an executable go program in directory `05_stringer/USERNAME`
- Make the `IPAddr` type implement `fmt.Stringer` to print the address as a dotted quad
- Find hints at [tour of go exercise: stringers](https://tour.golang.org/methods/18)
- Call `fmt.Println(IPAddr{127, 0, 0, 1})` in `main` to print:

```
127.0.0.1
```

# Lab 6 - CRUD puppy with interface

- Create an executable go program in directory `06_puppy/USERNAME` (see [hints](https://github.com/anz-bank/go-samplerest/blob/master/pkg/pet/types.go))
- Implement a `Puppy` struct containing `ID`, `Breed`, `Colour`, `Value`
- Create `Storer` interface with [crud](https://en.wikipedia.org/wiki/Create,\_read,\_update_and_delete) methods for `Puppy`
- Write a `MapStore` implementation of `Storer` backed by a `map`
- Write a `SyncStore` implementation of `Storer` backed by a [sync.Map](https://golang.org/pkg/sync/#Map)
- Keep all implementation files in the same folder and in package `main`
- Test against the `Storer` interface and run in [suite](https://godoc.org/github.com/stretchr/testify/suite) with both implementations

# Lab 7 - Errors

- Create an executable go program in directory `07_errors/USERNAME`
- Copy the CRUD puppy from upstream master `06_puppy/USERNAME`
- Add a custom error type `Error` with fields `Message` and `Code`
- Extend the `Storer` interface for all methods to also return `error`
- Create errors for:

        - Value < 0
        - ID not found in Read, Update and Delete

- Add locking for proper use of sync.Map
- Bonus points: Add a third `Storer` implementation using [leveldb](https://github.com/syndtr/goleveldb)

# Lab 8 - Project Layout

- Copy the CRUD puppy from upstream master `07_errors/USERNAME`
- Create directory `08_project/USERNAME` containing

```
├── README.md
├── pkg
│   └── puppy
│       ├── types.go
│       ├── types_test.go
│       ├── errors.go
│       ├── errors_test.go
│       └── store
│           ├── storer.go
│           └── .... store files and tests, e.g. mapstore.go
└── cmd
    └── puppy-server
        └── main.go
```

Add project introduction and how to build, run & test it to `README.md`

# Lab 9 - JSON puppy

- Create directory `09_json/USERNAME` containing a copy of upstream master `08_project/USERNAME`
- Add JSON tags to puppy data type
- Test marshalling and unmarshalling using [require.JSONEq](https://godoc.org/github.com/stretchr/testify/require#JSONEq)
- Add command line flag `-d FILE` with long form `--data FILE` using [kingpin.v2](https://godoc.org/gopkg.in/alecthomas/kingpin.v2)
- FILE should contain an array of puppies in JSON format. Parse this file and store its contents.

# Lab 10 - Puppy REST

- Create directory `10_rest/USERNAME` containing a copy of upstream master `09_json/USERNAME`
- Add file `pkg/puppy/rest.go` implementing:

```
GET    /api/puppy/{id}
POST   /api/puppy/          Payload: Puppy JSON without ID
PUT    /api/puppy/{id}      Payload: Puppy JSON without ID
DELETE /api/puppy/{id}
```

- Use [net/http/httptest](https://golang.org/pkg/net/http/httptest/) for testing
- Add flag `-p PORT` with long flag `--port PORT` to command line flags
- Add flag `-s STORE` with long flag `--store STORE` with accepted values:

```
map, sync, db
```
Document the API in README.md

# Lab 11 - Puppy Notifications

- Create directory `11_notify/USERNAME` containing a copy of upstream master `10_rest/USERNAME`
- Create `cmd/lostpuppy-service/main.go` running single endpoint:
```
POST   /api/lostpuppy/          Payload: { id: PUPPY_ID }
This stubbed endpoint returns with 2 second delay:
HTTP status 201 for even IDs
HTTP status 500 for odd IDs
```
Update Puppy Delete method to notify lostpuppy-service in a goroutine and log response code asynchronously.

