package main

import "sync"

type SyncStore struct {
	syncStore sync.Map
	counter   uint32
}

// NewMapStore initialise a new SyncStore
func NewSyncStore() *SyncStore {
	var newSyncStore = SyncStore{}
	newSyncStore.syncStore = sync.Map{}
	return &newSyncStore
}

// IncrementCounter increase the ID counter everytime a new Puppy is created.
func (s *SyncStore) IncrementCounter() {
	s.counter++
}

// CreatePuppy create a new puppy and store in mapStore.
func (s *SyncStore) CreatePuppy(p *Puppy) uint32 {
	s.IncrementCounter()
	p.ID = s.counter
	s.syncStore.Store(p.ID, *p)
	return p.ID
}

// ReadPuppy read a puppy given its id.
// It returns the pointer to that puppy.
func (s *SyncStore) ReadPuppy(id uint32) *Puppy {
	if puppy, ok := s.syncStore.Load(id); ok {
		returnPuppy := puppy.(Puppy)
		return &returnPuppy
	}
	return nil
}

// UpdatePuppy updates the store with key of id with the new puppy.
// It returns a boolean whether the operation is successful or not.
func (s *SyncStore) UpdatePuppy(id uint32, puppy *Puppy) bool {
	if _, ok := s.syncStore.Load(id); !ok {
		return false
	}
	puppy.ID = id
	s.syncStore.Store(id, puppy)
	return true
}

// DeletePuppy delete the puppy given the id.
// It returns true/success or false/unsuccessful.
func (s *SyncStore) DeletePuppy(id uint32) bool {
	if _, ok := s.syncStore.Load(id); ok {
		s.syncStore.Delete(id)
		return true
	}
	return false
}
