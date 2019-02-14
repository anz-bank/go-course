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

// StorerMap defines standard CRUD operations for puppy
// This has been created for dynamic ID implementation for each puppy
// I would like to know about extending the "Storer" interface "createPuppy" method with a return integer,
// Otherwise I can use the same interface for MapStore and SyncStore
type StorerMap interface {
	createPuppy(Puppy) uint
	readPuppy(uint) Puppy
	updatePuppy(uint, Puppy)
	deletePuppy(uint)
}
