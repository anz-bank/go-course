package puppy

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_marshall(t *testing.T) {
	var tests typetest
	buildTestvar(&tests)

	for name, tt := range tests {
		test := tt
		t.Run(name, func(t *testing.T) {
			got, err := json.Marshal(test.obj)
			assert.NoError(t, err)
			assert.JSONEq(t, test.jsn, string(got))
		})
	}
}

func Test_unmarshall(t *testing.T) {
	var tests typetest
	buildTestvar(&tests)

	for name, tt := range tests {
		test := tt
		t.Run(name, func(t *testing.T) {
			var p Puppy
			err := json.Unmarshal([]byte(test.jsn), &p)
			assert.NoError(t, err)
			assert.JSONEq(t, test.obj.JSONstr(), p.JSONstr())
		})
	}
}

type typetest map[string]struct {
	obj Puppy
	jsn string
}

func buildTestvar(t *typetest) {
	*t = typetest{
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
