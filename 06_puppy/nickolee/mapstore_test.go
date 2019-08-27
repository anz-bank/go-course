package puppystorer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Writing one test to give confidence that mapstore is creating and retrieving puppies correctly
func TestCreateAndRetrievePuppyMapStore(t *testing.T) {
	assert := assert.New(t)
	ms := NewMapStore()

	tests := []*Puppy{
		{Breed: "Snoopy", Colour: "Is sleepy", Value: 2300.90},
		{Breed: "Arcanine", Colour: "Level 100", Value: 9300.90},
		{Breed: "The Hound", Colour: "Of Baskerville", Value: 12300.90},
		{Breed: "Laika", Colour: "First dog on moon", Value: 22300.90},
		{Breed: "Simba", Colour: "Pumba", Value: 92300.90},
	}

	for i, testPuppy := range tests {
		createdPuppy := ms.CreatePuppy(testPuppy)
		assert.Equal(i+1, createdPuppy) // test that nextID is working properly
		retrievedPuppy, err := ms.ReadPuppy(createdPuppy)
		assert.NoError(err)
		assert.Equalf(testPuppy, retrievedPuppy, "Testcase TestCreateAndRetrievePuppySyncStore failed")
	}
}
