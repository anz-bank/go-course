package main

type Puppy struct {
	ID     int
	Breed  string
	Colour string
	Value  float64
}

// Storer defines standard CRUD operations for Puppies
type Storer interface {
	CreatePuppy(*Puppy) error
	ReadPuppy(int) (Puppy, error)
	UpdatePuppy(Puppy) error
	DeletePuppy(int) error
}
