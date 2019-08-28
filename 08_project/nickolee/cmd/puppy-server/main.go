package main

import (
	"fmt"
	"io"
	"os"

	"github.com/anz-bank/go-course/08_project/nickolee/puppystorer"
	// store "github.com/anz-bank/go-course/08_project/nickolee/pkg/puppy/store"
)

var out io.Writer = os.Stdout

func main() {
	// instantiate new puppy store
	store := store.NewMapStore()
	// create a puppy and store in store (no pun intended lol)
	createdPuppy, _ := store.CreatePuppy(&{Breed: "The Hound", Colour: "Of Baskerville", Value: 12300.90})
	fmt.Fprintln(out, "Puppy with ID %d has been created", createdID)

	retrievedPuppy, _ := store.ReadPuppy(createdPuppy)
	fmt.Fprintln(out, "Retrieved puppy ID: ", retrievedPuppy)

	updateResult := store.UpdatePuppy(createdID, &{{Breed: "Arcanine", Colour: "Level 100", Value: 9300.90}})
	fmt.Fprintln(out, "Update puppy operation result:", updateResult)

	deleteResult, _ := store.DeletePuppy(createdPuppy)
	fmt.Fprintln(out, "Delete puppy operation result:", deleteResult)
}
