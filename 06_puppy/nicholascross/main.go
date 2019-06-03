package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Puppy struct {
	ID     uint64
	Breed  string
	Colour string
	Value  float64
}

type Storer interface {
	CreatePuppy(p Puppy)
	RetrievePuppy(uint64) *Puppy
	UpdatePuppy(uint64, Puppy)
	DeletePuppy(uint64)
}

type MapStore map[uint64]Puppy

var out io.Writer = os.Stdout

func main() {
	p := Puppy{1, "Jack Russel Terrier", "white and brown", 550.0}
	fmt.Fprintf(out, "%d - %s [%s]: %f", p.ID, p.Breed, p.Colour, p.Value)
}

func (ms MapStore) CreatePuppy(p Puppy) {
	if _, exists := ms[p.ID]; !exists {
		ms[p.ID] = p
	}
}

func (ms MapStore) RetrievePuppy(id uint64) *Puppy {
	if _, exists := ms[id]; exists {
		pup := ms[id]
		return &pup
	}
	return nil
}

func (ms MapStore) UpdatePuppy(id uint64, p Puppy) {
	if _, exists := ms[id]; exists {
		p.ID = id
		ms[p.ID] = p
	}
}

func (ms MapStore) DeletePuppy(id uint64) {
	if _, exists := ms[id]; exists {
		delete(ms, id)
	}
}

type SyncStore struct {
	store *sync.Map
}

func (s SyncStore) CreatePuppy(p Puppy) {
	s.store.LoadOrStore(p.ID, p)
}

func (s SyncStore) RetrievePuppy(id uint64) *Puppy {
	if pup, exists := s.store.Load(id); exists {
		p := pup.(Puppy)
		return &p
	}
	return nil
}

func (s SyncStore) UpdatePuppy(id uint64, p Puppy) {
	if _, exists := s.store.Load(id); exists {
		p.ID = id
		s.store.Store(p.ID, p)
	}
}

func (s SyncStore) DeletePuppy(id uint64) {
	if _, exists := s.store.Load(id); exists {
		s.store.Delete(id)
	}
}
