package puppy

import (
	"strconv"
	"sync"
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
func (s *SyncStore) CreatePuppy(p *Puppy) (uint32, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	i, err := strconv.Atoi(p.Value)
	if err != nil {
		return 0, Errorf(ErrInvalidInput, "Puppy value can't be negative")
	}
	if i < 0 {
		return 0, Errorf(ErrInvalidInput, "Unrecognised input")
	}
	s.nextID++
	p.ID = s.nextID
	s.syncStore.Store(p.ID, *p)
	return p.ID, nil
}

// ReadPuppy read a puppy given its id.
// It returns the pointer to that puppy.
func (s *SyncStore) ReadPuppy(id uint32) (*Puppy, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if p, ok := s.syncStore.Load(id); ok {
		returnPuppy := p.(Puppy)
		return &returnPuppy, nil
	}
	return nil, Errorf(ErrNotFound, "Puppy ID can't be found")
}

// UpdatePuppy updates the store with key of id with the new puppy.
func (s *SyncStore) UpdatePuppy(id uint32, p *Puppy) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	i, err := strconv.Atoi(p.Value)
	if err != nil {
		return Errorf(ErrInvalidInput, "Puppy value is not recognised")
	} else if i < 0 {
		return Errorf(ErrInvalidInput, "Puppy value can't be negative")
	}
	if _, ok := s.syncStore.Load(id); ok {
		p.ID = id
		s.syncStore.Store(id, p)
		return nil
	}
	return Errorf(ErrNotFound, "Puppy ID can't be found, update operation failed")
}

// DeletePuppy delete the puppy given the id.
func (s *SyncStore) DeletePuppy(id uint32) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.syncStore.Load(id); ok {
		s.syncStore.Delete(id)
		return nil
	}
	return Errorf(ErrNotFound, "Puppy ID can't be found, delete operation failed")
}
