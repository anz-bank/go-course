package store

import (
	"fmt"

	puppy "github.com/anz-bank/go-course/08_project/n0npax/pkg/puppy"
)

// NewSyncStore creates new storer for SyncMap
func NewSyncStore() *SyncStore {
	return &SyncStore{}
}

// CreatePuppy creates puppy
func (s *SyncStore) CreatePuppy(p *puppy.Puppy) (int, error) {
	if p.Value < 0 {
		return -1, puppy.ErrInvalidInput(puppy.InvalidInputMsg)
	}
	s.Lock()
	defer s.Unlock()
	p.ID = s.total
	s.total++
	s.Store(p.ID, *p)
	return p.ID, nil
}

// ReadPuppy reads puppy from backend
func (s *SyncStore) ReadPuppy(id int) (*puppy.Puppy, error) {
	if puppyData, ok := s.Load(id); ok {
		p := puppyData.(puppy.Puppy)
		return &p, nil
	}
	return nil, puppy.ErrNotFound(fmt.Sprintf(puppy.PuppyNotFoundMsg, id))
}

// UpdatePuppy updates puppy
func (s *SyncStore) UpdatePuppy(id int, p *puppy.Puppy) error {
	if p.Value < 0 {
		return puppy.ErrInvalidInput(puppy.InvalidInputMsg)
	}
	if id != p.ID {
		return puppy.ErrInvalidInput(puppy.CorruptedIDMsg)
	}
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Load(id); !ok {
		return puppy.ErrNotFound(fmt.Sprintf(puppy.PuppyNotFoundMsg, id))
	}
	s.Store(id, *p)
	return nil
}

// DeletePuppy deletes puppy
func (s *SyncStore) DeletePuppy(id int) (bool, error) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Load(id); !ok {
		return false, puppy.ErrNotFound(fmt.Sprintf(puppy.PuppyNotFoundMsg, id))
	}
	s.Delete(id)
	return true, nil
}
