package main

//Storer provides CRUD methods
type Storer interface {
	CreatePuppy(*Puppy) int
	ReadPuppy(int) (*Puppy, error)
	UpdatePuppy(int, *Puppy) error
	DeletePuppy(int) error
}
