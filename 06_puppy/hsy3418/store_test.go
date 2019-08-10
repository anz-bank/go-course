package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func (suite *storerTestSuite) SetupTest() {
	suite.store = suite.makeStore()
	suite.nonExistPuppy = Puppy{ID: 123456, Breed: "Poodle", Colour: "White", Value: 1000.5}
	suite.toBeCreatedPuppy = Puppy{Breed: "Poodle", Colour: "White", Value: 1000.5}
	suite.existsPuppy = Puppy{Breed: "Poodle", Colour: "White", Value: 1280.5}
	suite.toBeUpdatedPuppy = Puppy{ID: suite.existsPuppy.ID, Breed: "Poodle", Colour: "White", Value: 2000}
	suite.store.CreatePuppy(suite.existsPuppy)
}

type storerTestSuite struct {
	suite.Suite
	store            Storer
	makeStore        func() Storer
	toBeCreatedPuppy Puppy
	existsPuppy      Puppy
	toBeUpdatedPuppy Puppy
	nonExistPuppy    Puppy
}

func (suite *storerTestSuite) TestCreatePuppy() {
	assert := assert.New(suite.T())
	testCases := []struct {
		title string
		input Puppy
	}{
		{"Create new puppy", suite.toBeCreatedPuppy},
	}
	for _, tc := range testCases {
		tc := tc
		suite.T().Run(tc.title, func(t *testing.T) {
			id := suite.store.CreatePuppy(tc.input)
			assert.True(id > 0)
		})
	}
}

func (suite *storerTestSuite) TestUpdatePuppy() {
	assert := assert.New(suite.T())
	testCases := []struct {
		title         string
		inputPuppy    Puppy
		expectedError error
	}{
		{"Update puppy successfully", suite.toBeUpdatedPuppy, nil},
		{"Update non-existing puppy", suite.nonExistPuppy,
			fmt.Errorf("puppy with %d ID does not exist", suite.nonExistPuppy.ID)},
	}
	for _, tc := range testCases {
		tc := tc
		suite.T().Run(tc.title, func(t *testing.T) {
			err := suite.store.UpdatePuppy(tc.inputPuppy)
			assert.Equal(tc.expectedError, err)
		})
	}
}

func (suite *storerTestSuite) TestReadPuppy() {
	assert := assert.New(suite.T())
	testCases := []struct {
		title         string
		input         int32
		expected      Puppy
		expectedError error
	}{
		{"Read puppy successfully", suite.existsPuppy.ID, suite.existsPuppy, nil},
		{"Read non-existing puppy", suite.nonExistPuppy.ID, Puppy{},
			fmt.Errorf("puppy with %d ID does not exist", suite.nonExistPuppy.ID)},
	}
	for _, tc := range testCases {
		tc := tc
		suite.T().Run(tc.title, func(t *testing.T) {
			readPuppy, err := suite.store.ReadPuppy(tc.input)
			assert.Equal(tc.expected, readPuppy)
			assert.Equal(tc.expectedError, err)
		})
	}
}

func (suite *storerTestSuite) TestDeletePuppy() {
	assert := assert.New(suite.T())
	testCases := []struct {
		title         string
		input         int32
		expectedError error
	}{
		{"Delete puppy successfully", suite.existsPuppy.ID, nil},
		{"Delete non-existing puppy", suite.nonExistPuppy.ID,
			fmt.Errorf("puppy with %d ID does not exist", suite.nonExistPuppy.ID)},
	}
	for _, tc := range testCases {
		tc := tc
		suite.T().Run(tc.title, func(t *testing.T) {
			err := suite.store.DeletePuppy(tc.input)
			assert.Equal(tc.expectedError, err)
		})
	}
}

func TestStorers(t *testing.T) {
	suite.Run(t, &storerTestSuite{
		makeStore: func() Storer { return NewMapStore() },
	})
	suite.Run(t, &storerTestSuite{
		makeStore: func() Storer { return NewSyncStore() },
	})
}
