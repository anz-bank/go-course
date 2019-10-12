package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	var puppyStore Storer = NewMapStore()
	pup := Puppy{ID: 1, Breed: "kelpie", Colour: "brown", Value: "indispensable"}
	id, _ := puppyStore.CreatePuppy(&pup)
	if pup, err := puppyStore.ReadPuppy(id); err == nil {
		fmt.Fprintf(out, "retrieved: %v %v %v\n", pup.Breed, pup.Colour, pup.Value)
	}
}
