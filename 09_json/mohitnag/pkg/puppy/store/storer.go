package store

import (
	"github.com/anz-bank/go-course/09_json/mohitnag/pkg/puppy"
)

// Puppy is an alias for puppy.Puppy
type Puppy = puppy.Puppy

// Storer defines the interface on Puppy
type Storer interface {
	CreatePuppy(Puppy) error
	ReadPuppy(ID uint32) (Puppy, error)
	UpdatePuppy(Puppy Puppy) error
	DeletePuppy(ID uint32) error
}
