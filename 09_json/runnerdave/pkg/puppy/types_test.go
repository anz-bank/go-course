package puppy

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshallingPuppies(t *testing.T) {
	tables := []struct {
		in  Puppy
		exp string
	}{
		{Puppy{ID: 11, Breed: "Chihuahua", Colour: "Brown", Value: 12.30},
			`{"breed": "Chihuahua", "color": "Brown", "id": 11, "value": 12.30}`},
		{Puppy{ID: 11, Colour: "Brown", Value: 12.30},
			`{"breed": "", "color": "Brown", "id": 11, "value": 12.30}`},
		{Puppy{},
			`{"breed": "", "id": 0, "value": 0}`},
	}

	for _, table := range tables {
		expected := table.exp
		bytes, _ := json.Marshal(table.in)
		actual := string(bytes)
		require.JSONEq(t, expected, actual, "Marshalling of Puppy (%v) was incorrect, got: %s, want: %s.",
			table.in, actual, expected)
	}
}

func TestUnmarshallingPuppies(t *testing.T) {
	r := require.New(t)
	tables := []struct {
		in  string
		exp Puppy
	}{
		{`{"breed": "Chihuahua", "color": "Brown", "id": 11, "value": 12.30}`,
			Puppy{ID: 11, Breed: "Chihuahua", Colour: "Brown", Value: 12.30}},
		{`{"breed": "", "color": "Brown", "id": 11, "value": 12.30}`,
			Puppy{ID: 11, Colour: "Brown", Value: 12.30}},
		{`{"breed": "", "id": 0, "value": 0}`,
			Puppy{}},
	}

	for _, table := range tables {
		expected := table.exp
		bytes := []byte(table.in)
		actual := Puppy{}
		err := json.Unmarshal(bytes, &actual)
		r.NoError(err, "Should unmarshall without problem")
		assert.Equal(t, expected, actual, "Unmarshalling of Puppy (%v) was incorrect, got: %s, want: %s.",
			table.in, actual, expected)
	}
}

func TestPuppyStringer(t *testing.T) {
	tables := []struct {
		in  Puppy
		exp string
	}{
		{Puppy{ID: 11, Breed: "Chihuahua", Colour: "Brown", Value: 12.30},
			`{"breed": "Chihuahua", "color": "Brown", "id": 11, "value": 12.30}`},
		{Puppy{ID: 11, Colour: "Brown", Value: 12.30},
			`{"breed": "", "color": "Brown", "id": 11, "value": 12.30}`},
		{Puppy{},
			`{"breed": "", "id": 0, "value": 0}`},
	}

	for _, table := range tables {
		expected := table.exp
		actual := table.in.String()
		require.JSONEq(t, expected, actual, "Stringer of Puppy (%v) was incorrect, got: %s, want: %s.",
			table.in, actual, expected)
	}
}
