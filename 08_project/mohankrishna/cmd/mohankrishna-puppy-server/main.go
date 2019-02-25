package main

import (
	"fmt"
	"io"
	"os"

	types "github.com/anz-bank/go-training/08_project/mohankrishna/pkg/mohankrishna-puppy"
	"github.com/anz-bank/go-training/08_project/mohankrishna/pkg/mohankrishna-puppy/store"
)

var out io.Writer = os.Stdout

func main() {
	mapStorage := store.NewMapStore()
	syncStorage := store.NewSyncStore()

	puppy := func() types.Puppy {
		return types.Puppy{
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
		err = mapStorage.DeletePuppy(pup.ID)
		if err == nil {
			err = mapStorage.DeletePuppy(11)
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
		err = syncStorage.DeletePuppy(pup.ID)
		if err == nil {
			err = mapStorage.DeletePuppy(11)
			fmt.Fprintln(out, err)
		}
	}
}
