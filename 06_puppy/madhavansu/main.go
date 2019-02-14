package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var out io.Writer = os.Stdout

//Sync store
type syncStore struct {
	sync.Mutex
	sync.Map
}

func newSyncStore() *syncStore {
	return &syncStore{}
}

func (s *syncStore) createPuppy(in Puppy) {
	s.Store(in.id, in)
}

func (s *syncStore) readPuppy(id uint32) Puppy {
	pd, ok := s.Load(id)
	if !ok {
		fmt.Printf("No puppy exists\n")
	}
	p, _ := pd.(Puppy)
	return p
}

func (s *syncStore) updatePuppy(id uint32, in Puppy) {
	s.Store(in.id, in)
}

func (s *syncStore) deletePuppy(id uint32) {
	s.Delete(id)
}

// Map Store
type mapStore struct {
	ms map[uint32]Puppy
}

func newMapStore() *mapStore {
	return &mapStore{make(map[uint32]Puppy)}
}

func (m *mapStore) createPuppy(in Puppy) {
	m.ms[in.id] = in
}

func (m *mapStore) readPuppy(id uint32) Puppy {
	p, ok := m.ms[id]
	if !ok {
		fmt.Printf("No puppy exists\n")
	}
	return p
}

func (m *mapStore) updatePuppy(id uint32, in Puppy) {
	m.ms[id] = in
}

func (m *mapStore) deletePuppy(id uint32) {
	_, ok := m.ms[id]
	if ok {
		delete(m.ms, id)
	}
}

func main() {
	// Sync storer
	var s Storer = newSyncStore()
	fmt.Fprintln(out, "~~~~~~~~~~")
	fmt.Fprintln(out, "Sync Store")
	fmt.Fprintln(out, "~~~~~~~~~~")
	// Create/Read puppy
	s.createPuppy(Puppy{101, "Poodle", "red", 18000})
	fmt.Fprintln(out, "101 : ", s.readPuppy(101))
	// Delete puppy
	s.createPuppy(Puppy{102, "Bulldog", "brownish", 9999})
	s.deletePuppy(102)
	fmt.Fprintln(out, "102 : ", s.readPuppy(102))
	// Update puppy
	s.createPuppy(Puppy{103, "Labrador Retriever", "purple", 987})
	s.updatePuppy(103, Puppy{103, "German Shepherd", "red", 4533})
	fmt.Fprintln(out, "103 : ", s.readPuppy(103))

	// Map storer
	var m Storer = newMapStore()
	fmt.Fprintln(out, "~~~~~~~~~")
	fmt.Fprintln(out, "Map Store")
	fmt.Fprintln(out, "~~~~~~~~~")
	// Create/Read puppy
	m.createPuppy(Puppy{104, "Pug", "brown", 0.91})
	fmt.Fprintln(out, "104 : ", m.readPuppy(104))
	// Delete puppy
	m.createPuppy(Puppy{105, "Beagle", "yellowish", 1233})
	m.deletePuppy(105)
	fmt.Fprintln(out, "105 : ", m.readPuppy(105))
	// Update puppy
	m.createPuppy(Puppy{106, "Boxer", "black", 45000.98})
	m.updatePuppy(106, Puppy{106, "Beagle", "brown", 0.91})
	fmt.Fprint(out, "106 :  ", m.readPuppy(106))
}
