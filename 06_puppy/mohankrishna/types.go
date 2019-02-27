package main

// Error codes
const (
	ErrInvalidInput int = iota
	ErrDuplicate
	ErrNotFound
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

type Puppy struct {
	ID     uint32
	Breed  string
	Colour string
	Value  int
}

type Storer interface {
	CreatePuppy(puppy *Puppy) error
	ReadPuppy(ID uint32) (*Puppy, error)
	UpdatePuppy(ID uint32, puppy *Puppy) error
	DeletePuppy(ID uint32) (bool, error)
}
