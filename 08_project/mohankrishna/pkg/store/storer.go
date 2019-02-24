package store

type Storer interface {
	CreatePuppy(puppy *Puppy) error
	ReadPuppy(ID uint32) (*Puppy, error)
	UpdatePuppy(ID uint32, puppy *Puppy) error
	DeletePuppy(ID uint32) error
}
