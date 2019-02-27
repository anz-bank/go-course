package main

import "fmt"

const (
	Invalid int = 1 << iota
	Conflict
	NotFound
)

type StorerError struct {
	Code    int
	Message string
}

type Storer interface {
	Create(puppy *Puppy) error
	Read(id uint64) (*Puppy, error)
	Update(id uint64, puppy *Puppy) error
	Delete(id uint64) error
}

func (e *StorerError) Error() string {
	return fmt.Sprintf("%s [Error Code: %d]", e.Message, e.Code)
}

func validateID(id uint64) error {
	if id == 0 {
		return &StorerError{Invalid, "The id should be a positive number."}
	}
	return nil
}
