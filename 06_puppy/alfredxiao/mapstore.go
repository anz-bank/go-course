package main

import (
	"fmt"
)

type MapStore struct {
	data map[string]Puppy
}

func (s *MapStore) CreatePuppy(p Puppy) error {
	_, ok := s.data[p.ID]
	if ok {
		return fmt.Errorf("puppy with ID[%s] already exists", p.ID)
	}

	s.data[p.ID] = p
	return nil
}

func (s *MapStore) ReadPuppy(id string) (Puppy, error) {
	p, ok := s.data[id]
	if !ok {
		return Puppy{}, fmt.Errorf("puppy with ID[%s] does not exists", p.ID)
	}

	return p, nil
}

func (s *MapStore) UpdatePuppy(p Puppy) error {
	_, ok := s.data[p.ID]
	if !ok {
		return fmt.Errorf("puppy with ID[%s] does not exists", p.ID)
	}
	s.data[p.ID] = p
	return nil
}

func (s *MapStore) DeletePuppy(id string) (bool, error) {
	p, ok := s.data[id]
	if !ok {
		return false, fmt.Errorf("puppy with ID[%s] does not exists", p.ID)
	}
	delete(s.data, id)
	return true, nil
}

func NewMapStore() *MapStore {
	return &MapStore{
		data: make(map[string]Puppy),
	}
}
