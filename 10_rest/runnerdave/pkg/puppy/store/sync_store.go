package store

import (
	"sort"
	"sync"

	puppy "github.com/anz-bank/go-course/10_rest/runnerdave/pkg/puppy"
)

// SyncStore implementation of the Storer interface
type SyncStore struct {
	sync.Mutex
	sync.Map
	nextID int16
}

// NewSyncStore creates a new in-memory store with map intialised
func NewSyncStore() *SyncStore {
	return &SyncStore{nextID: 1}
}

// CreatePuppy saves new puppy if not in store, if it is already returns error
func (s *SyncStore) CreatePuppy(p puppy.Puppy) error {
	if err := puppy.ValidateValue(p.Value); err != nil {
		return err
	}
	s.Lock()
	defer s.Unlock()
	p.ID = s.nextID
	s.Store(p.ID, p)
	s.nextID++
	return nil
}

// ReadPuppy gets a puppy from the store given an ID
func (s *SyncStore) ReadPuppy(id int16) (puppy.Puppy, error) {
	if puppyData, ok := s.Load(id); ok {
		puppy, _ := puppyData.(puppy.Puppy)
		return puppy, nil
	}
	return puppy.Puppy{}, puppy.Errorf(puppy.ErrIDNotFound, "puppy with ID:%d not found", id)
}

// UpdatePuppy puts new puppy data to the store, error if id does not exist
func (s *SyncStore) UpdatePuppy(id int16, p *puppy.Puppy) error {
	if err := puppy.ValidateValue(p.Value); err != nil {
		return err
	}
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Load(id); !ok {
		return puppy.Errorf(puppy.ErrIDNotFound, "puppy with ID:%d not found", id)
	}
	p.ID = id
	s.Store(id, *p)
	return nil
}

// DeletePuppy deletes a puppy from the store
func (s *SyncStore) DeletePuppy(id int16) error {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Load(id); !ok {
		return puppy.Errorf(puppy.ErrIDNotFound, "puppy with ID:%d not found", id)
	}
	s.Delete(id)
	return nil
}

func (s *SyncStore) String() string {
	str := ""
	var keys []int
	s.Range(func(key interface{}, value interface{}) bool {
		val := value.(puppy.Puppy)
		keys = append(keys, int(val.ID))
		return true
	})
	sort.Ints(keys)
	for _, key := range keys {
		if p, ok := s.Load(int16(key)); ok {
			val := p.(puppy.Puppy)
			str += val.String() + " "
		}
	}
	return str
}
