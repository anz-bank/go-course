package main

import "sync"

type Puppy struct {
	ID    int
	Breed string
	Color string
	Value int // cents
}

type Storer interface {
	Create(puppy Puppy) (int, error)
	Read(ID int) (Puppy, error)
	Update(ID int, puppy Puppy) (bool, error)
	Destroy(ID int) (bool, error)
}

type MapStore struct {
	uuid  int
	store map[int]Puppy
}

type SyncStore struct {
	uuid int
	sync.Map
}
