package main

import (
	"sync"
)

type Puppy struct {
	ID     int    `json:"id"`
	Value  string `json:"value"`
	Breed  string `json:"breed"`
	Colour string `json:"colour"`
}

type Storer interface {
	ReadPuppy(ID int) (*Puppy, error)
	UpdatePuppy(ID int, puppy *Puppy) error
	CreatePuppy(puppy *Puppy) int
	DeletePuppy(id int) (bool, error)
}

type SyncStore struct {
	sync.Map
}

type MemStore map[int](*Puppy)
