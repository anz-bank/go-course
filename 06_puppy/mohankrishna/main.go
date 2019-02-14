package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	mapStorage := CreateStore()
	syncStorage := CreateSyncStore()

	puppy := Puppy{
		ID:     11,
		Breed:  "Sheep herder",
		Colour: "Brown",
		Value:  1000,
	}

	//Map store
	mapStorage.CreatePuppy(&puppy)
	fmt.Fprintln(out, *mapStorage.ReadPuppy(11))

	puppy.Value = 10000

	mapStorage.UpdatePuppy(puppy.ID, &puppy)
	fmt.Fprintln(out, *mapStorage.ReadPuppy(11))

	mapStorage.DeletePuppy(puppy.ID)
	fmt.Fprintln(out, *mapStorage.ReadPuppy(11))

	//Sync.map store
	puppy.Value = 1000
	syncStorage.CreatePuppy(&puppy)
	fmt.Fprintln(out, *syncStorage.ReadPuppy(11))

	puppy.Value = 10000

	syncStorage.UpdatePuppy(puppy.ID, &puppy)
	fmt.Fprintln(out, *syncStorage.ReadPuppy(11))

	syncStorage.DeletePuppy(puppy.ID)
	fmt.Fprintln(out, *syncStorage.ReadPuppy(11))
}
