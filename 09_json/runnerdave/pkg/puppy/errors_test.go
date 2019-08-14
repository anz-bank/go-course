package puppy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateValue(t *testing.T) {
	tables := []struct {
		in  Puppy
		exp error
	}{
		{Puppy{ID: 11, Breed: "Chihuahua", Colour: "Brown", Value: 12.30}, nil},
		{Puppy{ID: 11, Colour: "Brown", Value: 12.30}, nil},
		{Puppy{ID: 15, Colour: "Silver", Value: 0}, nil},
		{Puppy{}, nil},
		{Puppy{ID: 11, Colour: "Brown", Value: -12.30}, Errorf(0x1, "puppy has invalid value (-12.300000)")},
	}
	for _, table := range tables {
		expected := table.exp
		actual := ValidateValue(table.in.Value)
		assert.Equalf(t, expected, actual, "Validate value of Puppy %v was incorrect, expected: %v but got %v",
			expected, expected, actual)
	}
}

func TestPuppyError(t *testing.T) {
	tables := []struct {
		in  Error
		exp string
	}{
		{Error{Code: ErrIDNotFound, Message: "Not found"},
			"2: Not found"},
		{Error{Code: ErrInvalidValue, Message: "Invalid value"},
			"1: Invalid value"},
		{Error{Code: ErrUnknown, Message: "Unknown error"},
			"0: Unknown error"},
	}

	for _, table := range tables {
		expected := table.exp
		actual := table.in.Error()
		assert.Equal(t, expected, actual, "Error of Puppy (%v) was incorrect, got: %s, want: %s.",
			table.in, actual, expected)
	}
}
