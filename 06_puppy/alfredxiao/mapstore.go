package main

import (
	"fmt"

	"github.com/google/uuid"
)

type mapStore struct {
	data map[string]Puppy
}

func (s *mapStore) CreatePuppy(p Puppy) string {
	id := uuid.New().String()

	p.ID = id
	s.data[id] = p
	return id
}

func (s *mapStore) ReadPuppy(id string) (Puppy, error) {
	p, ok := s.data[id]
	if !ok {
		return Puppy{}, fmt.Errorf("puppy with ID[%s] does not exists", p.ID)
	}

	return p, nil
}

func (s *mapStore) UpdatePuppy(p Puppy) error {
	_, ok := s.data[p.ID]
	if !ok {
		return fmt.Errorf("puppy with ID[%s] does not exists", p.ID)
	}
	s.data[p.ID] = p
	return nil
}

func (s *mapStore) DeletePuppy(id string) error {
	p, ok := s.data[id]
	if !ok {
		return fmt.Errorf("puppy with ID[%s] does not exists", p.ID)
	}
	delete(s.data, id)
	return nil
}

func NewMapStore() Storer {
	return &mapStore{
		data: make(map[string]Puppy),
	}
}
