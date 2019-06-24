package main

import (
	"fmt"
	"sync"
)

type syncStore struct {
	data sync.Map
}

func (s *syncStore) CreatePuppy(p Puppy) error {
	_, ok := s.data.Load(p.ID)
	if ok {
		return fmt.Errorf("puppy with ID[%s] already exists", p.ID)
	}

	s.data.Store(p.ID, p)
	return nil
}

func (s *syncStore) ReadPuppy(id string) (Puppy, error) {
	p, ok := s.data.Load(id)
	if !ok {
		return Puppy{}, fmt.Errorf("puppy with ID[%s] does not exists", id)
	}

	return p.(Puppy), nil
}

func (s *syncStore) UpdatePuppy(p Puppy) error {
	_, ok := s.data.Load(p.ID)
	if !ok {
		return fmt.Errorf("puppy with ID[%s] does not exists", p.ID)
	}
	s.data.Store(p.ID, p)
	return nil
}

func (s *syncStore) DeletePuppy(id string) (bool, error) {
	_, ok := s.data.Load(id)
	if !ok {
		return false, fmt.Errorf("puppy with ID[%s] does not exists", id)
	}
	s.data.Delete(id)
	return true, nil
}

func NewSyncStore() Storer {
	return &syncStore{}
}
