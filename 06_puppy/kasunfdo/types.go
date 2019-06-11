package main

// Puppy contains information about puppy
type Puppy struct {
	ID     uint64
	Breed  string
	Colour string
	Value  float64
}

// Storer interface defines CRUD methods for Store implementations
type Storer interface {
	CreatePuppy(puppy Puppy) uint64
	ReadPuppy(id uint64) (Puppy, error)
	UpdatePuppy(puppy Puppy) error
	DeletePuppy(id uint64) error
}
