package store

import (
	puppy "github.com/anz-bank/go-course/10_rest/runnerdave/pkg/puppy"
)

// MapStore represents a simple map storage for the Puppy store
type MapStore struct {
	puppies map[int16]puppy.Puppy
	nextID  int16
}

// NewMapStore creates a new in-memory store with map intialised
func NewMapStore() *MapStore {
	return &MapStore{
		nextID:  1,
		puppies: map[int16]puppy.Puppy{}}
}

// CreatePuppy saves new puppy if not in store, if it is already returns error
func (m *MapStore) CreatePuppy(p puppy.Puppy) error {
	if err := puppy.ValidateValue(p.Value); err != nil {
		return err
	}
	p.ID = m.nextID
	m.puppies[p.ID] = p
	m.nextID++
	return nil
}

// ReadPuppy reads store by Puppy ID
func (m *MapStore) ReadPuppy(id int16) (puppy.Puppy, error) {
	if puppy, ok := m.puppies[id]; ok {
		return puppy, nil
	}
	return puppy.Puppy{}, puppy.Errorf(puppy.ErrIDNotFound, "puppy with ID:%d not found", id)
}

// UpdatePuppy updates puppy with new value if ID present otherwise error
func (m *MapStore) UpdatePuppy(id int16, p *puppy.Puppy) error {
	if err := puppy.ValidateValue(p.Value); err != nil {
		return err
	}
	if _, ok := m.puppies[id]; !ok {
		return puppy.Errorf(puppy.ErrIDNotFound, "puppy with ID:%d not found", id)
	}
	p.ID = id
	m.puppies[id] = *p
	return nil
}

// DeletePuppy deletes a puppy by id from the store
func (m *MapStore) DeletePuppy(id int16) error {
	if _, ok := m.puppies[id]; ok {
		delete(m.puppies, id)
		return nil
	}
	return puppy.Errorf(puppy.ErrIDNotFound, "puppy with ID:%d not found", id)
}
