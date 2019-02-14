package main

// Puppy defines the attributes corresponding to puppy
type Puppy struct {
	id    uint32
	breed string
	color string
	value float32
}

// Storer defines standard CRUD operations for puppy
type Storer interface {
	createPuppy(Puppy)
	readPuppy(uint32) Puppy
	updatePuppy(uint32, Puppy)
	deletePuppy(uint32)
}
