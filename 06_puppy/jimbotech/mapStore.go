package main

import (
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

// Storer defines standard CRUD operations for Puppy
type Storer interface {
	CreatePuppy(p *Puppy) (uint32, error)
	ReadPuppy(ID uint32) (*Puppy, error)
	UpdatePuppy(ID uint32, Puppy *Puppy) error
	DeletePuppy(ID uint32)
}

// Puppy stores puppy details
type Puppy struct {
	ID     uint32
	Breed  string
	Colour string
	Value  string
}

// MapStore stores puppies
type MapStore map[uint32]*Puppy

// SyncMapStore stores puppies threadsafe
type SyncMapStore struct {
	sync.Map
}

type mapCheck interface {
	length() int
}

// length used for testing
func (s MapStore) length() int {
	return len(s)
}

func (s *SyncMapStore) length() int {
	var length int
	s.Range(func(key interface{}, value interface{}) bool {
		length++
		return true
	})
	return length
}

// CreatePuppy add a puppy to storage
func (s MapStore) CreatePuppy(p *Puppy) (uint32, error) {
	if err := checkCreation(s); err != nil {
		return 0, err
	}
	p.ID = uuid.New().ID()
	s[p.ID] = p
	return p.ID, nil
}

// CreatePuppy threadsafe adding a puppy to storage
func (s *SyncMapStore) CreatePuppy(p *Puppy) (uint32, error) {
	p.ID = uuid.New().ID()
	s.Store(p.ID, p)
	return p.ID, nil
}

// ReadPuppy retrieve your puppy
func (s MapStore) ReadPuppy(id uint32) (*Puppy, error) {
	if err := checkCreation(s); err != nil {
		return nil, err
	}
	val, found := s[id]
	if !found {
		return nil, fmt.Errorf("no puppy with ID %v found", id)
	}
	return val, nil
}

// ReadPuppy threadsafe retrieval of your puppy
func (s *SyncMapStore) ReadPuppy(id uint32) (*Puppy, error) {
	val, found := s.Load(id)
	if !found {
		return nil, fmt.Errorf("no puppy with ID %v found", id)
	}
	return val.(*Puppy), nil
}

// UpdatePuppy update your puppy store
func (s MapStore) UpdatePuppy(id uint32, puppy *Puppy) error {
	if err := checkCreation(s); err != nil {
		return err
	}
	if res := s[id]; res == nil {
		return fmt.Errorf("no puppy with ID %v found", id)
	}
	s[id] = puppy
	return nil
}

// UpdatePuppy threadsafe update your puppy store
func (s *SyncMapStore) UpdatePuppy(id uint32, puppy *Puppy) error {
	_, found := s.Load(id)
	if found {
		s.Store(id, puppy)
		return nil
	}
	return fmt.Errorf("no puppy with ID %v found", id)
}

// DeletePuppy remove the puppy from store
func (s MapStore) DeletePuppy(id uint32) {
	if err := checkCreation(s); err != nil {
		return
	}
	delete(s, id)
}

// DeletePuppy threadsafe removal of the puppy from store
func (s *SyncMapStore) DeletePuppy(id uint32) {
	s.Delete(id)
}

// NewPuppyStorer constructor creates the map
func NewPuppyStorer() MapStore {
	return MapStore{}
}

// checkCreation might not be required if it is not possible
// to create a Mapstore without calling the above constructor
func checkCreation(s MapStore) error {
	if s == nil {
		return errors.New("store not created, call initializing constructor first")
	}
	return nil
}
