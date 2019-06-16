package main

import (
	"sync"
)

type Puppy struct {
	ID     uint32
	Breed  string
	Colour string
	Value  int
}

// Storer defines standard CRUD operations for Puppys
type Storer interface {
	CreatePuppy(*Puppy)
	ReadPuppy(id uint32) *Puppy
	UpdatePuppy(id uint32, puppy *Puppy)
	DeletePuppy(id uint32) bool
}

// MapStore ////

type MapStore map[uint32]Puppy

func (store *MapStore) CreatePuppy(p *Puppy) {
	(*store)[p.ID] = Puppy{p.ID, p.Breed, p.Colour, p.Value}
}

func (store *MapStore) ReadPuppy(id uint32) *Puppy {
	p, ok := (*store)[id]
	if !ok {
		return nil
	}
	return &p
}

func (store *MapStore) UpdatePuppy(id uint32, puppy *Puppy) {
	(*store)[id] = Puppy{puppy.ID, puppy.Breed, puppy.Colour, puppy.Value}
}

func (store *MapStore) DeletePuppy(id uint32) bool {
	_, exists := (*store)[id]
	delete(*store, id)
	return exists
}

// SyncStore ////

// type SyncStore sync.Map[uint32]Puppy
type SyncStore struct {
	sync.Map
}

func (store *SyncStore) CreatePuppy(p *Puppy) {
	store.Store(p.ID, Puppy{p.ID, p.Breed, p.Colour, p.Value})
}

func (store *SyncStore) ReadPuppy(id uint32) *Puppy {
	generic, ok := store.Load(id)
	if !ok {
		return nil
	}
	puppy := generic.(Puppy)
	return &puppy
}

func (store *SyncStore) UpdatePuppy(id uint32, p *Puppy) {
	store.Store(p.ID, Puppy{p.ID, p.Breed, p.Colour, p.Value})
}

func (store *SyncStore) DeletePuppy(id uint32) bool {
	_, exists := store.Load(id)
	store.Delete(id)
	return exists
}

func main() {
}
