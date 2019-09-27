package main

import (
	"fmt"
)

type MapStore struct {
	ms     map[int]Puppy
	currID int
}

func NewMapStore() *MapStore {
	newmapstore := MapStore{}
	newmapstore.ms = map[int]Puppy{}
	return &newmapstore
}

// CreatePuppy takes a  user provided puppy, and creates a new Puppy in the store
// puppy ID is updated with the next ID in sequence
// returns the ID generated for the new Puppy.
func (m *MapStore) CreatePuppy(puppy *Puppy) int {
	m.currID++
	puppy.ID = m.currID
	m.ms[m.currID] = *puppy
	return puppy.ID
}

// UpdatePuppy overrides an existing puppy in the store,
// returns an error if id is not found or does not match the puppy ID.
func (m *MapStore) UpdatePuppy(p *Puppy) error {
	if _, ok := m.ms[p.ID]; !ok {
		return fmt.Errorf("puppy ID %d to update does not exist in the map", p.ID)
	}
	m.ms[p.ID] = *p
	return nil
}

// DeletePuppy deletes an existing puppy from the store.
func (m *MapStore) DeletePuppy(id int) error {
	if _, ok := m.ms[id]; !ok {
		return fmt.Errorf("puppy ID %d does not exist in the map", id)
	}
	delete(m.ms, id)
	return nil
}

// ReadPuppy reads an existing puppy from the store.
func (m *MapStore) ReadPuppy(id int) (*Puppy, error) {
	puppy, ok := m.ms[id]
	if !ok {
		return nil, fmt.Errorf("puppy ID %d does not exist in the map", id)
	}
	return &puppy, nil
}
