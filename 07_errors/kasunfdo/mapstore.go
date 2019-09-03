package main

import "fmt"

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
// Will return an error if value of puppy is negative
func (m *MapStore) CreatePuppy(puppy Puppy) (uint64, error) {
	if puppy.Value < 0 {
		return 0, NewError(ErrInvalid, "value of puppy is negative")
	}

	puppy.ID = m.nextID
	m.nextID++
	m.store[puppy.ID] = puppy
	return puppy.ID, nil
}

// ReadPuppy reads puppy with given id from the store
// Will return an error if puppy with given id does not exist
func (m *MapStore) ReadPuppy(id uint64) (Puppy, error) {
	puppy, exists := m.store[id]

	if !exists {
		return puppy, NewError(ErrNotFound, fmt.Sprintf("puppy with id: %v is not found", id))
	}
	return puppy, nil
}

// UpdatePuppy update puppy with given id in store if puppy exists
// Returns nil if puppy with given id exists. Otherwise returns an error
func (m *MapStore) UpdatePuppy(puppy Puppy) error {
	if puppy.Value < 0 {
		return NewError(ErrInvalid, "value of puppy is negative")
	}

	if _, exists := m.store[puppy.ID]; !exists {
		return NewError(ErrNotFound, fmt.Sprintf("puppy with id: %v is not found", puppy.ID))
	}

	m.store[puppy.ID] = puppy
	return nil
}

// DeletePuppy deletes puppy with given id from store
// Returns nil if puppy with given id exists. Otherwise returns an errorZ
func (m *MapStore) DeletePuppy(id uint64) error {
	if _, exists := m.store[id]; !exists {
		return NewError(ErrNotFound, fmt.Sprintf("puppy with id: %v is not found", id))
	}

	delete(m.store, id)
	return nil
}
