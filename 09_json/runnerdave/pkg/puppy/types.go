package puppy

import "encoding/json"

// Puppy represents an item in the puppy store
type Puppy struct {
	Breed  string  `json:"breed"`
	Colour string  `json:"color,omitempty"`
	ID     int16   `json:"id"`
	Value  float32 `json:"value"`
}

func (p *Puppy) String() string {
	s, _ := json.Marshal(p)
	return string(s)
}

// Storer CRUD methods for the Puppy store
type Storer interface {
	CreatePuppy(Puppy) error
	ReadPuppy(ID int16) (Puppy, error)
	UpdatePuppy(ID int16, p *Puppy) error
	DeletePuppy(ID int16) error
}
