package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	var store Storer
	store = NewSyncStore()
	ID := store.CreatePuppy(Puppy{2, "German Shepherd", "white", 200})
	v := store.ReadPuppy(ID)
	fmt.Fprintf(out, "Create Puppy breed = %s, Colour = %s\n", v.Breed, v.Colour)
	store.UpdatePuppy(ID, Puppy{2, "Pug", "Black", 200})
	v = store.ReadPuppy(ID)
	fmt.Fprintf(out, "Update Puppy breed = %s, Colour = %s\n", v.Breed, v.Colour)
	store.DeletePuppy(ID)
	v = store.ReadPuppy(ID)
	fmt.Fprintf(out, "Delete Puppy breed = %s, Colour = %s\n", v.Breed, v.Colour)
	fmt.Fprintf(out, "============================================\n")
	store = NewMapStore()
	ID = store.CreatePuppy(Puppy{3, "Bulldog", "white", 200})
	v = store.ReadPuppy(ID)
	fmt.Fprintf(out, "Puppy breed = %s, Colour = %s\n", v.Breed, v.Colour)
	store.UpdatePuppy(ID, Puppy{4, "Poodle", "Brown", 200})
	v = store.ReadPuppy(ID)
	fmt.Fprintf(out, "Update Puppy breed = %s, Colour = %s\n", v.Breed, v.Colour)
	store.DeletePuppy(ID)
	v = store.ReadPuppy(ID)
	fmt.Fprintf(out, "Delete Puppy breed = %s, Colour = %s\n", v.Breed, v.Colour)
}
