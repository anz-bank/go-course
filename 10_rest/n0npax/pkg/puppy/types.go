package puppy

// Puppy contains information about single puppy
type Puppy struct {
	ID     int    `json:"id"`
	Value  int    `json:"value"`
	Breed  string `json:"breed"`
	Colour string `json:"colour"`
}

// Storer interface for Store implementations
type Storer interface {
	ReadPuppy(ID int) (*Puppy, error)
	UpdatePuppy(ID int, puppy *Puppy) error
	CreatePuppy(puppy *Puppy) (int, error)
	DeletePuppy(ID int) (bool, error)
}
