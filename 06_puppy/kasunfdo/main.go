package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	store := NewMapStore()
	puppy := Puppy{ID: 1, Breed: "Husky", Colour: "White", Value: 4999.98}

	id := store.CreatePuppy(puppy)
	fmt.Fprintf(out, "Puppy(%d) added to store\n", id)
}
