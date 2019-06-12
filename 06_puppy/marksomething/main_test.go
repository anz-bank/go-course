package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PuppyStorerTest struct {
	suite.Suite
	store        PuppyStorer
	storeFactory func() PuppyStorer
}

// Create a new instance of PuppyStorer for each test case
func (suite *PuppyStorerTest) SetupTest() {
	suite.store = suite.storeFactory()
}

func (suite *PuppyStorerTest) TestReadNonExistantPuppy() {
	expected := Puppy{}

	actual := suite.store.ReadPuppy(99)

	suite.Equal(expected, actual)
}

func (suite *PuppyStorerTest) TestUpdateNonExistantPuppy() {
	newPuppy := Puppy{1, "Lab", "Brown", 1}
	expected := Puppy{}

	suite.store.UpdatePuppy(1, newPuppy)
	actual := suite.store.ReadPuppy(1)

	suite.Equal(expected, actual)
}

func (suite *PuppyStorerTest) TestUpdatePuppyBadId() {
	expected := Puppy{1, "Lab", "Brown", 1}
	differentPuppy := Puppy{7, "Poodle", "Brown", 1}

	suite.store.CreatePuppy(expected)

	suite.store.UpdatePuppy(1, differentPuppy)
	actual := suite.store.ReadPuppy(1)
	suite.Equal(expected, actual)

	differentPuppyFromStorage := suite.store.ReadPuppy(7)
	blankPuppy := Puppy{}
	suite.Equal(blankPuppy, differentPuppyFromStorage)
}

func (suite *PuppyStorerTest) TestCreateReadPuppy() {
	expected := Puppy{1, "Lab", "Brown", 1}

	suite.store.CreatePuppy(expected)

	actual := suite.store.ReadPuppy(1)
	suite.Equal(expected, actual)
}

func (suite *PuppyStorerTest) TestDoubleCreatePuppy() {
	expected := Puppy{1, "Lab", "Brown", 1}

	suite.store.CreatePuppy(expected)
	suite.store.CreatePuppy(expected)

	actual := suite.store.ReadPuppy(1)
	suite.Equal(expected, actual)
}

func (suite *PuppyStorerTest) TestCreateDeleteReadPuppy() {
	expected := Puppy{}

	suite.store.CreatePuppy(Puppy{1, "Lab", "Brown", 1})
	suite.store.DeletePuppy(1)
	actual := suite.store.ReadPuppy(1)

	suite.Equal(expected, actual)
}

func (suite *PuppyStorerTest) TestDeleteNonExistantPuppy() {
	expected := suite.store
	suite.store.DeletePuppy(1)
	suite.Equal(expected, suite.store)
}

func (suite *PuppyStorerTest) TestCreateReadMultiplePuppy() {
	expected1 := Puppy{1, "Lab", "Brown", 1}
	expected2 := Puppy{2, "Poodle", "Blue", 2}

	suite.store.CreatePuppy(expected1)
	suite.store.CreatePuppy(expected2)
	actual1 := suite.store.ReadPuppy(1)
	actual2 := suite.store.ReadPuppy(2)

	suite.Equal(expected1, actual1)
	suite.Equal(expected2, actual2)
}

func (suite *PuppyStorerTest) TestCreateUpdateReadPuppy() {
	originalPuppy := Puppy{1, "Lab", "Brown", 32}
	modifiedPuppy := Puppy{1, "Lab", "Brown", 22}

	suite.store.CreatePuppy(originalPuppy)
	suite.store.UpdatePuppy(1, modifiedPuppy)

	actual := suite.store.ReadPuppy(1)
	suite.Equal(modifiedPuppy, actual)
}

func TestMapStore(t *testing.T) {
	s := PuppyStorerTest{
		storeFactory: func() PuppyStorer { return MapStore{} },
	}
	suite.Run(t, &s)
}

func TestSyncStore(t *testing.T) {
	s := PuppyStorerTest{
		storeFactory: func() PuppyStorer { return &SyncStore{} },
	}
	suite.Run(t, &s)
}
