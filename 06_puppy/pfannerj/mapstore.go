package main

import "fmt"

// PuppyMap a map for storing puppies in memory.
type PuppyMap map[uint32](Puppy)

// MapStore struct for the puppy map & keeps track of the last puppy ID allocated.
type MapStore struct {
	puppyMap PuppyMap
	currID   uint32
}

// NewMapStore creates a new in-memory store with map intialised.
func NewMapStore() *MapStore {
	return &MapStore{puppyMap: PuppyMap{}}
}

// CreatePuppy adds a new puppy to the map store.
func (m *MapStore) CreatePuppy(p Puppy) (uint32, error) {
	m.currID++
	p.ID = m.currID //ensure the ID within p always matches the map store key (puppyID)
	m.puppyMap[m.currID] = p
	return p.ID, nil
}

//ReadPuppy gets a puppy from the map store with the given ID.
func (m *MapStore) ReadPuppy(puppyID uint32) (Puppy, error) {
	if puppyOut, ok := m.puppyMap[puppyID]; ok {
		return puppyOut, nil
	}
	return Puppy{}, fmt.Errorf("no puppy found with id %d", puppyID)
}

// UpdatePuppy modifies puppy data in the map store, either creating a new one or overwriting an old one.
func (m *MapStore) UpdatePuppy(puppyID uint32, p Puppy) (uint32, error) {
	if _, ok := m.puppyMap[puppyID]; !ok {
		m.currID++
		puppyID = m.currID
	}
	p.ID = puppyID //ensure the ID within p always matches the map store key (puppyID)
	m.puppyMap[puppyID] = p
	return p.ID, nil
}

// DeletePuppy deletes the puppy with the given ID from the map store.
func (m *MapStore) DeletePuppy(puppyID uint32) error {
	if _, ok := m.puppyMap[puppyID]; ok {
		delete(m.puppyMap, puppyID)
	}
	return nil
}
