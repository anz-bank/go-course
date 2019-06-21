package main

type Puppy struct {
  ID     string
  Breed  string
  Colour string
  Value  string
}

type Storer interface {
  CreatePuppy(p Puppy) error
	ReadPuppy(ID string) (Puppy, error)
	UpdatePuppy(p Puppy) error
	DeletePet(ID string) (bool, error)
}
