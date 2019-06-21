package main

import (
	"fmt"
	"io"
	"os"
	"sync" //https://golang.org/pkg/sync/#Map
)

var out io.Writer = os.Stdout

type Puppy struct {
	pid    int
	breed  string
	colour string
	value  string
}

type MapStore struct {
	store map[int]Puppy
}

type SyncStore struct {
	store sync.Map
}

type Storer interface {
	CreatePuppy(*Puppy) int
	ReadPuppy(int) *Puppy
	UpdatePuppy(*Puppy)
	DeletePuppy(int)
}

func (m *MapStore) CreatePuppy(p *Puppy) int {
	m.store[p.pid] = *p
	return p.pid
}
func (m *MapStore) ReadPuppy(pid int) *Puppy {
	pup, ok := m.store[pid]
	if ok {
		return &pup
	}
	return nil
}
func (m *MapStore) UpdatePuppy(p *Puppy) {
	m.store[p.pid] = *p
}
func (m *MapStore) DeletePuppy(pid int) {
	delete(m.store, pid)
}

func (s *SyncStore) CreatePuppy(p *Puppy) int {
	s.store.Store(p.pid, p)
	return p.pid
}
func (s *SyncStore) ReadPuppy(pid int) *Puppy {
	pup, ok := s.store.Load(pid)
	if ok {
		return pup.(*Puppy)
	}
	return nil
}
func (s *SyncStore) UpdatePuppy(p *Puppy) {
	s.store.Store(p.pid, p)
}
func (s *SyncStore) DeletePuppy(pid int) {
	s.store.Delete(pid)
}

// using the map store + methods
func usingNormMap(pups []Puppy) {
	// initialise the map store
	var puppyMS Storer = &MapStore{store: make(map[int]Puppy)}
	for _, n := range pups {
		pup := n
		puppyMS.CreatePuppy(&pup)
	}
	if p := puppyMS.ReadPuppy(pups[1].pid); p != nil {
		fmt.Fprintln(out, *p)
	}

	upDog := pups[0]
	upDog.colour = "red"
	puppyMS.UpdatePuppy(&upDog)
	puppyMS.DeletePuppy(pups[1].pid)

	fmt.Fprintln(out, puppyMS)
}

// using the sync map store + methods
func usingSyncMap(pups []Puppy) {
	// initialise the sync store
	var puppySS Storer = &SyncStore{}
	for _, n := range pups {
		pup := n
		puppySS.CreatePuppy(&pup)
	}
	if p := puppySS.ReadPuppy(pups[1].pid); p != nil {
		fmt.Fprintln(out, *p)
	}

	upDog := pups[0]
	upDog.colour = "red"
	puppySS.UpdatePuppy(&upDog)
	puppySS.DeletePuppy(pups[1].pid)

	for _, pup := range pups {
		if p := puppySS.ReadPuppy(pup.pid); p != nil {
			fmt.Fprintln(out, *p)
		}
	}
}

func main() {
	pups := []Puppy{
		{99, "poodle", "blue", "$10.99"},
		{100, "lab", "orange", "$9.99"},
		{101, "cat", "striped", "$99.99"},
	}
	usingNormMap(pups)
	usingSyncMap(pups)
}
