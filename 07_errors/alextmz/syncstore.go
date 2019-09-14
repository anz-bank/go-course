package main

import (
	"sync"
)

type SyncStore struct {
	sync.Map
	sync.Mutex
}

func NewSyncStore() *SyncStore {
	a := SyncStore{}
	return &a
}

func (m *SyncStore) CreatePuppy(p *Puppy) error {
	if _, ok := m.Load(p.ID); !ok {
		if p.Value < 0 {
			return NewError(ErrValueLessThanZero)
		}
		m.Lock()
		m.Store(p.ID, *p)
		m.Unlock()
		return nil
	}
	return NewError(ErrIDBeingCreatedAlreadyExists)
}

func (m *SyncStore) ReadPuppy(id PuppyID) (*Puppy, error) {
	r, ok := m.Load(id)
	if ok {
		puppy := r.(Puppy)
		return &puppy, nil
	}
	return nil, NewError(ErrIDBeingReadDoesNotExist)
}

func (m *SyncStore) UpdatePuppy(id PuppyID, p *Puppy) error {
	if _, ok := m.Load(id); ok {
		if p.Value < 0 {
			return NewError(ErrValueLessThanZero)
		}
		m.Lock()
		m.Store(id, *p)
		m.Unlock()
		return nil
	}
	return NewError(ErrIDBeingUpdatedDoesNotExist)
}

func (m *SyncStore) DeletePuppy(id PuppyID) error {
	_, ok := m.Load(id)
	if ok {
		m.Lock()
		m.Delete(id)
		m.Unlock()
		return nil
	}
	return NewError(ErrIDBeingDeletedDoesNotExist)
}
