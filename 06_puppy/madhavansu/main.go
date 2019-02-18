package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

/**
* Storer implementation in below scenarios
* Dynamic Map ID: MapStore implementation of Storer backed by a map
* Static Map ID: SyncStore implementation of Storer backed by a sync.Map
 */
func main() {
	var puppyID uint
	// Sync storer
	var s Storer = newSyncStore()
	fmt.Fprintln(out, "~~~~~~~~~~")
	fmt.Fprintln(out, "Sync Store")
	fmt.Fprintln(out, "~~~~~~~~~~")
	// Create/Read puppy
	puppyID = s.createPuppy(Puppy{101, "Poodle", "red", "18000"})
	fmt.Fprintln(out, "101 : ", s.readPuppy(puppyID))
	// Delete puppy
	puppyID = s.createPuppy(Puppy{102, "Bulldog", "brownish", "9999"})
	s.deletePuppy(puppyID)
	fmt.Fprintln(out, "102 : ", s.readPuppy(puppyID))
	// Update puppy
	puppyID = s.createPuppy(Puppy{103, "Labrador Retriever", "purple", "987"})
	s.updatePuppy(puppyID, Puppy{puppyID, "German Shepherd", "red", "4533"})
	fmt.Fprintln(out, "103 : ", s.readPuppy(puppyID))

	// Map storer
	var m Storer = newMapStore()
	fmt.Fprintln(out, "~~~~~~~~~")
	fmt.Fprintln(out, "Map Store")
	fmt.Fprintln(out, "~~~~~~~~~")
	// Create/Read puppy
	puppyID = m.createPuppy(Puppy{104, "Pug", "brown", "0.91"})
	fmt.Fprintln(out, puppyID, " : ", m.readPuppy(puppyID))
	// Delete puppy
	puppyID = m.createPuppy(Puppy{105, "Beagle", "yellowish", "1233"})
	m.deletePuppy(puppyID)
	fmt.Fprintln(out, puppyID, " : ", m.readPuppy(puppyID))
	// Update puppy
	puppyID = m.createPuppy(Puppy{106, "Boxer", "black", "45000.98"})
	m.updatePuppy(puppyID, Puppy{puppyID, "Beagle", "brown", "0.91"})
	fmt.Fprintln(out, puppyID, " : ", m.readPuppy(puppyID))
}
