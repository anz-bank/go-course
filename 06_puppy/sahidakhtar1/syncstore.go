package main

import "sync"

type SyncStore struct {
	sync.Mutex
	sync.Map
}

func newSyncStore() *SyncStore {
	return &SyncStore{}
}
func (m *SyncStore) CreatePuppy(p Puppy) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.Load(p.ID); !ok {
		m.Store(p.ID, p)
	}
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
