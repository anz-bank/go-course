package main

import "sync"

type SyncStore struct {
	sync.Mutex
	sync.Map
	maxID uint
}

func newSyncStore() *SyncStore {
	return &SyncStore{}
}
func (m *SyncStore) CreatePuppy(p Puppy) uint {
	m.Lock()
	defer m.Unlock()
	m.maxID++
	p.ID = m.maxID
	m.Store(p.ID, p)
	return p.ID
}

func (m *SyncStore) ReadPuppy(id uint) Puppy {
	p, ok := m.Load(id)
	if !ok {
		return Puppy{}
	}
	puppy := p.(Puppy)
	return puppy
}
func (m *SyncStore) UpdatePuppy(id uint, p Puppy) {
	m.Store(id, p)
}

func (m *SyncStore) DeletePuppy(id uint) bool {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.Load(id); !ok {
		return false
	}
	m.Delete(id)
	return true
}
