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
