package main

// Puppy represents an entry in the puppy store
type Puppy struct {
	ID    int
	Breed string
	Color string
	Value string
}

// Storer CRUD methods for the Puppy store
type Storer interface {
	CreatePuppy(*Puppy) int
	UpdatePuppy(puppy *Puppy) error
	ReadPuppy(ID int) (*Puppy, error)
	DeletePuppy(ID int) error
}
