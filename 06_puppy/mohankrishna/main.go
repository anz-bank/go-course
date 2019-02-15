package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	mapStorage := NewMapStore()
	syncStorage := NewSyncStore()

	puppy := func() Puppy {
		return Puppy{
			ID:     11,
			Breed:  "Sheep herder",
			Colour: "Brown",
			Value:  1000,
		}
	}

	//Map store
	newPup := puppy()
	err := mapStorage.CreatePuppy(&newPup)

	if err == nil {
		pup, err := mapStorage.ReadPuppy(newPup.ID)
		if err == nil {
			fmt.Fprintln(out, *pup)
		}

		ok, _ := mapStorage.DeletePuppy(pup.ID)
		if ok {
			_, err = mapStorage.DeletePuppy(11)
			fmt.Fprintln(out, err)
		}
	}

	//Sync.map store
	newPup = puppy()
	err = syncStorage.CreatePuppy(&newPup)
	if err == nil {
		pup, err := syncStorage.ReadPuppy(11)
		if err == nil {
			fmt.Fprintln(out, *pup)
		}

		ok, _ := syncStorage.DeletePuppy(pup.ID)
		if ok {
			_, err = mapStorage.DeletePuppy(11)
			fmt.Fprintln(out, err)
		}
	}
}
