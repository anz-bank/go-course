package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	store := NewMapStore()
	puppy := Puppy{Breed: "Husky", Colour: "White", Value: 4999.98}
	id, _ := store.CreatePuppy(puppy)
	fmt.Fprintf(out, "Puppy(%d) added to store\n", id)
}
