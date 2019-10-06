package puppy

import "strconv"

type Puppy struct {
	ID    uint32 `json:"id"`
	Breed string `json:"breed"`
	Color string `json:"color"`
	Value string `json:"value,omitempty"`
}

type Storer interface {
	CreatePuppy(*Puppy) (uint32, error)
	ReadPuppy(ID uint32) (*Puppy, error)
	UpdatePuppy(ID uint32, puppy *Puppy) error
	DeletePuppy(ID uint32) error
}

// Validate() function validate a puppy's value
func (p *Puppy) Validate() error {
	if i, err := strconv.Atoi(p.Value); err != nil {
		return Errorf(ErrInvalidInput, "Puppy value not a number")
	} else if i < 0 {
		return Errorf(ErrInvalidInput, "Puppy value can't be negative")
	}
	return nil
}
