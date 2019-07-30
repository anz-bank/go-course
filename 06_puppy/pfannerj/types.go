package main

//Puppy encapsulates important details about the Puppy including ID, Breed, Colour and Value
type Puppy struct {
	ID     uint32  `json:"id"`
	Breed  string  `json:"breed"`
	Colour string  `json:"colour"`
	Value  float64 `json:"value"`
}

// Storer defines standard CRUD operations for Puppy
type Storer interface {
	CreatePuppy(Puppy) (uint32, error)
	ReadPuppy(ID uint32) (Puppy, error)
	UpdatePuppy(ID uint32, puppy Puppy) (uint32, error)
	DeletePuppy(ID uint32) error
}
