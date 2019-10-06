package main

import (
	"strconv"
	"sync"
)

// SyncStore stores Puppy details with Puppy Id as Key and Puppy  as value
type SyncStore struct {
	sync.Map
	m sync.Mutex
}

// CreatePuppy creates a Puppy in syncstore
func (s *SyncStore) CreatePuppy(p Puppy) error {
	s.m.Lock()
	defer s.m.Unlock()
	if _, ok := s.Load(p.ID); ok {
		return ErrorF(Duplicate, "puppy with Id %d already exists", p.ID)
	}
	val, _ := strconv.Atoi(p.Value)
	if val < 0 {
		return ErrorF(Invalid, "puppy with value less than 0 not allowed")
	}
	s.Store(p.ID, p)
	return nil
}

// ReadPuppy reads a Puppy from syncstore
func (s *SyncStore) ReadPuppy(id uint32) (Puppy, error) {
	s.m.Lock()
	defer s.m.Unlock()
	if puppy, ok := s.Load(id); ok {
		return puppy.(Puppy), nil
	}
	return Puppy{}, ErrorF(NotFound, "puppy with Id %d does not exists", id)
}

// UpdatePuppy updates a Puppy in syncstore
func (s *SyncStore) UpdatePuppy(p Puppy) error {
	s.m.Lock()
	defer s.m.Unlock()
	if _, ok := s.Load(p.ID); !ok {
		return ErrorF(NotFound, "puppy with Id %d does not exists", p.ID)
	}
	s.Store(p.ID, p)
	return nil
}

// DeletePuppy deletes a Puppy from syncstore
func (s *SyncStore) DeletePuppy(id uint32) error {
	s.m.Lock()
	defer s.m.Unlock()
	if _, ok := s.Load(id); !ok {
		return ErrorF(NotFound, "puppy with Id %d does not exists", id)
	}
	s.Delete(id)
	return nil
}
