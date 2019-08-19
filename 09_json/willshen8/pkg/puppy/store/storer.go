package puppy

import "github.com/anz-bank/go-course/09_json/willshen8/pkg/puppy"

type Storer interface {
	CreatePuppy(*puppy.Puppy) (uint32, error)
	ReadPuppy(ID uint32) (*puppy.Puppy, error)
	UpdatePuppy(ID uint32, puppy *puppy.Puppy) (bool, error)
	DeletePuppy(ID uint32) (bool, error)
}
