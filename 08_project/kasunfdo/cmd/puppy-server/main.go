package main

import (
	"fmt"
	"io"
	"os"

	"github.com/anz-bank/go-course/08_project/kasunfdo/pkg/puppy"
	"github.com/anz-bank/go-course/08_project/kasunfdo/pkg/puppy/store"
)

var out io.Writer = os.Stdout

func main() {
	store := store.NewMapStore()
	p := puppy.Puppy{Breed: "Husky", Colour: "White", Value: 4999.98}
	id, _ := store.CreatePuppy(p)
	fmt.Fprintf(out, "Puppy(%d) added to store\n", id)
}
