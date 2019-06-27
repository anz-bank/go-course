package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	store := NewSyncStore()
	_ = store.CreatePuppy(Puppy{
		ID:     "1",
		Colour: "Red",
	})

	puppy, _ := store.ReadPuppy("1")
	fmt.Fprint(out, puppy.Colour)
}
