package puppy

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJson(t *testing.T) {
	testCases := []struct {
		object    Puppy
		marshaled string
	}{
		{marshaled: `{"id":0,"value":222,"breed":"Type: R","colour":"Red"}`,
			object: Puppy{Colour: "Red", Value: 222, Breed: "Type: R"}},
		{marshaled: `{"id":0,"value":0,"breed":"","colour":""}`,
			object: Puppy{}},
	}
	for _, test := range testCases {
		b, _ := json.Marshal(test.object)
		assert.Equal(t, test.marshaled, string(b))
	}
}
