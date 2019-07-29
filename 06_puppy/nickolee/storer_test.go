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
}

func (suite *StorerSuite) TestCreateAndRetrievePuppy() {
	puppyID := suite.store.CreatePuppy(puppies[3])
	puppy, _ := suite.store.ReadPuppy(puppyID)
	suite.Equal(puppy, puppies[3])
}

func (suite *StorerSuite) TestUpdatePuppy() {
	id := suite.store.CreatePuppy(puppies[0])
	// update the puppy we just created with another puppy
	err := suite.store.UpdatePuppy(id, puppies[3])
	suite.Nil(err) // if no error it means puppy was updated successfully

	// retrieve to double check update was correct in the following assertion on next line
	retrievedPuppy, _ := suite.store.ReadPuppy(id)
	suite.Equal(puppies[3], retrievedPuppy)
}

func (suite *StorerSuite) TestUpdateNonExistentPuppy() {
	err := suite.store.UpdatePuppy(7, puppies[3]) // trying to update a non-existent puppy
	suite.Error(err)                              // assert there should be an error cos puppy with id 7 doesn't exist
}

func (suite *StorerSuite) TestDeletePuppy() {
	id := suite.store.CreatePuppy(puppies[1])
	// delete the puppy we just created
	err := suite.store.DeletePuppy(id)
	// check that puppy is successfully deleted
	suite.Nil(err)

	// verify delete was successful by trying to retrieve puppy
	retrievedPuppy, err := suite.store.ReadPuppy(id)
	suite.Nil(retrievedPuppy) // should get back a nil pointer since there ain't no more puppy
	suite.Error(err)          // asserting there should be an error since you cannot retrieve a deleted puppy
}

func (suite *StorerSuite) TestDeleteNonExistentPuppy() {
	err := suite.store.DeletePuppy(7) // trying to delete non-existent puppy
	suite.Error(err)                  // assert there should be an error cos puppy with id 7 doesn't exist
}

// Run tests for Storer using MapStore implementation of Storer interface
func TestMapStore(t *testing.T) {
	suite.Run(t, &StorerSuite{store: NewMapStore()})
}

// Run tests for Storer using SyncStore implementation of Storer interface
func TestSyncStore(t *testing.T) {
	suite.Run(t, &StorerSuite{store: NewSyncStore()})
}
