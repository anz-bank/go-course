package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	p1 := Puppy{Breed: "Poodle", Colour: "White", Value: 1280.5}
	p2 := Puppy{Breed: "Poodle", Colour: "Grey", Value: 1340.5}
	mapStore := NewMapStore()
	syncStore := NewSyncStore()
	_ = mapStore.CreatePuppy(p1)
	_ = syncStore.CreatePuppy(p2)
	puppy, _ := mapStore.ReadPuppy(0)
	puppy2, _ := syncStore.ReadPuppy(0)
	fmt.Fprintf(out, "Puppy ID %d is %v", puppy.ID, puppy.Value)
	fmt.Fprintf(out, "Puppy ID %d is %v", puppy2.ID, puppy2.Value)
}
