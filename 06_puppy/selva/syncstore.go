package main

import (
	"sync"
)

// SyncStore struct
type SyncStore struct {
	sync.Mutex
	sync.Map
	maxID int
}

// NewSyncStore constructor
func NewSyncStore() *SyncStore {
	return &SyncStore{}
}

//CreatePuppy to create
func (m *SyncStore) CreatePuppy(pup Puppy) int {
	pup.ID = m.maxID
	m.Store(m.maxID, pup)
	m.maxID++
	return pup.ID
}

//ReadPuppy to read
func (m *SyncStore) ReadPuppy(pupID int) Puppy {
	pupData, _ := m.Load(pupID)
	pup, _ := pupData.(Puppy)
	return pup
}

//UpdatePuppy to read
func (m *SyncStore) UpdatePuppy(pupID int, pup Puppy) {
	m.Store(pupID, pup)
	//return pup
}

//DeletePuppy to delete
func (m *SyncStore) DeletePuppy(pupID int) bool {
	m.Delete(pupID)
	return true
}
