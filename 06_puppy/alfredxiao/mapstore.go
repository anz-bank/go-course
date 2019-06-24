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
    return fmt.Errorf("Puppy with ID[%s] already exists", p.ID)
  }

  s.data[p.ID] = p
  return nil
}

func (s *MapStore) ReadPuppy(ID string) (Puppy, error) {
  p, ok := s.data[ID]
  if !ok {
    return Puppy{}, fmt.Errorf("Puppy with ID[%s] does not exists", p.ID)
  }

  return p, nil
}

func (s *MapStore) UpdatePuppy(p Puppy) error {
  _, ok := s.data[p.ID]
  if !ok {
    return fmt.Errorf("Puppy with ID[%s] does not exists", p.ID)
  }
  s.data[p.ID] = p
  return nil
}

func (s *MapStore) DeletePuppy(ID string) (bool, error) {
  p, ok := s.data[ID]
  if !ok {
    return false, fmt.Errorf("Puppy with ID[%s] does not exists", p.ID)
  }
  delete(s.data, ID)
  return true, nil
}

func NewMapStore() *MapStore {
  return &MapStore{
    data: make(map[string]Puppy),
  }
}
