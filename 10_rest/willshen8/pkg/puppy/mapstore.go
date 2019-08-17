package puppy

import (
	"strconv"
)

type MapStore struct {
	ms     map[uint32]Puppy
	nextID uint32
}

// NewMapStore initialise a new MapStore
func NewMapStore() *MapStore {
	return &MapStore{ms: map[uint32]Puppy{}}
}

// CreatePuppy create a new puppy and store in mapStore.
func (m *MapStore) CreatePuppy(p *Puppy) (uint32, error) {
	i, err := strconv.Atoi(p.Value)
	if err != nil {
		return 0, Errorf(ErrInvalidInput, ErrInvalidInput.String())
	}
	if i < 0 {
		return 0, Errorf(ErrInvalidInput, ErrInvalidInput.String())
	}

	m.nextID++
	p.ID = m.nextID
	m.ms[p.ID] = *p
	return p.ID, nil
}

// ReadPuppy read a puppy given its id.
// It returns the pointer to that puppy.
func (m MapStore) ReadPuppy(id uint32) (*Puppy, error) {
	if p, ok := m.ms[id]; ok {
		return &p, nil
	}
	return nil, Errorf(ErrNotFound, "Puppy ID (%v) not found", id)
}

// UpdatePuppy updates the store with key of id with the new puppy.
func (m MapStore) UpdatePuppy(id uint32, p *Puppy) error {
	if _, ok := m.ms[id]; !ok {
		return Errorf(ErrNotFound, "Puppy ID can't be found, update operation failed")
	}
	i, err := strconv.Atoi(p.Value)
	if err != nil {
		return Errorf(ErrInvalidInput, "Puppy value is not recognised")
	} else if i < 0 {
		return Errorf(ErrInvalidInput, "Puppy value can't be negative")
	}
	p.ID = id
	m.ms[id] = *p
	return nil
}

// DeletePuppy delete the puppy given the id.
func (m MapStore) DeletePuppy(id uint32) error {
	if _, ok := m.ms[id]; ok {
		delete(m.ms, id)
		return nil
	}
	return Errorf(ErrNotFound, "Puppy ID can't be found, delete operation failed")
}
