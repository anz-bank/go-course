package main

type Puppy struct {
	ID                   uint
	Breed, Colour, Value string
}

type Storer interface {
	CreatePuppy(Puppy)
	ReadPuppy(uint) Puppy
	UpdatePuppy(uint, Puppy)
	DeletePuppy(uint) bool
}
