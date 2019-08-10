package main

import (
	"fmt"
	"sync"
)

// SynceStore  is a implementation of `Storer` backed by a sync.Map
type SyncStore struct {
	sync.Mutex
	sync.Map
	nextID int32
}

// NewSyncStore creates a new SyncStore
func NewSyncStore() *SyncStore {
	return &SyncStore{}
}

// CreatePuppy adds a nuw puppy to the puppy store and returns the id for the puppy
func (m *SyncStore) CreatePuppy(puppy Puppy) int32 {
	m.Lock()
	defer m.Unlock()
	puppy.ID = m.nextID
	m.nextID++
	m.Store(puppy.ID, puppy)
	return puppy.ID
}

// ReadPuppy retrieves the puppy for a given id from puppies store
func (m *SyncStore) ReadPuppy(id int32) (Puppy, error) {
	p, ok := m.Load(id)
	if !ok {
		return Puppy{}, fmt.Errorf("puppy with %d ID does not exist", id)
	}
	puppy, _ := p.(Puppy)
	return puppy, nil
}

//UpdatePuppy updates the puppy for the given id
func (m *SyncStore) UpdatePuppy(puppy Puppy) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.Load(puppy.ID); !ok {
		return fmt.Errorf("puppy with %d ID does not exist", puppy.ID)
	}
	m.Store(puppy.ID, puppy)
	return nil
}

//DeletePuppy delete the puppy for the given id from puppies store
func (m *SyncStore) DeletePuppy(id int32) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.Load(id); !ok {
		return fmt.Errorf("puppy with %d ID does not exist", id)
	}
	m.Delete(id)
	return nil
}
