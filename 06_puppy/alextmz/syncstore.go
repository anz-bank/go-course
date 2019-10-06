package main

import (
	"fmt"
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
		return fmt.Errorf("puppy pointer is nil")
	}
	if p.ID != 0 {
		return fmt.Errorf("trying to create a puppy already initialized with ID %d", p.ID)
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
		return Puppy{}, fmt.Errorf("puppy ID %d being read does not exist", id)
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
		return fmt.Errorf("puppy ID %d being updated does not exist", p.ID)
	}
	m.Store(p.ID, p)
	return nil
}

func (m *SyncStore) DeletePuppy(id int) error {
	m.Lock()
	defer m.Unlock()
	_, ok := m.Load(id)
	if !ok {
		return fmt.Errorf("puppy ID %d being deleted does not exist", id)
	}
	m.Delete(id)
	return nil
}
