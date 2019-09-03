package puppystorer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/// Writing one test to give confidence that mapstore is retrieving puppies correctly
func TestCreateAndRetrievePuppySyncStore(t *testing.T) {
	assert := assert.New(t)
	ss := SyncStore{}

	tests := []*Puppy{
		{Breed: "Snoopy", Colour: "Is sleepy", Value: 2300.90},
		{Breed: "Arcanine", Colour: "Level 100", Value: 9300.90},
		{Breed: "The Hound", Colour: "Of Baskerville", Value: 12300.90},
		{Breed: "Laika", Colour: "First dog on moon", Value: 22300.90},
		{Breed: "Simba", Colour: "Pumba", Value: 92300.90},
		{Breed: "Direwolf", Colour: "Isn't technically a wolf", Value: 22300.90},
	}

	for i, testPuppy := range tests {
		createdPuppy, err := ss.CreatePuppy(testPuppy)
		assert.NoError(err)
		assert.Equal(i+1, createdPuppy) // test that nextID is working properly
		retrievedPuppy, err := ss.ReadPuppy(testPuppy.ID)
		assert.NoError(err)
		assert.Equalf(testPuppy, retrievedPuppy, "Testcase TestCreateAndRetrievePuppySyncStore failed")
	}
}
