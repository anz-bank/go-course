package main

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// Storer defines standard CRUD operations for Puppy
type Storer interface {
	CreatePuppy(p *Puppy) (uint32, error)
	ReadPuppy(ID uint32) (*Puppy, error)
	UpdatePuppy(ID uint32, Puppy *Puppy) error
	DeletePuppy(ID uint32) error
}

// Puppy stores puppy details.
type Puppy struct {
	ID     uint32
	Breed  string
	Colour string
	Value  string
}

// MapStore stores puppies.
type MapStore map[uint32]*Puppy

type mapCheck interface {
	length() int
}

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
	return val, nil
}

// UpdatePuppy update your puppy store.
func (s MapStore) UpdatePuppy(id uint32, puppy *Puppy) error {
	if s == nil {
		return ErrNotConstructed
	}
	if res := s[id]; res == nil {
		return fmt.Errorf("no puppy with ID %v found", id)
	}
	s[id] = puppy
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
