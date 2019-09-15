package puppy

type Puppy struct {
	ID     int     `json:"id"`
	Breed  string  `json:"breed"`
	Colour string  `json:"colour"`
	Value  float64 `json:"value"`
}

// Storer defines standard CRUD operations for Puppies
type Storer interface {
	CreatePuppy(*Puppy) error
	ReadPuppy(int) (Puppy, error)
	UpdatePuppy(Puppy) error
	DeletePuppy(int) error
}
