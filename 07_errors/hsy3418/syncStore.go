package main

import (
	"sync"
)

type SyncStore struct {
	sync.Mutex
	sync.Map
}

func NewSyncStore() *SyncStore {
	return &SyncStore{}
}

// CreatePuppy adds a nuw puppy to the puppies store
func (m *SyncStore) CreatePuppy(puppy Puppy) error {
	m.Lock()
	defer m.Unlock()
	if puppy.Value < 0 {
		return ErrorEf(ErrInvalidInput, "The puppy value is invalidate")
	}
	if _, exists := m.Load(puppy.ID); exists {
		return ErrorEf(ErrDuplicate, "This puppy exists ")
	}
	m.Store(puppy.ID, puppy)
	return nil
}

// ReadPuppy retrieves the puppy for a given id from puppies store
func (m *SyncStore) ReadPuppy(id int32) (Puppy, error) {
	if p, exists := m.Load(id); exists {
		puppy, _ := p.(Puppy)
		return puppy, nil
	}
	return Puppy{}, ErrorEf(ErrNotFound, "This puppy does not exist")

}

//UpdatePuppy updates the puppy for the given id
func (m *SyncStore) UpdatePuppy(puppy Puppy) error {
	m.Lock()
	defer m.Unlock()
	if puppy.Value < 0 {
		return ErrorEf(ErrInvalidInput, "The puppy value is invalidate")
	}
	if _, exists := m.Load(puppy.ID); !exists {
		return ErrorEf(ErrNotFound, "This puppy does not exist")
	}
	m.Store(puppy.ID, puppy)
	return nil
}

//DeletePuppy delete the puppy for the given id from puppies store
func (m *SyncStore) DeletePuppy(id int32) error {
	m.Lock()
	defer m.Unlock()
	if _, exists := m.Load(id); exists {
		m.Delete(id)
		return nil
	}
	return ErrorEf(ErrNotFound, "This puppy does not exist")
}
