package puppy

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonMarshaling(t *testing.T) {
	type test struct {
		name         string
		puppyObject  Puppy
		expectedJSON string
	}

	tests := []test{
		{name: "test all fields populated", puppyObject: Puppy{ID: 7, Breed: "Firefox", Colour: "Red", Value: 777},
			expectedJSON: `{"id":7,"breed":"Firefox","colour":"Red","value":777}`},
		{name: "test id field nil", puppyObject: Puppy{Breed: "Firefox", Colour: "Red", Value: 777},
			expectedJSON: `{"id":0,"breed":"Firefox","colour":"Red","value":777}`},
		{name: "test nil puppy object", puppyObject: Puppy{},
			expectedJSON: `{"id":0,"breed":"","colour":"","value":0}`},
	}

	for _, test := range tests {
		jsonRepresentation, _ := json.Marshal(test.puppyObject) // note that jsonReprentation is a slice of bytes
		assert.JSONEq(t, test.expectedJSON, string(jsonRepresentation),
			fmt.Sprintf("Test case: %s failed", test.name))
	}
}

func TestJsonUnmarshaling(t *testing.T) {
	type test struct {
		name          string
		jsonPuppy     []byte
		expectedPuppy Puppy
	}

	tests := []test{
		{name: "test all fields populated", expectedPuppy: Puppy{ID: 7, Breed: "Firefox", Colour: "Red", Value: 777},
			jsonPuppy: []byte(`{"id":7,"breed":"Firefox","colour":"Red","value":777}`)},
		{name: "test id field nil", expectedPuppy: Puppy{ID: 0, Breed: "Firefox", Colour: "Red", Value: 777},
			jsonPuppy: []byte(`{"id":0,"breed":"Firefox","colour":"Red","value":777}`)},
		{name: "test nil puppy object", expectedPuppy: Puppy{},
			jsonPuppy: []byte(`{"id":0,"breed":"","colour":"","value":0}`)},
	}

	for _, test := range tests {
		var unmarshaled Puppy
		err := json.Unmarshal(test.jsonPuppy, &unmarshaled) // note that jsonReprentation is a slice of bytes
		assert.NoError(t, err)
		assert.Equal(t, test.expectedPuppy, unmarshaled, fmt.Sprintf("Test case: %s failed", test.name))
	}
}
