package main

import (
	"strconv"
)

// Storer defines standard CRUD operations for Puppy
type Storer interface {
	CreatePuppy(p *Puppy) (int32, error)
	ReadPuppy(ID int32) (*Puppy, error)
	UpdatePuppy(ID int32, Puppy *Puppy) error
	DeletePuppy(ID int32) error
}

// Puppy stores puppy details.
type Puppy struct {
	ID     int32
	Breed  string
	Colour string
	Value  string
}

// mapTest used during testing to verify underlaying map changes
type mapTest interface {
	length() int
}

// Error is a custom error
type Error struct {
	Message string
	Code    int
}

func (e *Error) Error() string {
	if e == nil {
		return "<nil>"
	}
	return e.Message + " Error Code: " + strconv.Itoa(e.Code)
}

const (
	// ecNegativeID error number if id is negative
	ecNegativeID = iota
	// ecNotFound error number if id not found
	ecNotFound = iota
	// ecNotConstructed if the interface was called without being constructed first
	ecNotConstructed = iota
)

// ErrValueBelowZero error generated if the calu is below zero
var ErrValueBelowZero = &Error{Message: "id below 0", Code: ecNegativeID}

// ErrIDNotFound error if the requested ID is not in the store
var ErrIDNotFound = &Error{Message: "id not found", Code: ecNotFound}

// ErrNotConstructed returned if the interface was called without
// first constructing the underlaying structure.
var ErrNotConstructed = &Error{Message: "store not created", Code: ecNotConstructed}
