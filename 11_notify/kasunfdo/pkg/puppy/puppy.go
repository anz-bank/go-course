package puppy

type Puppy struct {
	ID     uint64  `json:"id,omitempty"`
	Breed  string  `json:"breed"`
	Colour string  `json:"colour"`
	Value  float64 `json:"value"`
}

func (p *Puppy) Validate() error {
	if p.Value < 0 {
		return ErrorEf(ErrInvalid, nil, "value of puppy is negative")
	}
	return nil
}

type Storer interface {
	CreatePuppy(puppy Puppy) (uint64, error)
	ReadPuppy(id uint64) (Puppy, error)
	UpdatePuppy(puppy Puppy) error
	DeletePuppy(id uint64) error
}

type LostPuppyRequest struct {
	ID uint64 `json:"id"`
}
