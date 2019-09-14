package puppy

import "fmt"

type Error struct {
	Message string
	Code    uint16
}

const (
	_ = iota
	ErrGeneric
	_
	ErrValueLessThanZero
	_
	ErrIDBeingCreatedAlreadyExists
	ErrIDBeingReadDoesNotExist
	ErrIDBeingUpdatedDoesNotExist
	ErrIDBeingDeletedDoesNotExist
)

var ErrTypeLongStrings = []string{
	"error error",
	"error, generic",
	"error, custom",
	"error, value is less than zero",
	"error, ID not found",
	"error, ID already exists",
	"error, ID being read does not exist",
	"error, ID being updated does not exist",
	"error, ID being deleted does not exist",
}

func (e Error) Error() string {
	switch {
	case e.Code == 0 && len(e.Message) == 0:
		e.Code = ErrGeneric
		e.Message = ErrTypeLongStrings[e.Code]
	case e.Code != 0 && len(e.Message) == 0:
		e.Message = ErrTypeLongStrings[e.Code]
	}
	return fmt.Sprintf("%v", e.Message)
}

func NewError(i uint16) Error {
	var e Error
	e.Code = i
	e.Message = e.Error()
	return e
}
