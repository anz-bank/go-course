package puppystorer

type Puppy struct {
	ID     int
	Breed  string
	Colour string
	Value  float64
}

// Storer defines standard CRUD operations for Pets
type Storer interface {
	CreatePuppy(puppy *Puppy) int // takes an address as an argument which makes sense since you are modifying that object
	ReadPuppy(id int) (*Puppy, error)
	UpdatePuppy(id int, puppy *Puppy) error
	DeletePuppy(id int) error
}
