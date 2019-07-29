package main

import "fmt"

// CreatePuppy creates a Puppy in mapstore
func (m MapStore) CreatePuppy(p Puppy) error {
	_, ok := m[p.ID]
	if !ok {
		m[p.ID] = p
		return nil
	}
	return fmt.Errorf("puppy with Id %d already exists", p.ID)
}

// ReadPuppy reads a Puppy from mapstore
func (m MapStore) ReadPuppy(id uint32) (Puppy, error) {
	_, ok := m[id]
	if !ok {
		return Puppy{}, fmt.Errorf("puppy with Id %d does not exists", id)
	}
	return m[id], nil
}

// UpdatePuppy updates a Puppy in mapstore
func (m MapStore) UpdatePuppy(p Puppy) error {
	_, ok := m[p.ID]
	if !ok {
		return fmt.Errorf("puppy with Id %d does not exists", p.ID)
	}
	m[p.ID] = p
	return nil
}

// DeletePuppy deletes a Puppy from mapstore
func (m MapStore) DeletePuppy(id uint32) error {
	_, ok := m[id]
	if !ok {
		return fmt.Errorf("puppy with Id %d does not exists", id)
	}
	delete(m, id)
	return nil
}
