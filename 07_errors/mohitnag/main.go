package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	mapStore := MapStore{}
	syncStore := SyncStore{}
	puppy := Puppy{ID: 1, Breed: "dog", Colour: "white", Value: "$2"}

	//CreatePuppy
	_ = mapStore.CreatePuppy(puppy)
	_ = syncStore.CreatePuppy(puppy)

	// ReadPuppy
	puppyMemstore, _ := mapStore.ReadPuppy(1)
	puppySyncstore, _ := syncStore.ReadPuppy(1)
	fmt.Fprintln(out, "Memstore: "+puppyMemstore.Colour)
	fmt.Fprintln(out, "Syncstore: "+puppySyncstore.Colour)

	// Update Puppy
	updatePuppy := Puppy{ID: 1, Breed: "dog", Colour: "black", Value: "$2"}
	_ = mapStore.UpdatePuppy(updatePuppy)
	_ = syncStore.UpdatePuppy(updatePuppy)

	// DeletePuppy
	_ = mapStore.DeletePuppy(1)
	_ = syncStore.DeletePuppy(1)
}
