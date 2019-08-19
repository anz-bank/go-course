package main

type Puppy struct {
	ID    uint
	Breed string
	Color string
	Value string
}

type Storer interface {
	Create(puppy Puppy) uint
	Read(ID uint) Puppy
	Update(ID uint, puppy Puppy) bool
	Destroy(ID uint) bool
}
