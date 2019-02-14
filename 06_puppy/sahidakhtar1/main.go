package main

type Puppy struct {
	ID                   uint
	Breed, Colour, Value string
}

type Storer interface {
	CreatePuppy(*Puppy)
	ReadPuppy(int) *Puppy
	UpdatePuppy(int, *Puppy)
	DeletePuppy(int)
}

func main() {

}
