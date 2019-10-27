package main

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

// MapStore stores puppies.
type MapStore struct {
	mux    sync.Mutex
	pstore map[int32]*Puppy
}

// length used for testing.
func (s *MapStore) length() int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return len(s.pstore)
}

// ErrNotConstructed returned if the interface was called without
// first constructing the underlaying structure.
var ErrNotConstructed = errors.New("store not created")

// CreatePuppy add a puppy to storage
// but will modify the member ID.
func (s *MapStore) CreatePuppy(p *Puppy) (int32, error) {
	if s.pstore == nil {
		return 0, ErrNotConstructed
	}
	s.mux.Lock()
	defer s.mux.Unlock()
	p.ID = int32(uuid.New().ID()) & 0x00FF
	sp := *p
	s.pstore[p.ID] = &sp
	return p.ID, nil
}

// ReadPuppy retrieve your puppy.
func (s *MapStore) ReadPuppy(id int32) (*Puppy, error) {
	if s.pstore == nil {
		return nil, ErrNotConstructed
	}
	if id < 0 {
		return nil, ErrValueBelowZero
	}
	s.mux.Lock()
	defer s.mux.Unlock()
	val, found := s.pstore[id]
	if !found {
		return nil, ErrIDNotFound
	}
	return val, nil
}

// UpdatePuppy update your puppy store.
func (s *MapStore) UpdatePuppy(id int32, puppy *Puppy) error {
	_, err := s.ReadPuppy(id)
	if err != nil {
		return err
	}
	s.mux.Lock()
	defer s.mux.Unlock()
	puppy.ID = id
	sp := *puppy
	s.pstore[id] = &sp
	return nil
}

// DeletePuppy remove the puppy from store.
func (s *MapStore) DeletePuppy(id int32) error {
	_, err := s.ReadPuppy(id)
	if err != nil {
		return err
	}
	s.mux.Lock()
	defer s.mux.Unlock()
	delete(s.pstore, id)
	return nil
}

// NewMapStore constructor creates the map.
func NewMapStore() *MapStore {
	var m MapStore
	m.pstore = make(map[int32]*Puppy)
	return &m
}
