package main

import (
	"fmt"
)

// CreatePuppy creates a Puppy in syncstore
func (s *SyncStore) CreatePuppy(p Puppy) error {
	if _, ok := s.Load(p.ID); !ok {
		s.Store(p.ID, p)
		return nil
	}
	return fmt.Errorf("puppy with Id %d already exists", p.ID)
}

// ReadPuppy reads a Puppy from syncstore
func (s *SyncStore) ReadPuppy(id uint32) (Puppy, error) {
	if puppy, ok := s.Load(id); ok {
		return puppy.(Puppy), nil
	}
	return Puppy{}, fmt.Errorf("puppy with Id %d does not exists", id)
}

// UpdatePuppy updates a Puppy in syncstore
func (s *SyncStore) UpdatePuppy(p Puppy) error {
	if _, ok := s.Load(p.ID); !ok {
		return fmt.Errorf("puppy with Id %d does not exists", p.ID)
	}
	s.Store(p.ID, p)
	return nil
}

// DeletePuppy deletes a Puppy from syncstore
func (s *SyncStore) DeletePuppy(id uint32) error {
	if _, ok := s.Load(id); !ok {
		return fmt.Errorf("puppy with Id %d does not exists", id)
	}
	s.Delete(id)
	return nil
}
