package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	store := NewMemStore()
	puppy := Puppy{ID: 1, Breed: "Mix", Colour: "White", Value: 100}

	id, _ := store.CreatePuppy(&puppy)
	fmt.Fprint(out, "Puppy id: ", id)
}
