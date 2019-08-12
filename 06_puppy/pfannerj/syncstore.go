package main

import (
	"fmt"
	"sync"
)

// SyncStore is a sync.Map based in-memory implementation of PuppyStorer.
type SyncStore struct {
	sync.Map
	currID uint32
}

// NewSyncStore creates a new in-memory store with sync map intialised.
func NewSyncStore() *SyncStore {
	return &SyncStore{}
}

// CreatePuppy adds a new puppy to the sync store.
func (s *SyncStore) CreatePuppy(puppy Puppy) (uint32, error) {
	s.currID++
	puppy.ID = s.currID
	s.Store(puppy.ID, puppy)
	return puppy.ID, nil
}

// ReadPuppy gets a puppy from the sync store with the given ID.
func (s *SyncStore) ReadPuppy(puppyID uint32) (Puppy, error) {
	if puppy, ok := s.Load(puppyID); ok {
		puppyOut, _ := puppy.(Puppy)
		return puppyOut, nil
	}
	return Puppy{}, fmt.Errorf("no puppy found with id %d", puppyID)
}

// UpdatePuppy modifies puppy data in the sync store, either creating a new one or overwriting an old one.
func (s *SyncStore) UpdatePuppy(puppyID uint32, p Puppy) (uint32, error) {
	if _, ok := s.Load(puppyID); !ok {
		s.currID++
		puppyID = s.currID
	}
	p.ID = puppyID //ensure the ID within p always matches the sync store key (puppyID)
	s.Store(puppyID, p)
	return puppyID, nil
}

// DeletePuppy deletes the puppy with the given ID from the sync store.
func (s *SyncStore) DeletePuppy(puppyID uint32) error {
	if _, ok := s.Load(puppyID); ok {
		s.Delete(puppyID)
	}
	return nil
}
