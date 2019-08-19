package main

import (
	"sync"
)

type Puppy struct {
	ID    int
	Breed string
	Color string
	Value string
}

type Storer interface {
	Create(puppy Puppy) int
	Read(ID int) Puppy
	Update(ID int, puppy Puppy) bool
	Destroy(ID int) bool
}

type MapStore struct {
	uuid  int
	store map[int]Puppy
}

type SyncStore struct {
	uuid int
	sync.Map
}
