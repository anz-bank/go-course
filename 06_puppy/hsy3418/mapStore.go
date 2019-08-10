package main

import "fmt"

// MapStore is a implementation of storer for the storage of puppies
type MapStore struct {
	puppies map[int32]Puppy
	nextID  int32
}

// NewMapStore creates a new mapStore
func NewMapStore() *MapStore {
	return &MapStore{
		puppies: map[int32]Puppy{},
	}
}

// CreatePuppy adds a nuw puppy to the puppy store and returns the id for the puppy
func (m *MapStore) CreatePuppy(puppy Puppy) int32 {
	puppy.ID = m.nextID
	m.nextID++
	m.puppies[puppy.ID] = puppy
	return puppy.ID
}

// ReadPuppy retrieves the puppy for a given id from puppies store
func (m *MapStore) ReadPuppy(id int32) (Puppy, error) {
	if _, ok := m.puppies[id]; !ok {
		return Puppy{}, fmt.Errorf("puppy with %d ID does not exist", id)
	}
	return m.puppies[id], nil
}

//UpdatePuppy updates the puppy for the given id
func (m *MapStore) UpdatePuppy(puppy Puppy) error {
	if _, ok := m.puppies[puppy.ID]; !ok {
		return fmt.Errorf("puppy with %d ID does not exist", puppy.ID)
	}
	m.puppies[puppy.ID] = puppy
	return nil
}

//DeletePuppy delete the puppy for the given id from puppies store
func (m *MapStore) DeletePuppy(id int32) error {
	if _, ok := m.puppies[id]; !ok {
		return fmt.Errorf("puppy with %d ID does not exist", id)
	}
	delete(m.puppies, id)
	return nil
}
