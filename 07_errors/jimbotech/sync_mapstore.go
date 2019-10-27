package main

import (
	"sync"

	"github.com/google/uuid"
)

// SyncMapStore stores puppies threadsafe.
type SyncMapStore struct {
	mux sync.Mutex
	sync.Map
}

// length is not concorrency safe. As the go doc says:
// Range does not necessarily correspond to any consistent snapshot of the
// Map's contents: no key will be visited more than once, but if the value
// for any key is stored or deleted concurrently, Range may reflect any
// mapping for that key from any point during the Range call.
//
func (s *SyncMapStore) length() int {
	var length int
	s.mux.Lock()
	defer s.mux.Unlock()
	s.Range(func(key interface{}, value interface{}) bool {
		length++
		return true
	})
	return length
}

// CreatePuppy threadsafe adding a puppy to storage
// but will modify the member ID.
func (s *SyncMapStore) CreatePuppy(p *Puppy) (int32, error) {
	p.ID = int32(uuid.New().ID()) & 0X00FF
	sp := p
	s.mux.Lock()
	defer s.mux.Unlock()
	s.Store(p.ID, sp)
	return p.ID, nil
}

// ReadPuppy threadsafe retrieval of your puppy.
func (s *SyncMapStore) ReadPuppy(id int32) (*Puppy, error) {
	if id < 0 {
		return nil, ErrValueBelowZero
	}
	s.mux.Lock()
	defer s.mux.Unlock()
	val, found := s.Load(id)
	if !found {
		return nil, ErrIDNotFound
	}
	return val.(*Puppy), nil
}

// UpdatePuppy threadsafe update your puppy store.
func (s *SyncMapStore) UpdatePuppy(id int32, puppy *Puppy) error {
	_, err := s.ReadPuppy(id)
	if err != nil {
		return err
	}
	s.mux.Lock()
	defer s.mux.Unlock()
	puppy.ID = id
	sp := puppy
	s.Store(id, sp)
	return nil
}

// DeletePuppy threadsafe removal of the puppy from store.
func (s *SyncMapStore) DeletePuppy(id int32) error {
	_, err := s.ReadPuppy(id)
	if err != nil {
		return err
	}
	s.mux.Lock()
	defer s.mux.Unlock()
	s.Delete(id)
	return nil
}
