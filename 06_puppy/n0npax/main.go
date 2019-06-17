package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	store := MemStore{}
	puppy := Puppy{ID: 1, Breed: "Mix", Colour: "White", Value: "Free"}

	id := store.CreatePuppy(&puppy)
	fmt.Fprint(out, "Puppy id: ", id)
}
