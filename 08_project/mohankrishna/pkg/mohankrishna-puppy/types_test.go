package puppy

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypesPuppy(t *testing.T) {
	// Given
	r := require.New(t)
	pup := Puppy{
		ID:     0x1236,
		Breed:  "Sheep herder",
		Colour: "Brown",
		Value:  1000,
	}
	actual := fmt.Sprint(pup)
	r.Equalf("{ID:4662, Breed:Sheep herder, Colour:Brown, Value:1000}", actual, "Unexpected error message")
}

func TestJSONMarshal(t *testing.T) {
	// Given
	r := require.New(t)
	pup := Puppy{
		ID:     0x1236,
		Breed:  "Sheep herder",
		Colour: "Brown",
		Value:  1000,
	}
	data, _ := json.Marshal(pup)
	actual := string(data)
	expected := `{"id":4662, "breed":"Sheep herder", "colour":"Brown", "value":1000}`
	r.JSONEq(expected, actual, "Unexpected error message")
}

func TestJSONUNMarshal(t *testing.T) {
	// Given
	r := require.New(t)
	jsonData := []byte(`{"id":4662,"breed":"Sheep herder","colour":"Brown","value":1000}`)
	actual := Puppy{}
	_ = json.Unmarshal(jsonData, &actual)
	expected := Puppy{
		ID:     0x1236,
		Breed:  "Sheep herder",
		Colour: "Brown",
		Value:  1000,
	}
	r.Equal(expected, actual, "Unexpected error message")
}
