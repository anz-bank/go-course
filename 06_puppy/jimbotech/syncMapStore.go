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
	sp := p
	s.Store(p.ID, sp)
	return p.ID, nil
}

// ReadPuppy threadsafe retrieval of your puppy.
func (s *SyncMapStore) ReadPuppy(id uint32) (*Puppy, error) {
	val, found := s.Load(id)
	if !found {
		return nil, fmt.Errorf("no puppy with ID %v found", id)
	}
	return val.(*Puppy), nil
}

// UpdatePuppy threadsafe update your puppy store.
func (s *SyncMapStore) UpdatePuppy(id uint32, puppy *Puppy) error {
	_, found := s.Load(id)
	if found {
		s.Store(id, puppy)
		return nil
	}
	return fmt.Errorf("no puppy with ID %v found", id)
}

// DeletePuppy threadsafe removal of the puppy from store.
func (s *SyncMapStore) DeletePuppy(id uint32) error {
	s.Delete(id)
	return nil
}
