package store

import (
	"fmt"
	"testing"

	"github.com/anz-bank/go-course/08_project/nickolee/pkg/puppy"

	"github.com/stretchr/testify/suite"
)

type StorerSuite struct {
	suite.Suite
	store puppy.Storer
}

// A bunch of seed puppies
var puppies = []*puppy.Puppy{
	{Breed: "Laika", Colour: "First dog on moon", Value: 2130.50},
	{Breed: "Pomsky", Colour: "Violet", Value: 777.77},
	{Breed: "Dalmatian", Colour: "Cruella de vil", Value: 101},
	{Breed: "Rare Lvl 100 Arcanine", Colour: "Brown", Value: 9999.99},
	{Breed: "Cheap Pug", Colour: "Cheap Brown", Value: -500.50},
}

func (suite *StorerSuite) TestCreateAndRetrievePuppy() {
	puppyID, err := suite.store.CreatePuppy(puppies[3])
	suite.Nil(err)
	pup, err := suite.store.ReadPuppy(puppyID)
	suite.Nil(err)
	suite.Equal(pup, puppies[3])

	// Test negative value puppy
	_, rawErr := suite.store.CreatePuppy(puppies[4])
	// Require will stop the test if error is wrong type (graceful handle), preventing a panic on the next line
	suite.Require().IsType(&puppy.Error{}, rawErr)
	actualErr, _ := rawErr.(*puppy.Error) // Type cast, err now holds the actual error
	suite.Equal(puppy.ErrNegativePuppyID, actualErr.Code)
	suite.Equal("Sorry puppy value cannot be negative. The dog has to be worth something :)", actualErr.Message)

	// test Error() method in error.go for 100% coverage
	suite.Equal(actualErr.Error(), fmt.Sprintf("PuppyStoreError %d: %s", actualErr.Code, actualErr.Message))
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

	// Test error path
	rawErr := suite.store.UpdatePuppy(1829246, puppies[1]) // give non-existent id
	suite.Require().IsType(&puppy.Error{}, rawErr)         // handle gracefully
	actualErr, _ := rawErr.(*puppy.Error)                  // Type cast, err now holds the actual error
	suite.Equal(puppy.ErrPuppyNotFound, actualErr.Code)
	suite.Equal("Sorry puppy with ID 1829246 does not exist", actualErr.Message)
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
	rawErr := suite.store.DeletePuppy(1829246)     // give non-existent id
	suite.Require().IsType(&puppy.Error{}, rawErr) // handle gracefully
	actualErr, _ := rawErr.(*puppy.Error)          // Type cast, err now holds the actual error
	suite.Equal(puppy.ErrPuppyNotFound, actualErr.Code)
	suite.Equal("Sorry puppy with ID 1829246 does not exist", actualErr.Message)
}

// Run tests for Storer using MapStore implementation of Storer interface
func TestMapStore(t *testing.T) {
	suite.Run(t, &StorerSuite{store: NewMapStore()})
}

// Run tests for Storer using SyncStore implementation of Storer interface
func TestSyncStore(t *testing.T) {
	suite.Run(t, &StorerSuite{store: NewSyncStore()})
}
