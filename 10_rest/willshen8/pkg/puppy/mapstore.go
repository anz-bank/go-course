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
	if i, err := strconv.Atoi(p.Value); err == nil {
		if i < 0 {
			return 0, &Error{
				Message: PuppyValueLessThanZero,
				Code:    NegativeValue,
			}
		}
	} else {
		return 0, &Error{
			Message: InvalidPuppyValue,
			Code:    ErrorValueFormat,
		}
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
	return nil, &Error{
		Message: PuppyValueNotFound,
		Code:    NonExistentPuppy,
	}
}

// UpdatePuppy updates the store with key of id with the new puppy.
// It returns a boolean whether the operation is successful or not.
func (m MapStore) UpdatePuppy(id uint32, p *Puppy) (bool, error) {
	if _, ok := m.ms[id]; !ok {
		return false, &Error{
			Message: InvliadIDForUpdate,
			Code:    NonExistentPuppy,
		}
	}
	p.ID = id
	m.ms[id] = *p
	return true, nil
}

// DeletePuppy delete the puppy given the id.
// It returns true/success or false/unsuccessful.
func (m MapStore) DeletePuppy(id uint32) (bool, error) {
	if _, ok := m.ms[id]; ok {
		delete(m.ms, id)
		return true, nil
	}
	return false, &Error{
		Message: InvalidIDForDelete,
		Code:    NonExistentPuppy,
	}
}
