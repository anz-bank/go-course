package puppy

import (
	"fmt"
)

type Puppy struct {
	ID     uint32 `json:"id"`
	Breed  string `json:"breed"`
	Colour string `json:"colour"`
	Value  int    `json:"value"`
}

func (puppy Puppy) String() string {
	return fmt.Sprintf("{ID:%d, Breed:%s, Colour:%s, Value:%d}", puppy.ID, puppy.Breed, puppy.Colour, puppy.Value)
}
