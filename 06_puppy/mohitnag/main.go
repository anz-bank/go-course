package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	m := MemStore{}
	puppy := Puppy{ID: 1, Breed: "dog", Colour: "white", Value: "$2"}
	_ = m.CreatePuppy(puppy)
	puppy, _ = m.ReadPuppy(1)
	fmt.Fprint(out, puppy.Colour)
}
