package puppy

import "fmt"

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

// type StorerHTTPHandler struct {
// 	store puppy.Storer
// 	len int
// }

func (p *Puppy) JSONstr() string {
	r := `{"id":%d,"breed":"%s","colour":"%s","value": %.2f}`
	return fmt.Sprintf(r, p.ID, p.Breed, p.Colour, p.Value)
}
