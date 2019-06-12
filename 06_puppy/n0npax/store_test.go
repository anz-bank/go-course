package main

import (
	"testing"

	tassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type storerImpl int

const (
	syncStorer storerImpl = iota
	memStorer  storerImpl = iota
)

var (
	puppy0 = func() Puppy {
		return Puppy{
			Breed:  "Norwegian Forest cat",
			Colour: "Grey",
			Value:  "$300",
		}
	}
	puppy1 = func() Puppy {
		return Puppy{
			Breed:  "Maine Coon",
			Colour: "Brown",
			Value:  "€400",
		}
	}
)

type storerSuite struct {
	suite.Suite
	store Storer
	impl  storerImpl
}

func (s *storerSuite) SetupTest() {
	switch s.impl {
	case syncStorer:
		s.store = NewSyncStore()
		puppy := puppy0()
		s.store.CreatePuppy(&puppy)
	case memStorer:
		s.store = NewMemStore()
		puppy := puppy0()
		s.store.CreatePuppy(&puppy)
	default:
		panic("Unrecognised storer implementation")
	}
}

func (s *storerSuite) TestCreatePuppySuccessful() {
	assert := tassert.New(s.T())
	newPuppy0, newPuppy1 := puppy0(), puppy1()
	id0, id1 := s.store.CreatePuppy(&newPuppy0), s.store.CreatePuppy(&newPuppy1)
	assert.Equal(id0, id1-1, "2nd id should be 1st +1, got", id0, id1)
}

func (s *storerSuite) TestReadPuppySuccessful() {
	assert := tassert.New(s.T())
	expectedPuppy := puppy0()
	puppy, err := s.store.ReadPuppy(expectedPuppy.ID)

	if assert.NoError(err, "Should be able to read puppy0 from store") {
		assert.Equal(&expectedPuppy, puppy, "store should return a puppy identical to puppy0")
	}
}

func (s *storerSuite) TestReadPuppyIDDoesNotExist() {
	assert := tassert.New(s.T())
	_, err := s.store.ReadPuppy(1000)
	assert.Error(err, "Should get an error when attempting to read an non-existing puppy")
}

func (s *storerSuite) TestUpdatePuppy() {
	assert := tassert.New(s.T())
	existingPuppy := puppy0()
	existingPuppy.Colour = "Black"
	err := s.store.UpdatePuppy(existingPuppy.ID, &existingPuppy)

	assert.NoError(err, "Update should successfully update a puppy")
	puppy, err := s.store.ReadPuppy(existingPuppy.ID)
	assert.NoError(err, "Read of updated puppy should be successful")
	assert.Equal(existingPuppy.Colour, puppy.Colour, "Updated colour missmatch")
}

func (s *storerSuite) TestDeletePuppySuccessful() {
	assert := tassert.New(s.T())
	existingPuppy := puppy0()
	deleted, err := s.store.DeletePuppy(existingPuppy.ID)

	assert.NoError(err, "Delete should successfully delete a puppy")
	assert.True(deleted, "Delete should return true indicating a puppy was deleted")
	_, err = s.store.ReadPuppy(existingPuppy.ID)
	assert.Error(err, "Should not be able to read a deleted ID")
}

func (s *storerSuite) TestDeletePuppy_IDDoesNotExist() {
	assert := tassert.New(s.T())
	_, err := s.store.DeletePuppy(1000)
	assert.Error(err, "Should not be able to delete puppy with non existing ID")
}

func TestStorerSyncMap(t *testing.T) {
	syncSuite := storerSuite{
		store: NewSyncStore(),
		impl:  syncStorer,
	}
	suite.Run(t, &syncSuite)
}

func TestStorerMemoryMap(t *testing.T) {
	syncSuite := storerSuite{
		store: NewMemStore(),
		impl:  memStorer,
	}
	suite.Run(t, &syncSuite)
}
