package main

import "fmt"

type Error struct {
	Message string
	Code    int
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error code %d: %s",
		e.Code, e.Message)
}

const (
	NegativeValue    = 1001
	NonExistentPuppy = 1002
)
