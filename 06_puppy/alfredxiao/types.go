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
	UpdatePuppy(ID string, p Puppy) error
	DeletePuppy(ID string) error
}
