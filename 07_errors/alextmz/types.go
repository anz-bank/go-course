package main

type PuppyID uint64
type PuppyVal int32

type Puppy struct {
	ID     PuppyID  //`json:"id"`
	Breed  string   //`json:"breed"`
	Colour string   //`json:"colour"`
	Value  PuppyVal //`json:"value"`
}

// Storer defines standard CRUD operations for Puppies
type Storer interface {
	CreatePuppy(*Puppy) error
	ReadPuppy(ID PuppyID) (*Puppy, error)
	UpdatePuppy(ID PuppyID, p *Puppy) error
	DeletePuppy(ID PuppyID) error
}
