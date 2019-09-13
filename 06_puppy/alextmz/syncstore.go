package main

import (
	"fmt"
	"sync"
)

type SyncStore struct {
	sync.Map
	size int
}

func NewSyncStore() *SyncStore {
	a := SyncStore{}
	return &a
}

func (m *SyncStore) CreatePuppy(p *Puppy) error {
	if p.ID == 0 {
		p.ID = m.size + 1
		m.size++
		m.Store(p.ID, *p)
		return nil
	}
	if _, ok := m.Load(p.ID); ok {
		return fmt.Errorf("puppy ID %d being created already exists", p.ID)
	}
	return fmt.Errorf("trying to create a puppy that already has an ID %d", p.ID)
}

func (m *SyncStore) ReadPuppy(id int) (Puppy, error) {
	if r, ok := m.Load(id); ok {
		puppy := r.(Puppy)
		return puppy, nil
	}
	return Puppy{}, fmt.Errorf("puppy ID %d being read does not exist", id)
}

func (m *SyncStore) UpdatePuppy(p Puppy) error {
	if _, ok := m.Load(p.ID); ok {
		m.Store(p.ID, p)
		return nil
	}
	return fmt.Errorf("puppy ID %d being updated does not exist", p.ID)
}

func (m *SyncStore) DeletePuppy(id int) error {
	if _, ok := m.Load(id); ok {
		m.Delete(id)
		m.size--
		return nil
	}
	return fmt.Errorf("puppy ID %d being deleted does not exist", id)
}
