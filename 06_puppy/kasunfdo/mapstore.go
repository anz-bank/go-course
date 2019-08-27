package main

import (
	"fmt"
)

// MapStore is map based implementation of Storer interface
type MapStore struct {
	store  map[uint64]Puppy
	nextID uint64
}

// NewMapStore creates a new MapStore
func NewMapStore() *MapStore {
	return &MapStore{nextID: 1, store: map[uint64]Puppy{}}
}

// CreatePuppy inserts given puppy in store and returns given id to the puppy
func (m *MapStore) CreatePuppy(puppy Puppy) uint64 {
	puppy.ID = m.nextID
	m.nextID++
	m.store[puppy.ID] = puppy
	return puppy.ID
}

// ReadPuppy reads puppy with given id from the store
func (m *MapStore) ReadPuppy(id uint64) (Puppy, error) {
	puppy, exists := m.store[id]

	if !exists {
		return Puppy{}, fmt.Errorf("no puppy with id: %d", id)
	}
	return puppy, nil
}

// UpdatePuppy update puppy with given id in store if puppy exists
// Returns nil if puppy with given id exists. Otherwise returns an error
func (m *MapStore) UpdatePuppy(puppy Puppy) error {
	if _, exists := m.store[puppy.ID]; !exists {
		return fmt.Errorf("no puppy with id: %d", puppy.ID)
	}

	m.store[puppy.ID] = puppy
	return nil
}

// DeletePuppy deletes puppy with given id from store
// Returns nil if puppy with given id exists. Otherwise returns an error
func (m *MapStore) DeletePuppy(id uint64) error {
	if _, exists := m.store[id]; !exists {
		return fmt.Errorf("no puppy with id: %d", id)
	}

	delete(m.store, id)
	return nil
}
