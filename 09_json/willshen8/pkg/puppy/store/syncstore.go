package puppy

import (
	"strconv"
	"sync"

	puppy "github.com/anz-bank/go-course/09_json/willshen8/pkg/puppy"
)

type SyncStore struct {
	syncStore sync.Map
	nextID    uint32
	mu        sync.Mutex
}

// NewMapStore initialise a new SyncStore
func NewSyncStore() *SyncStore {
	return &SyncStore{syncStore: sync.Map{}}
}

// CreatePuppy create a new puppy and store in mapStore.
func (s *SyncStore) CreatePuppy(p *puppy.Puppy) (uint32, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if i, err := strconv.Atoi(p.Value); err == nil {
		if i < 0 {
			return 0, &puppy.Error{
				Message: "Puppy value can't be less than 0.",
				Code:    puppy.NegativeValue,
			}
		}
	} else {
		return 0, &puppy.Error{
			Message: "Unrecongised puppy value.",
			Code:    puppy.ErrorValueFormat,
		}
	}
	s.nextID++
	p.ID = s.nextID
	s.syncStore.Store(p.ID, *p)

	return p.ID, nil
}

// ReadPuppy read a puppy given its id.
// It returns the pointer to that puppy.
func (s *SyncStore) ReadPuppy(id uint32) (*puppy.Puppy, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if p, ok := s.syncStore.Load(id); ok {
		returnPuppy := p.(puppy.Puppy)
		return &returnPuppy, nil
	}
	return nil, &puppy.Error{
		Message: "Puppy ID can not be found, read operation failed.",
		Code:    puppy.NonExistentPuppy,
	}
}

// UpdatePuppy updates the store with key of id with the new puppy.
// It returns a boolean whether the operation is successful or not.
func (s *SyncStore) UpdatePuppy(id uint32, p *puppy.Puppy) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.syncStore.Load(id); ok {
		p.ID = id
		s.syncStore.Store(id, p)
		return true, nil
	}
	return false, &puppy.Error{
		Message: "Puppy ID can not be found, update operation failed.",
		Code:    puppy.NonExistentPuppy,
	}
}

// DeletePuppy delete the puppy given the id.
// It returns true/success or false/unsuccessful
func (s *SyncStore) DeletePuppy(id uint32) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.syncStore.Load(id); ok {
		s.syncStore.Delete(id)
		return true, nil
	}
	return false, &puppy.Error{
		Message: "Puppy ID can not be found, delete operation failed.",
		Code:    puppy.NonExistentPuppy,
	}
}
