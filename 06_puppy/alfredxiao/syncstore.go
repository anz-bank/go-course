package main

import (
	"fmt"
	"sync"
)

type SyncStore struct {
	data sync.Map
}

func (s *SyncStore) CreatePuppy(p Puppy) error {
	_, ok := s.data.Load(p.ID)
	if ok {
		return fmt.Errorf("puppy with ID[%s] already exists", p.ID)
	}

	s.data.Store(p.ID, p)
	return nil
}

func (s *SyncStore) ReadPuppy(id string) (Puppy, error) {
	p, ok := s.data.Load(id)
	if !ok {
		return Puppy{}, fmt.Errorf("puppy with ID[%s] does not exists", id)
	}

	return p.(Puppy), nil
}

func (s *SyncStore) UpdatePuppy(p Puppy) error {
	_, ok := s.data.Load(p.ID)
	if !ok {
		return fmt.Errorf("puppy with ID[%s] does not exists", p.ID)
	}
	s.data.Store(p.ID, p)
	return nil
}

func (s *SyncStore) DeletePuppy(id string) (bool, error) {
	_, ok := s.data.Load(id)
	if !ok {
		return false, fmt.Errorf("puppy with ID[%s] does not exists", id)
	}
	s.data.Delete(id)
	return true, nil
}

func NewSyncStore() *SyncStore {
	return &SyncStore{}
}
