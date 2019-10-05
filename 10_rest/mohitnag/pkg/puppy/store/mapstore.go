package store

import (
	"strconv"
	"sync"

	"github.com/anz-bank/go-course/10_rest/mohitnag/pkg/puppy"
)

// MapStore stores Puppy details with Puppy Id as Key and Puppy  as value
type MapStore struct {
	m map[uint32]puppy.Puppy
	sync.Mutex
}

func NewMapStore() *MapStore {
	return &MapStore{m: map[uint32]puppy.Puppy{}}
}

// CreatePuppy creates a Puppy in mapstore
func (m *MapStore) CreatePuppy(p puppy.Puppy) error {
	val, _ := strconv.Atoi(p.Value)
	if val < 0 {
		return puppy.ErrorF(puppy.Invalid, "puppy with value less than 0 not allowed")
	}
	m.Lock()
	defer m.Unlock()
	if _, ok := m.m[p.ID]; ok {
		return puppy.ErrorF(puppy.Duplicate, "puppy with Id %d already exists", p.ID)
	}
	m.m[p.ID] = p
	return nil
}

// ReadPuppy reads a Puppy from mapstore
func (m *MapStore) ReadPuppy(id uint32) (puppy.Puppy, error) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.m[id]; !ok {
		return puppy.Puppy{}, puppy.ErrorF(puppy.NotFound, "puppy with Id %d does not exists", id)
	}
	return m.m[id], nil
}

// UpdatePuppy updates a Puppy in mapstore
func (m *MapStore) UpdatePuppy(p puppy.Puppy) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.m[p.ID]; !ok {
		return puppy.ErrorF(puppy.NotFound, "puppy with Id %d does not exists", p.ID)
	}
	m.m[p.ID] = p
	return nil
}

// DeletePuppy deletes a Puppy from mapstore
func (m *MapStore) DeletePuppy(id uint32) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.m[id]; !ok {
		return puppy.ErrorF(puppy.NotFound, "puppy with Id %d does not exists", id)
	}
	delete(m.m, id)
	return nil
}
