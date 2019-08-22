package puppy

import (
	"encoding/json"
	"testing"

	puppy "github.com/anz-bank/go-course/09_json/willshen8/pkg/puppy"
	"github.com/stretchr/testify/assert"
)

func TestUnMarshallJsonData(t *testing.T) {
	input := `{
        "ID": 1,
        "Breed": "Jack Russell Terrier",
        "Color": "White and Brown",
        "Value": "1500"
	}`

	expected := puppy.Puppy{ID: 1, Breed: "Jack Russell Terrier",
		Color: "White and Brown", Value: "1500"}
	var result puppy.Puppy
	unMarshalError := json.Unmarshal([]byte(input), &result)
	if unMarshalError != nil {
		panic(unMarshalError)
	}

	assert.Equal(t, expected, result)
}

func TestMarshal(t *testing.T) {
	input := puppy.Puppy{ID: 1, Breed: "Jack Russell Terrier",
		Color: "White and Brown", Value: "1500"}

	expected := `{"breed":"Jack Russell Terrier", "color":"White and Brown", "id":1, "value":"1500"}`
	result, marshalError := json.Marshal(input)
	if marshalError != nil {
		panic(marshalError)
	}
	assert.JSONEq(t, expected, string(result))
}
