package puppystorer

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type StorerSuite struct {
	suite.Suite
	store Storer
}

// A bunch of seed puppies
var puppies = []*Puppy{
	{Breed: "Laika", Colour: "First dog on moon", Value: 2130.50},
	{Breed: "Pomsky", Colour: "Violet", Value: 777.77},
	{Breed: "Dalmatian", Colour: "Cruella de vil", Value: 101},
	{Breed: "Rare Lvl 100 Arcanine", Colour: "Brown", Value: 9999.99},
	{Breed: "Cheap Pug", Colour: "Cheap Brown", Value: -500.50},
}

func (suite *StorerSuite) TestCreateAndRetrievePuppy() {
	puppyID, err := suite.store.CreatePuppy(puppies[3])
	suite.Nil(err)
	puppy, err := suite.store.ReadPuppy(puppyID)
	suite.Nil(err)
	suite.Equal(puppy, puppies[3])

	// Test negative value puppy
	_, err = suite.store.CreatePuppy(puppies[4])
	suite.Equal("Error code 400: Sorry puppy value cannot be negative. The dog has to be worth something :)",
		err.Error())
}

func (suite *StorerSuite) TestUpdatePuppy() {
	// test happy path
	id, err := suite.store.CreatePuppy(puppies[0])
	suite.Nil(err)
	err = suite.store.UpdatePuppy(id, puppies[3])
	suite.Nil(err) // if no error it means puppy was updated successfully
	// retrieve to double check update was correct in the following assertion on next line
	retrievedPuppy, err := suite.store.ReadPuppy(id)
	suite.Nil(err)
	suite.Equal(puppies[3], retrievedPuppy)

	// test error path
	err = suite.store.UpdatePuppy(1829246, puppies[1]) // give non-existent id
	suite.Equal("Error code 404: Sorry puppy with ID 1829246 does not exist", err.Error())
}

func (suite *StorerSuite) TestDeletePuppy() {
	// Test happy path
	id, err := suite.store.CreatePuppy(puppies[1])
	suite.NoError(err)
	err = suite.store.DeletePuppy(id)
	suite.NoError(err) // if nil means successfully deleted
	// verify delete was successful by trying to retrieve puppy
	retrievedPuppy, err := suite.store.ReadPuppy(id)
	suite.Nil(retrievedPuppy) // should get back a nil pointer since there ain't no more puppy
	suite.Error(err)          // asserting there should be an error since you cannot retrieve a deleted puppy

	// Test error path
	err = suite.store.DeletePuppy(1829246)
	suite.Equal("Error code 404: Sorry puppy with ID 1829246 does not exist", err.Error())
}

// Run tests for Storer using MapStore implementation of Storer interface
func TestMapStore(t *testing.T) {
	suite.Run(t, &StorerSuite{store: NewMapStore()})
}

// Run tests for Storer using SyncStore implementation of Storer interface
func TestSyncStore(t *testing.T) {
	suite.Run(t, &StorerSuite{store: NewSyncStore()})
}
