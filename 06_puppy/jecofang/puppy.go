package main

import (
	"encoding/json"
	"fmt"
)

type Puppy struct {
	ID    uint64
	Breed string
	Color string
	Value float32
}

func (p Puppy) String() string {
	return fmt.Sprintf("Puppy[ID: %d, Breed: %s, Color: %s, Value: %.2f]", p.ID, p.Breed, p.Color, p.Value)
}

func (p Puppy) ToJSON() []byte {
	b, _ := json.Marshal(p)
	return b
}

func (p *Puppy) ParseJSON(b []byte) error {
	err := json.Unmarshal(b, p)
	return err
}
