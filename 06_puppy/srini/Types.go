package main

//PuppyRecord is internal type with embedded puppy
type PuppyRecord struct {
	id    int
	puppy Puppy
}

//Puppy is struct for Puppy bread
type Puppy struct {
	breed  string
	colour string
	value  float64
}

//Storer provides CRUD methods
type Storer interface {
	CreatePuppy(*Puppy) int
	ReadPuppy(int) (*Puppy, error)
	UpdatePuppy(int, *Puppy) error
	DeletePuppy(int) error
}
