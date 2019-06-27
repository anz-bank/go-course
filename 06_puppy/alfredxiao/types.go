package main

type Puppy struct {
	ID     string
	Breed  string
	Colour string
	Value  string
}

type Storer interface {
	CreatePuppy(p Puppy) string
	ReadPuppy(ID string) (Puppy, error)
	UpdatePuppy(p Puppy) error
	DeletePuppy(ID string) (bool, error)
}
