package main

import (
	"fmt"
)

func NewMemStore() MemStore {
	return MemStore{}
}

func (m MemStore) CreatePuppy(puppy *Puppy) int {
	id := len(m)
	m[id] = puppy
	return id
}

func (m MemStore) ReadPuppy(id int) (*Puppy, error) {
	if puppy, ok := m[id]; ok {
		return puppy, nil
	}
	return nil, fmt.Errorf("puppy with ID: %d does not exist", id)
}

func (m MemStore) UpdatePuppy(id int, puppy *Puppy) error {
	if _, ok := m[id]; !ok {
		return fmt.Errorf("puppy with ID: %d does not exist", id)
	}
	if id != puppy.ID {
		return fmt.Errorf("puppy ID corrupted")
	}
	m[id] = puppy
	return nil
}

func (m MemStore) DeletePuppy(id int) (bool, error) {
	if _, ok := m[id]; !ok {
		return false, fmt.Errorf("puppy with ID: %d does not exist", id)
	}
	delete(m, id)
	return true, nil
}
