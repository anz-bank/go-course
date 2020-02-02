package main

import (
	"sync"
)

type SyncStore struct {
	nextID int
	sync.Map
	sync.Mutex
}

func NewSyncStore() *SyncStore {
	a := SyncStore{}
	return &a
}

func (m *SyncStore) CreatePuppy(p *Puppy) error {
	if p == nil {
		return NewError(ErrInvalidRequest, "puppy pointer is nil")
	}
	if p.Value < 0 {
		return NewErrorf(ErrNegativeValue, "puppy value (%f) is < 0", p.Value)
	}
	if p.ID != 0 {
		return NewErrorf(ErrInvalidRequest, "trying to create a puppy already initialized with ID %d", p.ID)
	}
	m.Lock()
	defer m.Unlock()
	p.ID = m.nextID + 1
	m.nextID++
	m.Store(p.ID, *p)
	return nil
}

func (m *SyncStore) ReadPuppy(id int) (Puppy, error) {
	v, ok := m.Load(id)
	if !ok {
		return Puppy{}, NewErrorf(ErrNotFound, "puppy ID %d being read does not exist", id)
	}
	m.Lock()
	defer m.Unlock()
	return v.(Puppy), nil
}

func (m *SyncStore) UpdatePuppy(p Puppy) error {
	m.Lock()
	defer m.Unlock()
	_, ok := m.Load(p.ID)
	if !ok {
		return NewErrorf(ErrNotFound, "puppy ID %d being updated does not exist", p.ID)
	}
	if p.Value < 0 {
		return NewErrorf(ErrNegativeValue, "puppy value (%f) is < 0", p.Value)
	}
	m.Store(p.ID, p)
	return nil
}

func (m *SyncStore) DeletePuppy(id int) error {
	m.Lock()
	defer m.Unlock()
	_, ok := m.Load(id)
	if !ok {
		return NewErrorf(ErrNotFound, "puppy ID %d being deleted does not exist", id)
	}
	m.Delete(id)
	return nil
}
