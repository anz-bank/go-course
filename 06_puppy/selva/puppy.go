package main

// Puppy struct
type Puppy struct {
	ID     int
	Breed  string
	Colour string
	Value  int
}

//Storer interface
type Storer interface {
	CreatePuppy(Puppy) int
	ReadPuppy(ID int) Puppy
	UpdatePuppy(ID int, pup Puppy)
	DeletePuppy(ID int) bool
}
