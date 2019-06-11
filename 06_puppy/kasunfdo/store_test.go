package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StorerTest struct {
	suite.Suite
	store Storer
	id    uint64
}

func (suite *StorerTest) SetupTest() {
	suite.id = suite.store.CreatePuppy(Puppy{Breed: "Labrador", Colour: "Cream", Value: 2999.99})
}

func (suite *StorerTest) TestCreatePuppy() {
	id := suite.store.CreatePuppy(Puppy{Breed: "German Shepard", Colour: "Brown", Value: 3499.99})
	suite.True(id > 1)
}

func (suite *StorerTest) TestReadPuppy() {
	puppy, err := suite.store.ReadPuppy(suite.id)

	suite.Nil(err)
	suite.Equal(puppy.ID, suite.id)
	suite.Equal(puppy.Breed, "Labrador")
	suite.Equal(puppy.Colour, "Cream")
	suite.Equal(puppy.Value, 2999.99)

	_, err = suite.store.ReadPuppy(100)
	suite.Equal("no puppy with id: 100", err.Error())
}

func (suite *StorerTest) TestUpdatePuppy() {
	err := suite.store.UpdatePuppy(Puppy{ID: suite.id, Breed: "Labrador Retriever", Colour: "Brown", Value: 3999.99})
	suite.Nil(err)

	puppy, err := suite.store.ReadPuppy(suite.id)

	suite.Nil(err)
	suite.Equal(puppy.ID, suite.id)
	suite.Equal(puppy.Breed, "Labrador Retriever")
	suite.Equal(puppy.Colour, "Brown")
	suite.Equal(puppy.Value, 3999.99)

	err = suite.store.UpdatePuppy(Puppy{ID: 100, Breed: "Poodle", Colour: "White", Value: 1999.99})
	suite.Equal("no puppy with id: 100", err.Error())
}

func (suite *StorerTest) TestDeletePuppy() {
	err := suite.store.DeletePuppy(suite.id)
	suite.Nil(err)

	_, err = suite.store.ReadPuppy(suite.id)
	suite.Equal(fmt.Sprintf("no puppy with id: %d", suite.id), err.Error())

	err = suite.store.DeletePuppy(suite.id)
	suite.Equal(fmt.Sprintf("no puppy with id: %d", suite.id), err.Error())
}

func TestMapStore(t *testing.T) {
	suite.Run(t, &StorerTest{store: NewMapStore()})
}

func TestSyncStore(t *testing.T) {
	suite.Run(t, &StorerTest{store: NewSyncStore()})
}
