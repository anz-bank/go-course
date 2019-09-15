package main

import (
	"fmt"
	"sync"
)

type SyncStore struct {
	sync.Map
	sync.Mutex
	currID int
}

func NewSyncStore() *SyncStore {
	return &SyncStore{}
}

// CreatePuppy takes a user provided puppy and creates a new Puppy in the store
// returns the ID of the new Puppy.
func (m *SyncStore) CreatePuppy(puppy *Puppy) (int, error) {
	m.Lock()
	defer m.Unlock()
	if puppy.Value < 0 {
		return -1, NewError(ErrInvalidValue, "Puppy value must be greater than 0")
	}
	puppy.ID = m.currID
	m.currID++
	m.Store(puppy.ID, *puppy)
	return puppy.ID, nil
}

// UpdatePuppy overrides an existing puppy in the store,
// returns an error if id is not found or does not match the puppy ID
func (m *SyncStore) UpdatePuppy(p *Puppy) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.Load(p.ID); !ok {
		return fmt.Errorf("puppy ID %d to update does not exist in the map", p.ID)
	}
	m.Store(p.ID, p)
	return nil
}

// DeletePuppy deletes an existing puppy from the store
func (m *SyncStore) DeletePuppy(id int) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.Load(id); ok {
		m.Delete(id)
		return nil
	}
	return fmt.Errorf("puppy ID %d does not exist in the map", id)
}

// ReadPuppy reads an existing puppy from the store
func (m *SyncStore) ReadPuppy(id int) (*Puppy, error) {
	if puppyData, ok := m.Load(id); ok {
		puppy, _ := puppyData.(Puppy)
		return &puppy, nil
	}
	return nil, fmt.Errorf("puppy ID %d does not exist in the map", id)
}
