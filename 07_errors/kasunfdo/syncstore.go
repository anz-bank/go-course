package main

import (
	"fmt"
	"sync"
)

// SyncStore is sync.Map based implementation of Storer interface
type SyncStore struct {
	store sync.Map
	sync.Mutex
	nextID uint64
}

// NewSyncStore creates a new SyncStore
func NewSyncStore() *SyncStore {
	return &SyncStore{nextID: 1, store: sync.Map{}}
}

// CreatePuppy inserts given puppy in store and returns given id to the puppy
// Will return an error if value of puppy is negative
func (s *SyncStore) CreatePuppy(puppy Puppy) (uint64, error) {
	if puppy.Value < 0 {
		return 0, NewError(ErrInvalid, "value of puppy is negative")
	}

	s.Lock()
	defer s.Unlock()

	puppy.ID = s.nextID
	s.nextID++
	s.store.Store(puppy.ID, &puppy)
	return puppy.ID, nil
}

// ReadPuppy reads puppy with given id from the store
// Will return an error if puppy with given id does not exist
func (s *SyncStore) ReadPuppy(id uint64) (Puppy, error) {
	value, exists := s.store.Load(id)

	if !exists {
		return Puppy{}, NewError(ErrNotFound, fmt.Sprintf("puppy with id: %v is not found", id))
	}
	return *value.(*Puppy), nil
}

// UpdatePuppy update puppy with given id in store if puppy exists
// Returns nil if puppy with given id exists. Otherwise returns an error
func (s *SyncStore) UpdatePuppy(puppy Puppy) error {
	if puppy.Value < 0 {
		return NewError(ErrInvalid, "value of puppy is negative")
	}

	s.Lock()
	defer s.Unlock()

	if _, exists := s.store.Load(puppy.ID); !exists {
		return NewError(ErrNotFound, fmt.Sprintf("puppy with id: %v is not found", puppy.ID))
	}

	s.store.Store(puppy.ID, &puppy)
	return nil
}

// DeletePuppy deletes puppy with given id from store
// Returns nil if puppy with given id exists. Otherwise returns an error
func (s *SyncStore) DeletePuppy(id uint64) error {
	s.Lock()
	defer s.Unlock()

	if _, exists := s.store.Load(id); !exists {
		return NewError(ErrNotFound, fmt.Sprintf("puppy with id: %v is not found", id))
	}

	s.store.Delete(id)
	return nil
}
