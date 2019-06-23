package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Puppy struct {
	ID     int64
	Breed  string
	Colour string
	Value  float64
}

type Storer interface {
	CreatePuppy(p Puppy) (bool, error)
	RetrievePuppy(int64) (*Puppy, error)
	UpdatePuppy(int64, Puppy) (bool, error)
	DeletePuppy(int64) (bool, error)
}

type MapStore map[int64]Puppy

const (
	missingPup      = 0
	invalidPupValue = 1
)

type PupError struct {
	Message string
	Code    int
}

func (e *PupError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func (e *PupError) IsMissingPup() bool {
	return e.Code == missingPup
}

func (e *PupError) InvalidPupValue() bool {
	return e.Code == invalidPupValue
}

func missingPuppy(id int64) error {
	return &PupError{fmt.Sprintf("Puppy not found: %d", id), missingPup}
}

func invalidPuppyValue(value float64) error {
	return &PupError{fmt.Sprintf("Puppy value must be non negative: %f", value), invalidPupValue}
}

var out io.Writer = os.Stdout

func main() {
	p := Puppy{1, "Jack Russel Terrier", "white and brown", 550.0}
	fmt.Fprintf(out, "%d - %s [%s]: %f", p.ID, p.Breed, p.Colour, p.Value)
}

func (ms MapStore) CreatePuppy(p Puppy) (bool, error) {
	if p.Value < 0 {
		return false, invalidPuppyValue(p.Value)
	}

	if _, ok := ms[p.ID]; !ok {
		ms[p.ID] = p
	}
	return true, nil
}

func (ms MapStore) RetrievePuppy(id int64) (*Puppy, error) {
	if _, ok := ms[id]; ok {
		pup := ms[id]
		return &pup, nil
	}
	return nil, missingPuppy(id)
}

func (ms MapStore) UpdatePuppy(id int64, p Puppy) (bool, error) {
	if p.Value < 0 {
		return false, invalidPuppyValue(p.Value)
	}

	if _, ok := ms[id]; ok {
		p.ID = id
		ms[p.ID] = p
		return true, nil
	}
	return false, missingPuppy(id)
}

func (ms MapStore) DeletePuppy(id int64) (bool, error) {
	if _, ok := ms[id]; ok {
		delete(ms, id)
		return true, nil
	}
	return false, missingPuppy(id)
}

type SyncStore struct {
	store *sync.Map
	lock  *sync.Mutex
}

func (s SyncStore) CreatePuppy(p Puppy) (bool, error) {
	if p.Value < 0 {
		return false, invalidPuppyValue(p.Value)
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	s.store.LoadOrStore(p.ID, p)
	return true, nil
}

func (s SyncStore) RetrievePuppy(id int64) (*Puppy, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if pup, ok := s.store.Load(id); ok {
		p := pup.(Puppy)
		return &p, nil
	}
	return nil, missingPuppy(id)
}

func (s SyncStore) UpdatePuppy(id int64, p Puppy) (bool, error) {
	if p.Value < 0 {
		return false, invalidPuppyValue(p.Value)
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.store.Load(id); ok {
		p.ID = id
		s.store.Store(p.ID, p)
		return true, nil
	}
	return false, missingPuppy(id)
}

func (s SyncStore) DeletePuppy(id int64) (bool, error) {

	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.store.Load(id); ok {
		s.store.Delete(id)
		return true, nil
	}
	return false, missingPuppy(id)
}
