package store

import (
	"fmt"

	"github.com/anz-bank/go-course/08_project/kasunfdo/pkg/puppy"
)

// MapStore is map based implementation of Storer interface
type MapStore struct {
	store  map[uint64]puppy.Puppy
	nextID uint64
}

// NewMapStore creates a new MapStore
func NewMapStore() *MapStore {
	return &MapStore{nextID: 1, store: map[uint64]puppy.Puppy{}}
}

// CreatePuppy inserts given puppy in store and returns given id to the puppy
// Will return an error if value of puppy is negative
func (m *MapStore) CreatePuppy(p puppy.Puppy) (uint64, error) {
	if p.Value < 0 {
		return 0, puppy.NewError(puppy.ErrInvalid, "value of puppy is negative")
	}

	p.ID = m.nextID
	m.nextID++
	m.store[p.ID] = p
	return p.ID, nil
}

// ReadPuppy reads puppy with given id from the store
// Will return an error if puppy with given id does not exist
func (m *MapStore) ReadPuppy(id uint64) (puppy.Puppy, error) {
	p, exists := m.store[id]

	if !exists {
		return p, puppy.NewError(puppy.ErrNotFound, fmt.Sprintf("puppy with id: %v is not found", id))
	}
	return p, nil
}

// UpdatePuppy update puppy with given id in store if puppy exists
// Returns nil if puppy with given id exists. Otherwise returns an error
func (m *MapStore) UpdatePuppy(p puppy.Puppy) error {
	if p.Value < 0 {
		return puppy.NewError(puppy.ErrInvalid, "value of puppy is negative")
	}

	if _, exists := m.store[p.ID]; !exists {
		return puppy.NewError(puppy.ErrNotFound, fmt.Sprintf("puppy with id: %v is not found", p.ID))
	}

	m.store[p.ID] = p
	return nil
}

// DeletePuppy deletes puppy with given id from store
// Returns nil if puppy with given id exists. Otherwise returns an errorZ
func (m *MapStore) DeletePuppy(id uint64) error {
	if _, exists := m.store[id]; !exists {
		return puppy.NewError(puppy.ErrNotFound, fmt.Sprintf("puppy with id: %v is not found", id))
	}

	delete(m.store, id)
	return nil
}
