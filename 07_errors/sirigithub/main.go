package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	mapStore := NewMapStore()
	puppy1 := Puppy{Breed: "Mix", Color: "White", Value: 300}
	id, _ := mapStore.CreatePuppy(&puppy1)
	println(id)
	syncStore := NewSyncStore()
	id2, _ := syncStore.CreatePuppy(&puppy1)
	println(id2)

	p1, _ := mapStore.ReadPuppy(id)
	fmt.Fprintln(out, "Read puppy from Mapstore with ID:", p1.ID)
	p2, _ := syncStore.ReadPuppy(id2)
	fmt.Fprintln(out, "Read puppy from SyncStore with ID:", p2.ID)
}
