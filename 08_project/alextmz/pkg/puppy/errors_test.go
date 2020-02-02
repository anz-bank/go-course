package puppy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnownErrorCodeDescr(t *testing.T) {
	var tests = map[string]struct {
		code      int
		parameter string
		want      string
	}{
		"ErrNilPuppyPointer": {ErrNilPuppyPointer, "",
			"puppy pointer is nil"},
		"ErrNegativePuppyValueOnCreate": {ErrNegativePuppyValueOnCreate, "",
			"trying to create a puppy with a negative value"},
		"ErrNegativePuppyValueOnCreate with parameter": {ErrNegativePuppyValueOnCreate, "-555.55",
			"trying to create a puppy with a negative value (-555.55)"},
		"ErrNegativePuppyValueOnUpdate": {ErrNegativePuppyValueOnUpdate, "",
			"trying to update a puppy with a negative value"},
		"ErrNegativePuppyValueOnUpdate with parameter": {ErrNegativePuppyValueOnUpdate, "-22.22",
			"trying to update a puppy with a negative value (-22.22)"},
		"ErrPuppyAlreadyIdentified": {ErrPuppyAlreadyIdentified, "",
			"puppy already initialized"},
		"ErrPuppyAlreadyIdentified with parameter": {ErrPuppyAlreadyIdentified, "42",
			"puppy already initialized with ID 42"},
		"ErrPuppyNotFoundOnRead": {ErrPuppyNotFoundOnRead, "",
			"puppy being read does not exist"},
		"ErrPuppyNotFoundOnRead with parameter": {ErrPuppyNotFoundOnRead, "11",
			"puppy with ID 11 being read does not exist"},
		"ErrPuppyNotFoundOnUpdate": {ErrPuppyNotFoundOnUpdate, "",
			"puppy being updated does not exist"},
		"ErrPuppyNotFoundOnUpdate with parameter": {ErrPuppyNotFoundOnUpdate, "22",
			"puppy with ID 22 being updated does not exist"},
		"ErrPuppyNotFoundOnDelete": {ErrPuppyNotFoundOnDelete, "",
			"puppy being deleted does not exist"},
		"ErrPuppyNotFoundOnDelete with parameter": {ErrPuppyNotFoundOnDelete, "33",
			"puppy with ID 33 being deleted does not exist"},
		"UnknownErrorCode": {
			61621, "",
			"undefined error"},
	}
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			e := Error{Code: test.code}
			got := e.errorCodeDescription(test.parameter)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestError(t *testing.T) {
	e := Error{Code: ErrNilPuppyPointer}
	assert.Equal(t, "puppy pointer is nil", e.Error())
	e = Error{Code: ErrPuppyNotFoundOnDelete}
	assert.Equal(t, "puppy being deleted does not exist", e.Error())
}

func TestErrorp(t *testing.T) {
	// test empyy parameters
	e := Errorp(ErrNilPuppyPointer, "")
	assert.Equal(t, "puppy pointer is nil", e.Error())
	// test string parameters
	e = Errorp(ErrPuppyNotFoundOnDelete, "999")
	assert.Equal(t, "puppy with ID 999 being deleted does not exist", e.Error())
	// test int parameters
	e = Errorp(ErrPuppyNotFoundOnDelete, 999)
	assert.Equal(t, "puppy with ID 999 being deleted does not exist", e.Error())
	// test float parameters
	e = Errorp(ErrPuppyNotFoundOnDelete, 9.99)
	assert.Equal(t, "puppy with ID 9.99 being deleted does not exist", e.Error())
	// test unimplemented parameter type
	assert.Panics(t, func() {
		_ = Errorp(ErrPuppyNotFoundOnDelete, true)
	})
}

func TestNewError(t *testing.T) {
	e1 := NewError(999, "error 999")
	var e2 Error
	e2.Code = 999
	e2.Message = "error 999"
	assert.Equal(t, e2, e1)
}

func TestNewErrorf(t *testing.T) {
	e1 := NewErrorf(999, "error %d", 999)
	var e2 Error
	e2.Code = 999
	e2.Message = "error 999"
	assert.Equal(t, e2, e1)
}
