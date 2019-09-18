package puppy

// Puppy data structure stores puppy properties
type Puppy struct {
	ID     int     `json:"id"`
	Breed  string  `json:"breed"`
	Colour string  `json:"colour"`
	Value  float64 `json:"value"`
}

// Storer defines standard CRUD operations for Pets
type Storer interface {
	CreatePuppy(puppy *Puppy) (int, error) // takes a pointer which makes sense since you are modifying that object
	ReadPuppy(id int) (*Puppy, error)
	UpdatePuppy(id int, puppy *Puppy) error
	DeletePuppy(id int) error
}
