package main

import (
	"fmt"
)

func NewSyncStore() *SyncStore {
	return &SyncStore{}
}

func (m *SyncStore) CreatePuppy(puppy *Puppy) int {
	m.Lock()
	defer m.Unlock()
	puppy.ID = m.total
	m.total++
	m.Store(puppy.ID, *puppy)
	return puppy.ID
}

func (m *SyncStore) ReadPuppy(id int) (*Puppy, error) {
	if puppyData, ok := m.Load(id); ok {
		puppy := puppyData.(Puppy)
		return &puppy, nil
	}
	return nil, fmt.Errorf("puppy with ID: %d does not exist", id)

}

func (m *SyncStore) UpdatePuppy(id int, puppy *Puppy) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.Load(id); !ok {
		return fmt.Errorf("puppy with ID: %d does not exist", id)
	}
	if id != puppy.ID {
		return fmt.Errorf("puppy ID corrupted")
	}
	m.Store(id, *puppy)
	return nil
}

func (m *SyncStore) DeletePuppy(id int) (bool, error) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.Load(id); !ok {
		return false, fmt.Errorf("puppy with ID: %d does not exist", id)
	}
	m.Delete(id)
	return true, nil
}
