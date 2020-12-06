package store

import (
	p "github.com/anz-bank/go-course/08_project/patrickmarabeas/pkg/puppy"
)

type Storer interface {
	Create(puppy p.Puppy) (int, error)
	Read(ID int) (p.Puppy, error)
	Update(ID int, puppy p.Puppy) (bool, error)
	Destroy(ID int) (bool, error)
}
