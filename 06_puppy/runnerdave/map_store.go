package main

import "fmt"

// MapStore represents a simple map storage for the Puppy store
type MapStore struct {
	puppies map[uint16]Puppy
}

// NewMapStore creates a new in-memory store with map intialised
func NewMapStore() *MapStore {
	var m MapStore
	m.puppies = make(map[uint16]Puppy)
	return &m
}

// CreatePuppy saves new puppy if not in store, if it is already returns error
func (m *MapStore) CreatePuppy(p *Puppy) error {
	if _, ok := m.puppies[p.ID]; ok {
		return fmt.Errorf("puppy with id %d already exists", p.ID)
	}
	m.puppies[p.ID] = *p
	return nil
}

// ReadPuppy reads store by Puppy ID
func (m *MapStore) ReadPuppy(id uint16) (Puppy, error) {
	if puppy, ok := m.puppies[id]; ok {
		return puppy, nil
	}
	return Puppy{}, fmt.Errorf("puppy not found")
}

// UpdatePuppy updates puppy with new value if ID present otherwise creates new Puppy
func (m *MapStore) UpdatePuppy(id uint16, p *Puppy) error {
	m.puppies[id] = *p
	return nil
}

// DeletePuppy deletes a puppy by id from the store
func (m *MapStore) DeletePuppy(id uint16) (bool, error) {
	if _, ok := m.puppies[id]; ok {
		delete(m.puppies, id)
		return true, nil
	}
	return false, fmt.Errorf("puppy with id %d not found", id)
}
