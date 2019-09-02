package puppy

// Puppy defines the data structure corresponding to a Puppy
type Puppy struct {
	ID     uint32 `json:"puppyid"`
	Breed  string `json:"breed"`
	Colour string `json:"colour"`
	Value  string `json:"value"`
}
