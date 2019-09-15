package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type storesSuite struct {
	suite.Suite
	store     Storer
	newStorer func() Storer
	puppy1    Puppy
}

func (suite *storesSuite) SetupTest() {
	suite.store = suite.newStorer()
	suite.puppy1 = Puppy{Breed: "cavalier", Color: "White", Value: 1700}
	_, err := suite.store.CreatePuppy(&suite.puppy1)
	if err != nil {
		suite.FailNow("Test setup failed to create puppy", err)
	}

}
func (suite *storesSuite) TestCreatePuppySuccesful() {
	//given
	assert := assert.New(suite.T())
	newPuppy := Puppy{Breed: "Dobberman", Color: "Black", Value: 800}
	//when
	id, err := suite.store.CreatePuppy(&newPuppy)
	//then
	assert.NoError(err, "Puppy creation should be succesful")
	assert.True(newPuppy.ID == suite.puppy1.ID+1, "Puppy Id increments sequentially")
	puppy, err := suite.store.ReadPuppy(id)
	if assert.NoError(err, "Read of created puppy should be successful") {
		assert.Equal(puppy, &newPuppy)
	}
}
func (suite *storesSuite) TestCreatePuppyForNonZeroValue() {
	//given
	assert := assert.New(suite.T())
	newPuppy := Puppy{Breed: "New", Color: "Black", Value: -1}
	//when
	_, err := suite.store.CreatePuppy(&newPuppy)
	//then
	assert.Error(err, "Negative Value should throw an error")
}
func (suite *storesSuite) TestReadPuppySuccesful() {
	//given
	assert := assert.New(suite.T())
	//when
	puppy, err := suite.store.ReadPuppy(suite.puppy1.ID)
	//then
	if assert.NoError(err, "Read of puppy should be successful") {
		assert.Equal(puppy, &suite.puppy1)
	}
}
func (suite *storesSuite) TestReadPuppyDoesNotExist() {
	assert := assert.New(suite.T())
	_, err := suite.store.ReadPuppy(100)
	assert.Error(err, "Read of non existing puppy Id should throw an error")
}
func (suite *storesSuite) TestUpdatePuppySuccesful() {
	//given
	assert := assert.New(suite.T())
	existingPuppy := suite.puppy1
	//when
	existingPuppy.Color = "Brown"
	existingPuppy.Breed = "Dobberman"
	err := suite.store.UpdatePuppy(&existingPuppy)
	//then
	assert.NoError(err)
	updatedPuppy, err := suite.store.ReadPuppy(existingPuppy.ID)
	assert.NoError(err)
	assert.NotEqual(updatedPuppy, &suite.puppy1, "Puppy values are updated")
}
func (suite *storesSuite) TestUpdatePuppyIdDoesNotExist() {
	assert := assert.New(suite.T())
	puppy := Puppy{ID: 100, Breed: "cavalier", Color: "White", Value: 1700}
	err := suite.store.UpdatePuppy(&puppy)
	assert.Error(err, "Update of non existing puppy Id should throw an error")
}
func (suite *storesSuite) TestDeletePuppySuccesful() {
	assert := assert.New(suite.T())
	//when
	err := suite.store.DeletePuppy(suite.puppy1.ID)
	//then
	assert.NoError(err, "Delete should successfully delete existing puppy")
	_, err = suite.store.ReadPuppy(suite.puppy1.ID)
	assert.Error(err, "Read of non existing puppy Id should throw an error")
}
func (suite *storesSuite) TestDeletePuppyFailure() {
	assert := assert.New(suite.T())
	err := suite.store.DeletePuppy(12)
	assert.Error(err, "Delete of non existing puppy Id should throw an error ")
}
func TestSuite(t *testing.T) {
	suite.Run(t, &storesSuite{newStorer: func() Storer { return NewMapStore() }})
	suite.Run(t, &storesSuite{newStorer: func() Storer { return NewSyncStore() }})
}
