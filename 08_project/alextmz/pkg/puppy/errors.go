package puppy

import (
	"fmt"
)

type Error struct {
	Message string
	Code    int
}

const (
	ErrNilPuppyPointer = iota
	ErrNegativePuppyValueOnCreate
	ErrNegativePuppyValueOnUpdate
	ErrPuppyAlreadyIdentified
	ErrPuppyNotFoundOnRead
	ErrPuppyNotFoundOnUpdate
	ErrPuppyNotFoundOnDelete
	ErrNilPuppyPointerStr            = "puppy pointer is nil%s"
	ErrNegativePuppyValueOnCreateStr = "trying to create a puppy with a negative value%s"
	ErrNegativePuppyValueOnUpdateStr = "trying to update a puppy with a negative value%s"
	ErrPuppyAlreadyIdentifiedStr     = "puppy already initialized%s"
	ErrPuppyNotFoundOnReadStr        = "puppy%s being read does not exist"
	ErrPuppyNotFoundOnUpdateStr      = "puppy%s being updated does not exist"
	ErrPuppyNotFoundOnDeleteStr      = "puppy%s being deleted does not exist"
)

// errorCodeDescription returns the verbose error description corresponding
// to a known/static error Code, allowing for an optional parameter to be passed
// for a more verbose error'ing.
// Passing an empty string makes it return the default error string only.
func (e Error) errorCodeDescription(param string) string {
	errormap := map[int]struct {
		errmsg, paramprefix, paramsuffix string
	}{
		ErrNilPuppyPointer:            {ErrNilPuppyPointerStr, "", ""},
		ErrNegativePuppyValueOnCreate: {ErrNegativePuppyValueOnCreateStr, " (", ")"},
		ErrNegativePuppyValueOnUpdate: {ErrNegativePuppyValueOnUpdateStr, " (", ")"},
		ErrPuppyAlreadyIdentified:     {ErrPuppyAlreadyIdentifiedStr, " with ID ", ""},
		ErrPuppyNotFoundOnRead:        {ErrPuppyNotFoundOnReadStr, " with ID ", ""},
		ErrPuppyNotFoundOnUpdate:      {ErrPuppyNotFoundOnUpdateStr, " with ID ", ""},
		ErrPuppyNotFoundOnDelete:      {ErrPuppyNotFoundOnDeleteStr, " with ID ", ""},
	}

	v, ok := errormap[e.Code]
	if !ok {
		return "undefined error"
	}

	if param != "" {
		return fmt.Sprintf(v.errmsg, v.paramprefix+param+v.paramsuffix)
	}

	return fmt.Sprintf(v.errmsg, "")
}

func (e Error) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprint(e.errorCodeDescription(""))
}

// Errorp returns the known parametrized (verbose) error string for
// a given error code.
func Errorp(err int, param interface{}) Error {
	var e Error
	e.Code = err
	switch v := param.(type) {
	case string:
		e.Message = e.errorCodeDescription(v)
	case int:
		e.Message = e.errorCodeDescription(fmt.Sprintf("%d", param))
	case float64:
		e.Message = e.errorCodeDescription(fmt.Sprintf("%.2f", param))
	default:
		panic("not implemented: param type is not either string, int or float")
	}
	return e
}

func NewError(err int, m string) Error {
	var e Error
	e.Code = err
	e.Message = m
	return e
}

func NewErrorf(err int, f string, m ...interface{}) Error {
	var e Error
	e.Code = err
	e.Message = fmt.Sprintf(f, m...)
	return e
}
