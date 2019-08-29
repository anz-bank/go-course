package main

import (
	"fmt"
	"io"
	"os"

	puppy "github.com/anz-bank/go-course/08_project/nickolee/pkg/puppy"
	store "github.com/anz-bank/go-course/08_project/nickolee/pkg/puppy/store"
)

var out io.Writer = os.Stdout

func main() {
	// instantiate new puppy store
	store := store.NewMapStore()
	// create a puppy and store in store (no pun intended lol)
	createdPuppy, _ := store.CreatePuppy(&puppy.Puppy{Breed: "The Hound", Colour: "Of Baskerville", Value: 12300.90})
	fmt.Fprintf(out, "Puppy with ID %d has been created\n", createdPuppy)

	retrievedPuppy, _ := store.ReadPuppy(createdPuppy)
	fmt.Fprintln(out, "Retrieved puppy:", retrievedPuppy)

	updateResult := store.UpdatePuppy(createdPuppy, &puppy.Puppy{Breed: "Arcanine", Colour: "Level 100", Value: 9300.90})
	fmt.Fprintln(out, "Update puppy operation result:", updateResult)

	deleteResult := store.DeletePuppy(createdPuppy)
	fmt.Fprintln(out, "Delete puppy operation result:", deleteResult)
}
