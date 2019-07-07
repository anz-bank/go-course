package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var out io.Writer = os.Stdout

type Puppy struct {
	ID     uint32
	Breed  string
	Colour string
	Value  int
}

// Storer defines standard CRUD operations for Puppys.
type Storer interface {
	CreatePuppy(*Puppy)
	ReadPuppy(id uint32) *Puppy
	UpdatePuppy(id uint32, puppy *Puppy)
	DeletePuppy(id uint32) bool
}

// MapStore ////

type MapStore map[uint32]Puppy

// CreatePuppy stores a copy of Puppy p.
func (store *MapStore) CreatePuppy(p *Puppy) {
	(*store)[p.ID] = *p
}

// ReadPuppy retrieves a previously stored Puppy.
func (store *MapStore) ReadPuppy(id uint32) *Puppy {
	p, ok := (*store)[id]
	if !ok {
		return nil
	}
	return &p
}

// UpdatePuppy updates details of Puppy in store.
func (store *MapStore) UpdatePuppy(id uint32, puppy *Puppy) {
	if id == puppy.ID {
		(*store)[id] = *puppy
	}
}

// DeletePuppy deletes puppy from store.
func (store *MapStore) DeletePuppy(id uint32) bool {
	_, exists := (*store)[id]
	delete(*store, id)
	return exists
}

// SyncStore ////

type SyncStore struct {
	sync.Map
}

// CreatePuppy stores a copy of Puppy p.
func (store *SyncStore) CreatePuppy(p *Puppy) {
	store.Store(p.ID, Puppy{p.ID, p.Breed, p.Colour, p.Value})
}

// ReadPuppy retrieves a previously stored Puppy.
func (store *SyncStore) ReadPuppy(id uint32) *Puppy {
	generic, ok := store.Load(id)
	if !ok {
		return nil
	}
	puppy := generic.(Puppy)
	return &puppy
}

// UpdatePuppy updates details of Puppy in store.
func (store *SyncStore) UpdatePuppy(id uint32, p *Puppy) {
	if id == p.ID {
		store.Store(p.ID, Puppy{p.ID, p.Breed, p.Colour, p.Value})
	}
}

// DeletePuppy deletes puppy from store.
func (store *SyncStore) DeletePuppy(id uint32) bool {
	_, exists := store.Load(id)
	store.Delete(id)
	return exists
}

func main() {
	storer := MapStore{}
	p1 := Puppy{7, "bulldog", "black", 100}
	storer.CreatePuppy(&p1)
	p2 := storer.ReadPuppy(7)
	fmt.Fprintln(out, *p2)
}
