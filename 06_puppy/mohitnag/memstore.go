package main

import "fmt"

// CreatePuppy creates a Puppy in memstore
func (m *MemStore) CreatePuppy(p Puppy) error {
	if _, ok := (*m)[p.ID]; !ok {
		(*m)[p.ID] = p
		return nil
	}
	return fmt.Errorf("puppy with Id %d already exists", p.ID)
}

// ReadPuppy reads a Puppy from memstore
func (m *MemStore) ReadPuppy(id uint32) (Puppy, error) {
	if _, ok := (*m)[id]; !ok {
		return Puppy{}, fmt.Errorf("puppy with Id %d does not exists", id)
	}
	return (*m)[id], nil
}

// UpdatePuppy updates a Puppy
func (m *MemStore) UpdatePuppy(p Puppy) error {
	if _, ok := (*m)[p.ID]; !ok {
		return fmt.Errorf("puppy with Id %d does not exists", p.ID)
	}
	(*m)[p.ID] = p
	return nil
}

// DeletePuppy deletes a Puppy
func (m *MemStore) DeletePuppy(id uint32) bool {
	if _, ok := (*m)[id]; !ok {
		return false
	}
	delete(*m, id)
	return true
}
