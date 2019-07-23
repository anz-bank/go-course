package puppy

type Puppy struct {
	ID    uint32 `json:"id"`
	Breed string `json:"breed"`
	Color string `json:"color"`
	Value string `json:"value,omitempty"`
}
