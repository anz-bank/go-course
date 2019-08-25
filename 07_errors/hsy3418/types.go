package main

// Puppy defines the data structure corresponding to a pet
type Puppy struct {
	ID     int32   `json:"id"`
	Value  float32 `json:"value"`
	Breed  string  `json:"breed"`
	Colour string  `json:"colour"`
}

//Storer define standard CRUD operations for puppys
type Storer interface {
	CreatePuppy(Puppy) error
	ReadPuppy(ID int32) (Puppy, error)
	UpdatePuppy(puppy Puppy) error
	DeletePuppy(ID int32) error
}
