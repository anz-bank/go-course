package main

import (
	"errors"
)

func NewSyncStore() *SyncStore {
	return &SyncStore{}
}

func (m *SyncStore) CreatePuppy(puppy *Puppy) int {
	var newID int
	m.Range(func(_, _ interface{}) bool {
		newID++
		return true
	})
	puppy.ID = newID
	m.Store(puppy.ID, *puppy)
	return puppy.ID
}

func (m *SyncStore) ReadPuppy(id int) (*Puppy, error) {
	puppyData, ok := m.Load(id)
	if !ok {
		return nil, errors.New("doesn't exists")
	}
	puppy := puppyData.(Puppy)
	return &puppy, nil
}

func (m *SyncStore) UpdatePuppy(id int, puppy *Puppy) error {
	m.Store(id, *puppy)
	return nil
}

func (m *SyncStore) DeletePuppy(id int) (bool, error) {
	if _, ok := m.Load(id); !ok {
		return false, errors.New("doesn't exist")
	}
	m.Delete(id)
	return true, nil
}
