package main

import (
	"fmt"
	"io"
	"os"

	puppy "github.com/anz-bank/go-course/08_project/mohitnag/pkg/puppy"
	store "github.com/anz-bank/go-course/08_project/mohitnag/pkg/puppy/store"
)

var out io.Writer = os.Stdout

func main() {
	mapStore := store.MapStore{}
	syncStore := store.SyncStore{}
	pup := puppy.Puppy{ID: 1, Breed: "dog", Colour: "white", Value: "$2"}

	//CreatePuppy
	_ = mapStore.CreatePuppy(pup)
	_ = syncStore.CreatePuppy(pup)

	// ReadPuppy
	puppyMemstore, _ := mapStore.ReadPuppy(1)
	puppySyncstore, _ := syncStore.ReadPuppy(1)
	fmt.Fprintln(out, "Mapstore: "+puppyMemstore.Colour)
	fmt.Fprintln(out, "Syncstore: "+puppySyncstore.Colour)

	// // Update Puppy
	updatePuppy := puppy.Puppy{ID: 1, Breed: "dog", Colour: "black", Value: "$2"}
	_ = mapStore.UpdatePuppy(updatePuppy)
	_ = syncStore.UpdatePuppy(updatePuppy)

	// DeletePuppy
	_ = mapStore.DeletePuppy(1)
	_ = syncStore.DeletePuppy(1)
}
