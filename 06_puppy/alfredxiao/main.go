package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	store := NewMapStore()
	id := store.CreatePuppy(Puppy{
		Colour: "Red",
	})

	puppy, _ := store.ReadPuppy(id)
	fmt.Fprint(out, puppy.Colour)
}
