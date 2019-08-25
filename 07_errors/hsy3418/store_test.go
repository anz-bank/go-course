package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func (suite *storerTestSuite) SetupTest() {
	suite.store = suite.makeStore()
	suite.toBeCreatedPuppy = Puppy{ID: 101, Breed: "Poodle", Colour: "White", Value: 1000.5}
	suite.existsPuppy = Puppy{ID: 102, Breed: "Poodle", Colour: "White", Value: 1280.5}
	suite.toBeUpdatedPuppy = Puppy{ID: 102, Breed: "Poodle", Colour: "White", Value: 2000}
	suite.invalidPuppy = Puppy{ID: 103, Breed: "Poodle", Colour: "White", Value: -1000}
	suite.invaildError = ErrorEf(ErrInvalidInput, "The puppy value is invalidate")
	suite.invalidIDError = ErrorEf(ErrNotFound, "This puppy does not exist")
	suite.dupicatesError = ErrorEf(ErrDuplicate, "This puppy exists ")
	err := suite.store.CreatePuppy(suite.existsPuppy)
	if err != nil {
		suite.FailNow("Failed to setup test")
	}
}

type storerTestSuite struct {
	suite.Suite
	store            Storer
	makeStore        func() Storer
	toBeCreatedPuppy Puppy
	existsPuppy      Puppy
	toBeUpdatedPuppy Puppy
	invalidPuppy     Puppy
	invaildError     error
	invalidIDError   error
	dupicatesError   error
}

//successfully create puppy, new puppy, existing puppy, value <0
func (suite *storerTestSuite) TestCreatePuppy() {
	assert := assert.New(suite.T())
	testCases := []struct {
		title    string
		input    Puppy
		expected error
	}{
		{"Create new puppy", suite.toBeCreatedPuppy, nil},
		{"Create existing puppy", suite.toBeCreatedPuppy, suite.dupicatesError},
		{"Create a invalid puppy", suite.invalidPuppy, suite.invaildError},
	}
	for _, tc := range testCases {
		tc := tc
		suite.T().Run(tc.title, func(t *testing.T) {
			err := suite.store.CreatePuppy(tc.input)
			assert.Equal(tc.expected, err)
		})
	}

}

//
func (suite *storerTestSuite) TestUpdatePuppy() {
	assert := assert.New(suite.T())
	testCases := []struct {
		title         string
		inputPuppy    Puppy
		expectedError error
	}{
		{"Update puppy successfully", suite.toBeUpdatedPuppy, nil},
		{"Update a invalid puppy", suite.invalidPuppy, suite.invaildError},
		{"Update non-existing puppy", Puppy{}, suite.invalidIDError},
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
		{"Read non-existing puppy", suite.toBeCreatedPuppy.ID, Puppy{}, suite.invalidIDError},
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
		{"Delete non-existing puppy", suite.toBeCreatedPuppy.ID, suite.invalidIDError},
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
