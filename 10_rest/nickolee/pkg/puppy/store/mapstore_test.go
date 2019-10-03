package store

import (
	"testing"

	"github.com/anz-bank/go-course/10_rest/nickolee/pkg/puppy"
	"github.com/stretchr/testify/assert"
)

// Writing one test to give confidence that mapstore is creating and retrieving puppies correctly
func TestCreateAndRetrievePuppyMapStore(t *testing.T) {
	assert := assert.New(t)
	ms := MapStore{store: make(map[int]puppy.Puppy), nextID: 0}

	tests := []*puppy.Puppy{
		{Breed: "Snoopy", Colour: "Is sleepy", Value: 2300.90},
		{Breed: "Arcanine", Colour: "Level 100", Value: 9300.90},
		{Breed: "The Hound", Colour: "Of Baskerville", Value: 12300.90},
		{Breed: "Laika", Colour: "First dog on moon", Value: 22300.90},
		{Breed: "Simba", Colour: "Pumba", Value: 92300.90},
	}

	for i, testPuppy := range tests {
		createdPuppy, err := ms.CreatePuppy(testPuppy)
		assert.NoError(err)
		assert.Equal(i+1, createdPuppy) // test that nextID is working properly
		retrievedPuppy, err := ms.ReadPuppy(createdPuppy)
		assert.NoError(err)
		assert.Equalf(testPuppy, retrievedPuppy, "Testcase TestCreateAndRetrievePuppySyncStore failed")
	}
}
