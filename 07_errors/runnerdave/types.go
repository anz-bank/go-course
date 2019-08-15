package main

// Puppy represents an item in the puppy store
type Puppy struct {
	Breed  string  `json:"name"`
	Colour string  `json:"color"`
	ID     uint16  `json:"id"`
	Value  float32 `json:"value"`
}

// Storer CRUD methods for the Puppy store
type Storer interface {
	CreatePuppy(*Puppy) error
	ReadPuppy(ID uint16) (Puppy, error)
	UpdatePuppy(ID uint16, p *Puppy) error
	DeletePuppy(ID uint16) error
}
