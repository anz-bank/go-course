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
    return fmt.Errorf("Puppy with ID[%s] already exists", p.ID)
  }

  s.data.Store(p.ID, p)
  return nil
}

func (s *SyncStore) ReadPuppy(ID string) (Puppy, error) {
  p, ok := s.data.Load(ID)
  if !ok {
    return Puppy{}, fmt.Errorf("Puppy with ID[%s] does not exists", p.(Puppy).ID)
  }

  return p.(Puppy), nil
}

func (s *SyncStore) UpdatePuppy(p Puppy) error {
  _, ok := s.data.Load(p.ID)
  if !ok {
    return fmt.Errorf("Puppy with ID[%s] does not exists", p.ID)
  }
  s.data.Store(p.ID, p)
  return nil
}

func (s *SyncStore) DeletePet(ID string) (bool, error) {
  p, ok := s.data.Load(ID)
  if !ok {
    return false, fmt.Errorf("Puppy with ID[%s] does not exists", p.(Puppy).ID)
  }
  s.data.Delete(ID)
  return true, nil
}

func NewSyncStore() *SyncStore {
  return &SyncStore{}
}
