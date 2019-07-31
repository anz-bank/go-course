package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

type Puppy struct {
	ID    uint32
	Breed string
	Color string
	Value string
}

func main() {
	cutePuppy := Puppy{1, "Jack Russell", "White and Brown", "500"}
	naughtyPuppy := Puppy{1234, "Fox Terrier", "Black", "1300"}

	var store = MapStore{}
	store.ms = map[uint32]Puppy{}
	createdID, _ := store.CreatePuppy(&cutePuppy)
	fmt.Fprintln(out, "First Created Puppy ID:", createdID)

	readPuppy, _ := store.ReadPuppy(createdID)
	fmt.Fprintln(out, "ReadPuppy:", readPuppy)

	updateResult, _ := store.UpdatePuppy(createdID, &naughtyPuppy)
	fmt.Fprintln(out, "Update puppy result:", updateResult)

	deleteSuccess, _ := store.DeletePuppy(createdID)
	fmt.Fprintln(out, "Result of deleting puppy:", deleteSuccess)
}
