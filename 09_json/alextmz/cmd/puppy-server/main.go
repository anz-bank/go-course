package main

import (
	"fmt"
	"io"
	"os"

	"github.com/anz-bank/go-course/08_project/alextmz/pkg/puppy"
	"github.com/anz-bank/go-course/08_project/alextmz/pkg/puppy/store"
)

var out io.Writer = os.Stdout

func main() {
	ps := store.NewmapStore()
	p1 := &puppy.Puppy{ID: 1, Breed: "Dogo", Colour: "white", Value: 50.0}

	fmt.Printf("Trying to create first puppy on store - ID %d {Breed: %s, Colour: %s, Value %d}.\n",
		p1.ID, p1.Breed, p1.Colour, p1.Value)

	_ = ps.CreatePuppy(p1)
	fmt.Printf("Puppy ID %d created on store.\n", p1.ID)

	fmt.Printf("Trying to read puppy values back from the store...\n")
	p2, _ := ps.ReadPuppy(p1.ID)

	fmt.Fprintf(out, "Read Puppy by ID: %d {Breed: %s, Colour: %s, Value %d}.\n",
		p2.ID, p2.Breed, p2.Colour, p2.Value)
}
