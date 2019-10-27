package main

import (
	"errors"

	"github.com/google/uuid"
)

// MapStore stores puppies.
type MapStore map[int32]*Puppy

// length used for testing.
func (s MapStore) length() int {
	return len(s)
}

// ErrNotConstructed returned if the interface was called without
// first constructing the underlaying structure.
var ErrNotConstructed = errors.New("store not created")

// CreatePuppy add a puppy to storage
// but will modify the member ID.
func (s MapStore) CreatePuppy(p *Puppy) (int32, error) {
	if s == nil {
		return 0, ErrNotConstructed
	}

	p.ID = int32(uuid.New().ID()) & 0x00FF
	sp := *p
	s[p.ID] = &sp
	return p.ID, nil
}

// ReadPuppy retrieve your puppy.
func (s MapStore) ReadPuppy(id int32) (*Puppy, error) {
	if s == nil {
		return nil, ErrNotConstructed
	}
	if id < 0 {
		return nil, ErrValueBelowZero
	}
	val, found := s[id]
	if !found {
		return nil, ErrIDNotFound
	}
	return val, nil
}

// UpdatePuppy update your puppy store.
func (s MapStore) UpdatePuppy(id int32, puppy *Puppy) error {
	_, err := s.ReadPuppy(id)
	if err != nil {
		return err
	}
	puppy.ID = id
	sp := *puppy
	s[id] = &sp
	return nil
}

// DeletePuppy remove the puppy from store.
func (s MapStore) DeletePuppy(id int32) error {
	_, err := s.ReadPuppy(id)
	if err != nil {
		return err
	}
	delete(s, id)
	return nil
}

// NewMapStore constructor creates the map.
func NewMapStore() MapStore {
	return MapStore{}
}
