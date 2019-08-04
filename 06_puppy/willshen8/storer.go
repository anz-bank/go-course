package main

type Storer interface {
	CreatePuppy(*Puppy) uint32
	ReadPuppy(ID uint32) *Puppy
	UpdatePuppy(ID uint32, puppy *Puppy) bool
	DeletePuppy(ID uint32) bool
}
