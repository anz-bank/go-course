package main

import (
	"strconv"
)

// CreatePuppy creates a Puppy in memstore
func (s *SyncStore) CreatePuppy(p Puppy) error {
	s.m.Lock()
	defer s.m.Unlock()
	if _, ok := s.Load(p.ID); !ok {
		val, _ := strconv.Atoi(p.Value)
		if val < 0 {
			return ErrorF(InvalidValue, "puppy with value less than 0 not allowed")
		}
		s.Store(p.ID, p)
		return nil
	}
	return ErrorF(Duplicate, "puppy with Id %d already exists", p.ID)
}

// ReadPuppy reads a Puppy from memstore
func (s *SyncStore) ReadPuppy(id uint32) (Puppy, error) {
	s.m.Lock()
	defer s.m.Unlock()
	if puppy, ok := s.Load(id); ok {
		return puppy.(Puppy), nil
	}
	return Puppy{}, ErrorF(NotFound, "puppy with Id %d does not exists", id)
}

// UpdatePuppy updates a Puppy
func (s *SyncStore) UpdatePuppy(p Puppy) error {
	s.m.Lock()
	defer s.m.Unlock()
	if _, ok := s.Load(p.ID); !ok {
		return ErrorF(NotFound, "puppy with Id %d does not exists", p.ID)
	}
	s.Store(p.ID, p)
	return nil
}

// DeletePuppy deletes a Puppy
func (s *SyncStore) DeletePuppy(id uint32) bool {
	s.m.Lock()
	defer s.m.Unlock()
	if _, ok := s.Load(id); !ok {
		return false
	}
	s.Delete(id)
	return true
}
