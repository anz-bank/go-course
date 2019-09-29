package puppy

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type typetest map[string]struct {
	puppy     Puppy
	puppyJSON string
}

func newPuppyFixture() typetest {
	return typetest{
		"normal puppy": {
			Puppy{ID: 1, Breed: "Wolfhound", Colour: "Gray", Value: 50},
			`{"id":1,"breed":"Wolfhound","colour":"Gray","value": 50}`,
		},
		"empty puppy": {
			Puppy{},
			`{"id":0,"breed":"","colour":"","value":0}`,
		},
		"invalid puppy": {
			Puppy{Value: -100},
			`{"id":0,"breed":"","colour":"","value":-100}`,
		},
	}
}

func TestMarshall(t *testing.T) {
	tests := newPuppyFixture()

	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			got, err := json.Marshal(test.puppy)
			assert.NoError(t, err)
			assert.JSONEq(t, test.puppyJSON, string(got))
		})
	}
}

func TestUnmarshall(t *testing.T) {
	tests := newPuppyFixture()
	for name, test := range tests {
		test := test

		t.Run(name, func(t *testing.T) {
			var p Puppy
			err := json.Unmarshal([]byte(test.puppyJSON), &p)
			assert.NoError(t, err)
			assert.Equal(t, test.puppy, p)
		})
	}
}
