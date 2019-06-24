package store

import (
	"fmt"
	"sync"

	puppy "github.com/anz-bank/go-course/08_project/n0npax/pkg/puppy"
)

//SyncStore sync.Map based type for storing puppies data
type SyncStore struct {
	sync.Map
	sync.Mutex
	nextID int
}

// NewSyncStore creates new storer for SyncMap
func NewSyncStore() *SyncStore {
	return &SyncStore{}
}

// CreatePuppy creates puppy
func (s *SyncStore) CreatePuppy(p *puppy.Puppy) (int, error) {
	if p.Value < 0 {
		return -1, puppy.Errorf(puppy.ErrInvalidInputCode, "Puppy value have to be positive number")
	}
	s.Lock()
	defer s.Unlock()
	p.ID = s.nextID
	s.nextID++
	s.Store(p.ID, *p)
	return p.ID, nil
}

// ReadPuppy reads puppy from backend
func (s *SyncStore) ReadPuppy(id int) (*puppy.Puppy, error) {
	if puppyData, ok := s.Load(id); ok {
		p := puppyData.(puppy.Puppy)
		return &p, nil
	}
	return nil, puppy.Errorf(puppy.ErrNotFoundCode, fmt.Sprintf("Puppy with ID (%v) not found", id))
}

// UpdatePuppy updates puppy
func (s *SyncStore) UpdatePuppy(id int, p *puppy.Puppy) error {
	if p.Value < 0 {
		return puppy.Errorf(puppy.ErrInvalidInputCode, "Puppy value have to be positive number")
	}
	if id != p.ID {
		return puppy.Errorf(puppy.ErrInvalidInputCode, "ID is corrupted. Please ensure object ID matched provided ID")
	}
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Load(id); !ok {
		return puppy.Errorf(puppy.ErrNotFoundCode, fmt.Sprintf("Puppy with ID (%v) not found", id))
	}
	s.Store(id, *p)
	return nil
}

// DeletePuppy deletes puppy
func (s *SyncStore) DeletePuppy(id int) (bool, error) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Load(id); !ok {
		return false, puppy.Errorf(puppy.ErrNotFoundCode, fmt.Sprintf("Puppy with ID (%v) not found", id))
	}
	s.Delete(id)
	return true, nil
}
