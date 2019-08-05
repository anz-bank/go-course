package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	puppy1 := Puppy{ID: 11, Breed: "Chihuahua", Colour: "Brown", Value: 12.30}
	puppy2 := Puppy{ID: 11, Breed: "Chihuahua", Colour: "Brown", Value: 10.30}

	mapStore := NewMapStore()
	mapCreateErr := mapStore.CreatePuppy(&puppy1)

	puppyMap, _ := mapStore.ReadPuppy(11)
	fmt.Fprintf(out, "Puppy created in map of Breed:%s, errors at creation:%v\n", puppyMap.Breed, mapCreateErr)

	syncStore := NewSyncStore()
	syncCreateErr := syncStore.CreatePuppy(&puppy1)
	syncUpdateErr := syncStore.UpdatePuppy(11, &puppy2)
	puppySync, _ := syncStore.ReadPuppy(11)
	fmt.Fprintf(out, "Puppy created in sync of Breed:%s, value updated to:%f, error at creation:%v, error in update:%v\n",
		puppySync.Breed, puppySync.Value, syncCreateErr, syncUpdateErr)
}
