package store

import (
	"strconv"

	"github.com/anz-bank/go-course/09_json/mohitnag/pkg/puppy"
)

// MapStore stores Puppy details with Puppy Id as Key and Puppy  as value
type MapStore map[uint32]puppy.Puppy

// CreatePuppy creates a Puppy in mapstore
func (m MapStore) CreatePuppy(p puppy.Puppy) error {
	if _, ok := m[p.ID]; ok {
		return puppy.ErrorF(puppy.Duplicate, "puppy with Id %d already exists", p.ID)
	}
	val, _ := strconv.Atoi(p.Value)
	if val < 0 {
		return puppy.ErrorF(puppy.Invalid, "puppy with value less than 0 not allowed")
	}
	m[p.ID] = p
	return nil
}

// ReadPuppy reads a Puppy from mapstore
func (m MapStore) ReadPuppy(id uint32) (puppy.Puppy, error) {
	if _, ok := m[id]; !ok {
		return puppy.Puppy{}, puppy.ErrorF(puppy.NotFound, "puppy with Id %d does not exists", id)
	}
	return m[id], nil
}

// UpdatePuppy updates a Puppy in mapstore
func (m MapStore) UpdatePuppy(p puppy.Puppy) error {
	if _, ok := m[p.ID]; !ok {
		return puppy.ErrorF(puppy.NotFound, "puppy with Id %d does not exists", p.ID)
	}
	m[p.ID] = p
	return nil
}

// DeletePuppy deletes a Puppy from mapstore
func (m MapStore) DeletePuppy(id uint32) error {
	if _, ok := m[id]; !ok {
		return puppy.ErrorF(puppy.NotFound, "puppy with Id %d does not exists", id)
	}
	delete(m, id)
	return nil
}
