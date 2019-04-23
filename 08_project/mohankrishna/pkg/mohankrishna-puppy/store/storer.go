package store

import (
	types "github.com/anz-bank/go-training/08_project/mohankrishna/pkg/mohankrishna-puppy"
)

type Storer interface {
	CreatePuppy(puppy *types.Puppy) error
	ReadPuppy(ID uint32) (*types.Puppy, error)
	UpdatePuppy(ID uint32, puppy *types.Puppy) error
	DeletePuppy(ID uint32) error
}
