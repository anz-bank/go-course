package main

type Puppy struct {
	ID     uint32 `json:"ID"`
	Breed  string `json:"breed"`
	Colour string `json:"colour"`
	Value  string `json:"value"`
}

func (p *Puppy) Clone() *Puppy {
	puppy := new(Puppy)
	puppy.ID = p.ID
	puppy.Breed = p.Breed
	puppy.Colour = p.Colour
	puppy.Value = p.Value
	return puppy
}

type Storer interface {
	CreatePuppy(*Puppy) error
	ReadPuppy(id uint32) (*Puppy, error)
	UpdatePuppy(id uint32, puppy *Puppy) error
	DeletePuppy(id uint32) (bool, error)
}
