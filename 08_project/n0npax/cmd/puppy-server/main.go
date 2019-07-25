package main

import (
	"fmt"
	"io"
	"os"

	puppy "github.com/anz-bank/go-course/08_project/n0npax/pkg/puppy"
	store "github.com/anz-bank/go-course/08_project/n0npax/pkg/puppy/store"
)

var out io.Writer = os.Stdout

func main() {
	store := store.MemStore{}
	p := puppy.Puppy{Breed: "Mix", Colour: "White", Value: 100}

	id, _ := store.CreatePuppy(&p)
	fmt.Fprint(out, "Puppy id: ", id)
}
