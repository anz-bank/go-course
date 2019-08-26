package main

import (
	"fmt"
	"sync"
)

// SyncStore is sync.Map based implementation of Storer interface
type SyncStore struct {
	store  sync.Map
	nextID uint64
}

// NewSyncStore creates a new SyncStore
func NewSyncStore() *SyncStore {
	return &SyncStore{nextID: 1, store: sync.Map{}}
}

// CreatePuppy inserts given puppy in store and returns given id to the puppy
func (s *SyncStore) CreatePuppy(puppy Puppy) uint64 {
	puppy.ID = s.nextID
	s.nextID++
	s.store.Store(puppy.ID, &puppy)
	return puppy.ID
}

// ReadPuppy reads puppy with given id from the store
func (s *SyncStore) ReadPuppy(id uint64) (Puppy, error) {
	if value, exists := s.store.Load(id); exists {
		return *value.(*Puppy), nil
	}
	return Puppy{}, fmt.Errorf("no puppy with id: %d", id)
}

// UpdatePuppy update puppy with given id in store if puppy exists
// Returns nil if puppy with given id exists. Otherwise returns an error
func (s *SyncStore) UpdatePuppy(puppy Puppy) error {
	if _, exists := s.store.Load(puppy.ID); exists {
		s.store.Store(puppy.ID, &puppy)
		return nil
	}
	return fmt.Errorf("no puppy with id: %d", puppy.ID)
}

// DeletePuppy deletes puppy with given id from store
// Returns nil if puppy with given id exists. Otherwise returns an error
func (s *SyncStore) DeletePuppy(id uint64) error {
	if _, exists := s.store.Load(id); exists {
		s.store.Delete(id)
		return nil
	}
	return fmt.Errorf("no puppy with id: %d", id)
}
