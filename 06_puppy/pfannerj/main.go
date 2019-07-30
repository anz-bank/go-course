package main

import (
	"fmt"
	"io"
	"os"
)

var mainout io.Writer = os.Stdout

func main() {
	mapStore := NewMapStore()
	puppy := Puppy{ID: 12345, Breed: "Labrador", Colour: "Brown", Value: 999.99}
	mapPuppyID, _ := mapStore.CreatePuppy(puppy)
	fmt.Fprintf(mainout, "Map Puppy Created with ID: %d\n", mapPuppyID)
	mapPuppy, _ := mapStore.ReadPuppy(mapPuppyID)
	fmt.Fprintf(mainout, "Map Puppy read: %v\n", mapPuppy)
	mapPuppyID, _ = mapStore.UpdatePuppy(mapPuppyID, mapPuppy)
	fmt.Fprintf(mainout, "Map Puppy updated: %d\n", mapPuppyID)
	mapErr := mapStore.DeletePuppy(mapPuppyID)
	fmt.Fprintf(mainout, "Map Puppy deleted! Error returned was: %v\n", mapErr)
	syncStore := NewSyncStore()
	syncPuppyID, _ := syncStore.CreatePuppy(puppy)
	fmt.Fprintf(mainout, "Sync Puppy Created with ID: %d\n", syncPuppyID)
	syncPuppy, _ := syncStore.ReadPuppy(syncPuppyID)
	fmt.Fprintf(mainout, "Sync Puppy read: %v\n", syncPuppy)
	syncPuppyID, _ = syncStore.UpdatePuppy(syncPuppyID, syncPuppy)
	fmt.Fprintf(mainout, "Sync Puppy updated: %d\n", syncPuppyID)
	syncErr := syncStore.DeletePuppy(syncPuppyID)
	fmt.Fprintf(mainout, "Sync Puppy deleted! Error returned was: %v\n", syncErr)
}
