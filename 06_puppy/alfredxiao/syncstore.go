package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type syncStore struct {
	data sync.Map
}

func (s *syncStore) CreatePuppy(p Puppy) string {
	id := uuid.New().String()

	p.ID = id
	s.data.Store(id, p)
	return id
}

func (s *syncStore) ReadPuppy(id string) (Puppy, error) {
	p, ok := s.data.Load(id)
	if !ok {
		return Puppy{}, fmt.Errorf("puppy with ID[%s] does not exists", id)
	}

	return p.(Puppy), nil
}

func (s *syncStore) UpdatePuppy(id string, p Puppy) error {
	if id != p.ID {
		return fmt.Errorf("bad update request, two IDs (%s, %s) do not match",
			id, p.ID)
	}

	_, ok := s.data.Load(id)
	if !ok {
		return fmt.Errorf("puppy with ID[%s] does not exists", p.ID)
	}

	s.data.Store(id, p)
	return nil
}

func (s *syncStore) DeletePuppy(id string) error {
	_, ok := s.data.Load(id)
	if !ok {
		return fmt.Errorf("puppy with ID[%s] does not exists", id)
	}
	s.data.Delete(id)
	return nil
}

func NewSyncStore() Storer {
	return &syncStore{}
}
