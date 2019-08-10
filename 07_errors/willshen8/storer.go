package main

type Storer interface {
	CreatePuppy(*Puppy) (uint32, error)
	ReadPuppy(ID uint32) (*Puppy, error)
	UpdatePuppy(ID uint32, puppy *Puppy) (bool, error)
	DeletePuppy(ID uint32) (bool, error)
}
