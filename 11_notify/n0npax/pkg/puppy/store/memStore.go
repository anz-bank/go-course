package store

import (
	"sync"

	puppy "github.com/anz-bank/go-course/11_notify/n0npax/pkg/puppy"
)

// MemStore map based type for storing puppies data
type MemStore struct {
	m      map[int]puppy.Puppy
	nextID int
	sync.Mutex
}

// NewMemStore creates new storer for map
func NewMemStore() *MemStore {
	return &MemStore{m: map[int]puppy.Puppy{}}
}

// CreatePuppy creates puppy
func (m *MemStore) CreatePuppy(p *puppy.Puppy) (int, error) {
	if p.Value < 0 {
		return -1, puppy.Errorf(puppy.ErrCodeInvalidInput, "puppy value have to be positive number")
	}
	m.Lock()
	defer m.Unlock()
	id := m.nextID
	m.nextID++
	p.ID = id
	m.m[id] = *p
	return id, nil
}

// ReadPuppy reads puppy from backend
func (m *MemStore) ReadPuppy(id int) (*puppy.Puppy, error) {
	if puppy, ok := m.m[id]; ok {
		return &puppy, nil
	}
	return nil, puppy.Errorf(puppy.ErrCodeNotFound, "puppy with ID (%v) not found", id)
}

// UpdatePuppy updates puppy
func (m *MemStore) UpdatePuppy(id int, p *puppy.Puppy) error {
	if p.Value < 0 {
		return puppy.Errorf(puppy.ErrCodeInvalidInput, "puppy value have to be positive number")
	}
	if id != p.ID {
		return puppy.Errorf(puppy.ErrCodeInvalidInput, "ID is corrupted. Please ensure object ID matched provided ID")
	}
	if _, ok := m.m[id]; !ok {
		return puppy.Errorf(puppy.ErrCodeNotFound, "puppy with ID (%v) not found", id)
	}
	m.Lock()
	defer m.Unlock()
	m.m[id] = *p
	return nil
}

// DeletePuppy deletes puppy
func (m *MemStore) DeletePuppy(id int) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.m[id]; !ok {
		return puppy.Errorf(puppy.ErrCodeNotFound, "puppy with ID (%v) not found", id)
	}
	delete(m.m, id)
	return nil
}
