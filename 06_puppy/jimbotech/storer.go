package main

// Storer defines standard CRUD operations for Puppy
type Storer interface {
	CreatePuppy(p *Puppy) (uint32, error)
	ReadPuppy(ID uint32) (*Puppy, error)
	UpdatePuppy(ID uint32, Puppy *Puppy) error
	DeletePuppy(ID uint32) error
}

// Puppy stores puppy details.
type Puppy struct {
	ID     uint32
	Breed  string
	Colour string
	Value  string
}

// mapTest used during testing to verify underlaying map changes
type mapTest interface {
	length() int
}
