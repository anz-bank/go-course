package main

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// MapStore stores puppies.
type MapStore map[uint32]*Puppy

// length used for testing.
func (s MapStore) length() int {
	return len(s)
}

// ErrNotConstructed returned if the interface was called without
// first constructing the underlaying structure.
var ErrNotConstructed = errors.New("store not created")

// CreatePuppy add a puppy to storage
// but will modify the member ID.
func (s MapStore) CreatePuppy(p *Puppy) (uint32, error) {
	if s == nil {
		return 0, ErrNotConstructed
	}
	p.ID = uuid.New().ID()
	sp := *p
	s[p.ID] = &sp
	return p.ID, nil
}

// ReadPuppy retrieve your puppy.
func (s MapStore) ReadPuppy(id uint32) (*Puppy, error) {
	if s == nil {
		return nil, ErrNotConstructed
	}
	val, found := s[id]
	if !found {
		return nil, fmt.Errorf("no puppy with ID %v found", id)
	}
	retVal := *val
	return &retVal, nil
}

// UpdatePuppy update your puppy store.
func (s MapStore) UpdatePuppy(id uint32, puppy *Puppy) error {
	if s == nil {
		return ErrNotConstructed
	}
	if _, ok := s[id]; !ok {
		return fmt.Errorf("no puppy with ID %v found", id)
	}
	puppy.ID = id
	sp := *puppy
	s[id] = &sp
	return nil
}

// DeletePuppy remove the puppy from store.
func (s MapStore) DeletePuppy(id uint32) error {
	if s == nil {
		return ErrNotConstructed
	}
	delete(s, id)
	return nil
}

// NewMapStore constructor creates the map.
func NewMapStore() MapStore {
	return MapStore{}
}
