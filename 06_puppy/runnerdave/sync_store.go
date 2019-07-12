package main

import (
	"fmt"
	"sync"
)

// SyncStore implementation of the Storer interface
type SyncStore struct {
	sync.Mutex
	sync.Map
}

// NewSyncStore creates a new in-memory store with map intialised
func NewSyncStore() *SyncStore {
	return &SyncStore{}
}

// CreatePuppy saves new puppy if not in store, if it is already returns error
func (s *SyncStore) CreatePuppy(p *Puppy) error {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Load(p.ID); ok {
		return fmt.Errorf("puppy with id %d already exists", p.ID)
	}
	s.Store(p.ID, *p)
	return nil
}

// ReadPuppy gets a puppy from the store given an ID
func (s *SyncStore) ReadPuppy(id uint16) (Puppy, error) {
	if puppyData, ok := s.Load(id); ok {
		puppy, _ := puppyData.(Puppy)
		return puppy, nil
	}
	return Puppy{}, fmt.Errorf("no puppy exists with id %d", id)
}

// UpdatePuppy puts new puppy data to the store, either creating a new one or overriding an old
func (s *SyncStore) UpdatePuppy(id uint16, puppy *Puppy) error {
	s.Store(id, *puppy)
	return nil
}

// DeletePuppy deletes a puppy from the store
func (s *SyncStore) DeletePuppy(id uint16) (bool, error) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Load(id); !ok {
		return false, fmt.Errorf("no puppy exists with id %d", id)
	}
	s.Delete(id)
	return true, nil
}
