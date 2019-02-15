package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var out io.Writer = os.Stdout

/**
* Storer implementation in below scenarios
* Dynamic Map ID: MapStore implementation of Storer backed by a map
* Static Map ID: SyncStore implementation of Storer backed by a sync.Map
 */
//Sync store
type syncStore struct {
	sync.Mutex
	sync.Map
}

func newSyncStore() *syncStore {
	return &syncStore{}
}

func (s *syncStore) createPuppy(in Puppy) uint {
	s.Lock()
	defer s.Unlock()
	s.Store(in.id, in)
	return in.id
}

func (s *syncStore) readPuppy(id uint) Puppy {
	pd, ok := s.Load(id)
	if !ok {
		return Puppy{}
	}
	p, _ := pd.(Puppy)
	return p
}

func (s *syncStore) updatePuppy(id uint, in Puppy) {
	s.Store(in.id, in)
}

func (s *syncStore) deletePuppy(id uint) {
	s.Delete(id)
}

// Map Store
type mapStore struct {
	ms    map[uint]Puppy
	mapID uint
}

func newMapStore() *mapStore {
	return &mapStore{ms: make(map[uint]Puppy)}
}

func (m *mapStore) createPuppy(in Puppy) uint {
	m.mapID++
	in.id = m.mapID
	m.ms[in.id] = in
	return m.mapID
}

func (m *mapStore) readPuppy(id uint) Puppy {
	p, ok := m.ms[id]
	if !ok {
		return Puppy{}
	}
	return p
}

func (m *mapStore) updatePuppy(id uint, in Puppy) {
	m.ms[id] = in
}

func (m *mapStore) deletePuppy(id uint) {
	delete(m.ms, id)
}

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
