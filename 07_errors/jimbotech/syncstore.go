package main

import (
	"math/rand"
	"sync"
)

// syncStore stores puppies threadsafe.
type syncStore struct {
	m  sync.Mutex
	pm sync.Map
}

// length is not concorrency safe. As the go doc says:
// Range does not necessarily correspond to any consistent snapshot of the
// Map's contents: no key will be visited more than once, but if the value
// for any key is stored or deleted concurrently, Range may reflect any
// mapping for that key from any point during the Range call.
//
func (s *syncStore) length() int {
	var length int
	s.m.Lock()
	defer s.m.Unlock()
	s.pm.Range(func(key interface{}, value interface{}) bool {
		length++
		return true
	})
	return length
}

// CreatePuppy threadsafe adding a puppy to storage
// but will modify the member ID.
func (s *syncStore) CreatePuppy(p *Puppy) (int32, error) {
	s.m.Lock()
	defer s.m.Unlock()
	for notUnique := true; notUnique; p.ID = rand.Int31() {
		_, notUnique = s.pm.Load(p.ID)
	}
	sp := *p
	s.pm.Store(p.ID, &sp)
	return p.ID, nil
}

// ReadPuppy threadsafe retrieval of your puppy.
func (s *syncStore) ReadPuppy(id int32) (*Puppy, error) {
	s.m.Lock()
	defer s.m.Unlock()
	return s.unsafeRetrievePuppy(id)
}

func (s *syncStore) unsafeRetrievePuppy(id int32) (*Puppy, error) {
	if id < 0 {
		return nil, ErrValueBelowZero
	}
	val, found := s.pm.Load(id)
	if !found {
		return nil, ErrIDNotFound
	}
	retPup := *val.(*Puppy)
	return &retPup, nil
}

// UpdatePuppy threadsafe update your puppy store.
func (s *syncStore) UpdatePuppy(id int32, puppy *Puppy) error {
	s.m.Lock()
	defer s.m.Unlock()
	_, err := s.unsafeRetrievePuppy(id)
	if err != nil {
		return err
	}
	puppy.ID = id
	sp := *puppy
	s.pm.Store(id, &sp)
	return nil
}

// DeletePuppy threadsafe removal of the puppy from store.
func (s *syncStore) DeletePuppy(id int32) error {
	s.m.Lock()
	defer s.m.Unlock()
	_, err := s.unsafeRetrievePuppy(id)
	if err != nil {
		return err
	}
	s.pm.Delete(id)
	return nil
}
