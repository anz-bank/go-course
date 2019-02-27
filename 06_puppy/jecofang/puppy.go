package main

import "fmt"

type Puppy struct {
	ID    uint64
	Breed string
	Color string
	Value float32
}

func (puppy Puppy) String() string {
	return fmt.Sprintf("Puppy[ID: %d, Breed: %s, Color: %s, Value: %.2f]", puppy.ID, puppy.Breed, puppy.Color, puppy.Value)
}
