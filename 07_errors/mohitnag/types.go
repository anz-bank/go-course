package main

// Puppy defines the data structure corresponding to a Puppy
type Puppy struct {
	ID     uint32
	Breed  string
	Colour string
	Value  string
}

// Storer defines the interface on Puppy
type Storer interface {
	CreatePuppy(Puppy) error
	ReadPuppy(ID uint32) (Puppy, error)
	UpdatePuppy(Puppy Puppy) error
	DeletePuppy(ID uint32) error
}
