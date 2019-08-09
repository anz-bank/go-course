package main

import "sync"

// Puppy defines the data structure corresponding to a Puppy
type Puppy struct {
	ID     uint32
	Breed  string
	Colour string
	Value  string
}

// MapStore stores Puppy details with Puppy Id as Key and Puppy  as value
type MapStore map[uint32]Puppy

// SyncStore stores Puppy details with Puppy Id as Key and Puppy  as value
type SyncStore struct {
	sync.Map
}

// Storer defines the interface on Puppy
type Storer interface {
	CreatePuppy(Puppy) error
	ReadPuppy(ID uint32) (Puppy, error)
	UpdatePuppy(Puppy Puppy) error
	DeletePuppy(ID uint32) error
}
