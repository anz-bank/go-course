package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

// SyncMapStore stores puppies threadsafe.
type SyncMapStore struct {
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
	s.Range(func(key interface{}, value interface{}) bool {
		length++
		return true
	})
	return length
}

// CreatePuppy threadsafe adding a puppy to storage
// but will modify the member ID.
func (s *SyncMapStore) CreatePuppy(p *Puppy) (uint32, error) {
	p.ID = uuid.New().ID()
	sp := *p
	s.Store(p.ID, &sp)
	return p.ID, nil
}

// ReadPuppy threadsafe retrieval of your puppy.
func (s *SyncMapStore) ReadPuppy(id uint32) (*Puppy, error) {
	val, found := s.Load(id)
	if !found {
		return nil, fmt.Errorf("no puppy with ID %v found", id)
	}
	retPup := *val.(*Puppy)
	return &retPup, nil
}

// UpdatePuppy threadsafe update your puppy store.
func (s *SyncMapStore) UpdatePuppy(id uint32, puppy *Puppy) error {
	_, found := s.Load(id)
	if !found {
		return fmt.Errorf("no puppy with ID %v found", id)
	}
	puppy.ID = id
	sp := *puppy
	s.Store(id, &sp)
	return nil
}

// DeletePuppy threadsafe removal of the puppy from store.
func (s *SyncMapStore) DeletePuppy(id uint32) error {
	s.Delete(id)
	return nil
}
