package store

import (
	"sync"

	"github.com/anz-bank/go-course/11_notify/kasunfdo/pkg/puppy"
)

// SyncStore is sync.Map based implementation of Storer interface
type SyncStore struct {
	store sync.Map
	sync.Mutex
	nextID uint64
}

// NewSyncStore creates a new SyncStore
func NewSyncStore() *SyncStore {
	return &SyncStore{nextID: 1, store: sync.Map{}}
}

// CreatePuppy inserts given puppy in store and returns given id to the puppy
// Will return an error if value of puppy is negative
func (s *SyncStore) CreatePuppy(p puppy.Puppy) (uint64, error) {
	if err := p.Validate(); err != nil {
		return 0, err
	}

	s.Lock()
	defer s.Unlock()

	p.ID = s.nextID
	s.nextID++
	s.store.Store(p.ID, &p)
	return p.ID, nil
}

// ReadPuppy reads puppy with given id from the store
// Will return an error if puppy with given id does not exist
func (s *SyncStore) ReadPuppy(id uint64) (puppy.Puppy, error) {
	value, ok := s.store.Load(id)
	if !ok {
		return puppy.Puppy{}, puppy.ErrorEf(puppy.ErrNotFound, nil, "puppy with id: %v is not found", id)
	}

	return *value.(*puppy.Puppy), nil
}

// UpdatePuppy update puppy with given id in store if puppy exists
// Returns nil if puppy with given id exists. Otherwise returns an error
func (s *SyncStore) UpdatePuppy(p puppy.Puppy) error {
	if err := p.Validate(); err != nil {
		return err
	}

	s.Lock()
	defer s.Unlock()

	if _, ok := s.store.Load(p.ID); !ok {
		return puppy.ErrorEf(puppy.ErrNotFound, nil, "puppy with id: %v is not found", p.ID)
	}

	s.store.Store(p.ID, &p)
	return nil
}

// DeletePuppy deletes puppy with given id from store
// Returns nil if puppy with given id exists. Otherwise returns an error
func (s *SyncStore) DeletePuppy(id uint64) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.store.Load(id); !ok {
		return puppy.ErrorEf(puppy.ErrNotFound, nil, "puppy with id: %v is not found", id)
	}

	s.store.Delete(id)
	return nil
}
