package store

import (
	"fmt"

	puppy "github.com/anz-bank/go-course/08_project/n0npax/pkg/puppy"
)

// NewMemStore creates new storer for map
func NewMemStore() MemStore {
	return MemStore{}
}

// CreatePuppy creates puppy
func (m MemStore) CreatePuppy(p *puppy.Puppy) (int, error) {
	if p.Value < 0 {
		return -1, puppy.ErrInvalidInput(puppy.InvalidInputMsg)
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
	return nil, puppy.ErrNotFound(fmt.Sprintf(puppy.PuppyNotFoundMsg, id))
}

// UpdatePuppy updates puppy
func (m MemStore) UpdatePuppy(id int, p *puppy.Puppy) error {
	if p.Value < 0 {
		return puppy.ErrInvalidInput(puppy.InvalidInputMsg)
	}
	if id != p.ID {
		return puppy.ErrInvalidInput(puppy.CorruptedIDMsg)
	}
	if _, ok := m[id]; !ok {
		return puppy.ErrNotFound(fmt.Sprintf(puppy.PuppyNotFoundMsg, id))
	}
	m[id] = *p
	return nil
}

// DeletePuppy deletes puppy
func (m MemStore) DeletePuppy(id int) (bool, error) {
	if _, ok := m[id]; !ok {
		return false, puppy.ErrNotFound(fmt.Sprintf(puppy.PuppyNotFoundMsg, id))
	}
	delete(m, id)
	return true, nil
}
