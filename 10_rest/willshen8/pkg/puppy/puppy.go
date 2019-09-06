package puppy

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
