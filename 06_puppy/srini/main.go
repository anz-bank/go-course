package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	labrodar := Puppy{"labrodar", "Red", 34343.43}
	poodle := Puppy{"Poodle", "White", 5000.43}

	//create stores
	mapStore := GetMapStore()
	syncStore := GetSyncStore()

	//Create a puppy
	mid := mapStore.CreatePuppy(&labrodar)
	fmt.Fprintln(out, fmt.Sprintf("Created a %v puppy with ID %d in Map Store", labrodar.breed, mid))

	sid := syncStore.CreatePuppy(&poodle)
	fmt.Fprintln(out, fmt.Sprintf("Created a %v puppy with ID %d in sync store", poodle.breed, sid))

	//Read the puppy
	puppyL, err := mapStore.ReadPuppy(1)
	if err == nil {
		fmt.Fprintln(out, fmt.Sprintf("Read a puppy from map store %v", *puppyL))
	} else {
		fmt.Fprintln(out, fmt.Sprintf("Error: %v", err))
	}

	puppyP, err := syncStore.ReadPuppy(1)
	if err == nil {
		fmt.Fprintln(out, fmt.Sprintf("Read a puppy from sync store %v", *puppyP))
	} else {
		fmt.Fprintln(out, fmt.Sprintf("Error: %v", err))
	}
	//Update the puppy color to black
	labrodar.colour = "black"
	mapStore.UpdatePuppy(1, &labrodar)
	if err == nil {
		fmt.Fprintln(out, fmt.Sprintf("Update the puppy color in map store %v", labrodar))
	} else {
		fmt.Fprintln(out, fmt.Sprintf("Error: %v", err))
	}

	//Update the puppy color to black
	poodle.colour = "black"
	err = syncStore.UpdatePuppy(1, &poodle)
	if err == nil {
		fmt.Fprintln(out, fmt.Sprintf("Update the puppy color in sync store %v", poodle))
	} else {
		fmt.Fprintln(out, fmt.Sprintf("Error: %v", err))
	}
	//Deletethe puppy
	err = mapStore.DeletePuppy(1)
	if err == nil {
		fmt.Fprintln(out, fmt.Sprintf("Deleted the puppy from map store"))
	} else {
		fmt.Fprintln(out, fmt.Sprintf("Error: %v", err))
	}

	err = syncStore.DeletePuppy(1)
	if err == nil {
		fmt.Fprintln(out, fmt.Sprintf("Deleted the puppy from sync store"))
	} else {
		fmt.Fprintln(out, fmt.Sprintf("Error: %v", err))
	}

}
