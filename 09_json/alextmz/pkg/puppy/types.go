package puppy

type Pid uint64
type Pval int32

type Puppy struct {
	ID     Pid    `json:"id"`
	Breed  string `json:"breed"`
	Colour string `json:"colour"`
	Value  Pval   `json:"value"`
}

// Storer defines standard CRUD operations for Puppies
type Storer interface {
	CreatePuppy(*Puppy) error
	ReadPuppy(ID Pid) (*Puppy, error)
	UpdatePuppy(ID Pid, p *Puppy) error
	DeletePuppy(ID Pid) error
}
