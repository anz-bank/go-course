package store

import (
	"fmt"

	puppy "github.com/anz-bank/go-course/08_project/n0npax/pkg/puppy"
)

// MemStore map based type for storing puppies data
type MemStore map[int]puppy.Puppy

// NewMemStore creates new storer for map
func NewMemStore() MemStore {
	return MemStore{}
}

// CreatePuppy creates puppy
func (m MemStore) CreatePuppy(p *puppy.Puppy) (int, error) {
	if p.Value < 0 {
		return -1, puppy.Errorf(puppy.ErrInvalidInputCode, "Puppy value have to be positive number")
	}
	id := len(m)
	m[id] = *p
	return id, nil
}

// ReadPuppy reads puppy from backend
func (m MemStore) ReadPuppy(id int) (*puppy.Puppy, error) {
	if puppy, ok := m[id]; ok {
		return &puppy, nil
	}
	return nil, puppy.Errorf(puppy.ErrNotFoundCode, fmt.Sprintf("Puppy with ID (%v) not found", id))
}

// UpdatePuppy updates puppy
func (m MemStore) UpdatePuppy(id int, p *puppy.Puppy) error {
	if p.Value < 0 {
		return puppy.Errorf(puppy.ErrInvalidInputCode, "Puppy value have to be positive number")
	}
	if id != p.ID {
		return puppy.Errorf(puppy.ErrInvalidInputCode, "ID is corrupted. Please ensure object ID matched provided ID")
	}
	if _, ok := m[id]; !ok {
		return puppy.Errorf(puppy.ErrNotFoundCode, fmt.Sprintf("Puppy with ID (%v) not found", id))
	}
	m[id] = *p
	return nil
}

// DeletePuppy deletes puppy
func (m MemStore) DeletePuppy(id int) (bool, error) {
	if _, ok := m[id]; !ok {
		return false, puppy.Errorf(puppy.ErrNotFoundCode, fmt.Sprintf("Puppy with ID (%v) not found", id))
	}
	delete(m, id)
	return true, nil
}
