package main

// Puppy defines the attributes corresponding to puppy
type Puppy struct {
	id    uint
	breed string
	color string
	value float32
}

// Storer defines standard CRUD operations for puppy
type Storer interface {
	createPuppy(Puppy)
	readPuppy(uint) Puppy
	updatePuppy(uint, Puppy)
	deletePuppy(uint)
}
