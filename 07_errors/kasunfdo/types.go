package main

type Puppy struct {
	ID     uint64  `json:"id"`
	Breed  string  `json:"breed"`
	Colour string  `json:"colour"`
	Value  float64 `json:"value"`
}

type Storer interface {
	CreatePuppy(puppy Puppy) (uint64, error)
	ReadPuppy(id uint64) (Puppy, error)
	UpdatePuppy(puppy Puppy) error
	DeletePuppy(id uint64) error
}
