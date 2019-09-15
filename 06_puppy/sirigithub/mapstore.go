package main

import (
	"fmt"
)

type MapStore map[int]Puppy

func NewMapStore() MapStore {
	return MapStore{}
}

// CreatePuppy takes a  user provided puppy, and creates a new Puppy in the store
// returns the ID of the new Puppy.
func (m MapStore) CreatePuppy(puppy *Puppy) int {
	puppy.ID = len(m)
	m[puppy.ID] = *puppy
	return puppy.ID
}

// UpdatePuppy overrides an existing puppy in the store,
// returns an error if id is not found or does not match the puppy ID
func (m MapStore) UpdatePuppy(p *Puppy) error {
	if _, ok := m[p.ID]; !ok {
		return fmt.Errorf("puppy ID %d to update does not exist in the map", p.ID)
	}
	m[p.ID] = *p
	return nil
}

// DeletePuppy deletes an existing puppy from the store
func (m MapStore) DeletePuppy(id int) error {
	if _, ok := m[id]; ok {
		delete(m, id)
		return nil
	}
	return fmt.Errorf("puppy ID %d does not exist in the map", id)
}

// ReadPuppy reads an existing puppy from the store
func (m MapStore) ReadPuppy(id int) (*Puppy, error) {
	if puppy, ok := m[id]; ok {
		return &puppy, nil
	}
	return nil, fmt.Errorf("puppy ID %d does not exist in the map", id)
}
