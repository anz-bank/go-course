package main

import (
	"testing"

	tassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	puppy1 = func() Puppy {
		return Puppy{ID: 11, Breed: "Chihuahua", Colour: "Brown", Value: 12.30}
	}
)

type storerSuite struct {
	suite.Suite
	store Storer
}

func (s *storerSuite) TestUpdatePuppyIDDoesNotExist() {
	// given
	assert := tassert.New(s.T())
	testPuppy := puppy1()

	// when
	err := s.store.UpdatePuppy(11, &testPuppy)

	// then
	assert.NoError(err, "Updating to non-existing puppy ID is not an error")
	newPuppy, err := s.store.ReadPuppy(11)
	if assert.NoError(err, "Should be able to read a newly added puppy via update") {
		assert.Equal(&testPuppy, newPuppy, "Newly added puppy should be equal to test puppy")
	}
}

func (s *storerSuite) TestReadPuppy() {
	// given
	assert := tassert.New(s.T())
	testPuppy := puppy1()
	createError := s.store.CreatePuppy(&testPuppy)
	s.T().Log(createError)

	// when
	newPuppy, err := s.store.ReadPuppy(11)

	// then
	if assert.NoError(err, "Should be able to read a newly added puppy") {
		assert.Equal(&testPuppy, newPuppy, "Newly added puppy should be equal to test puppy")
	}
}

func (s *storerSuite) TestDeletePuppy() {
	// given
	assert := tassert.New(s.T())
	testPuppy := puppy1()
	createError := s.store.CreatePuppy(&testPuppy)
	s.T().Log(createError)

	// when
	isDeleted, err := s.store.DeletePuppy(11)

	// then
	if assert.NoError(err, "Should be able to delete a newly added puppy") {
		assert.True(isDeleted)
		_, err := s.store.ReadPuppy(11)
		assert.Error(err, "Puppy not found")
	}
}

func (s *storerSuite) TestDeleteNonExistingPuppy() {
	// given
	assert := tassert.New(s.T())

	// when
	isDeleted, err := s.store.DeletePuppy(11)

	// then
	if assert.Error(err, "Should not be able to delete a non existing puppy") {
		assert.False(isDeleted)
	}
}

func (s *storerSuite) TestCreateExistingPuppy() {
	// given
	assert := tassert.New(s.T())
	testPuppy := puppy1()
	createError := s.store.CreatePuppy(&testPuppy)
	s.T().Log(createError)

	// when
	err := s.store.CreatePuppy(&testPuppy)

	// then
	assert.Error(err, "Should not be able to create twice a the same puppy")
	// cleanup
	_, deleteError := s.store.DeletePuppy(11)
	s.T().Log(deleteError)
}

func TestStorer(t *testing.T) {
	syncSuite := storerSuite{
		store: NewSyncStore(),
	}
	mapSuite := storerSuite{
		store: NewMapStore(),
	}
	suite.Run(t, &syncSuite)
	suite.Run(t, &mapSuite)
}
