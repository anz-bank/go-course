package main

import (
	"testing"

	tassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var (
	puppy1 = func() Puppy {
		return Puppy{ID: 11, Breed: "Chihuahua", Colour: "Brown", Value: 12.30}
	}
	puppy2 = func() Puppy {
		return Puppy{ID: 12, Breed: "Cacri", Colour: "Undefined", Value: 1.30}
	}
	puppy3 = func() Puppy {
		return Puppy{ID: 12, Breed: "Imaginary", Colour: "Undefined", Value: -1.30}
	}
)

type storerSuite struct {
	suite.Suite
	store      Storer
	storerType func() Storer
}

func (s *storerSuite) TestUpdatePuppyIDDoesNotExist() {
	// given
	assert := tassert.New(s.T())
	testPuppy := puppy1()

	// when
	err := s.store.UpdatePuppy(13, &testPuppy)

	// then
	assert.Error(err, "Should produce an error if id not found")
	serr, ok := err.(*Error)
	assert.True(ok)
	assert.Equal(uint16(0x2), serr.Code)
}

func (s *storerSuite) TestUpdatePuppy() {
	// given
	assert := tassert.New(s.T())
	testPuppy := puppy1()
	targetPuppy := puppy2()
	cerr := s.store.CreatePuppy(&testPuppy)
	r := require.New(s.T())
	r.NoError(cerr, "Create should not produce an error")

	// when
	uerr := s.store.UpdatePuppy(11, &targetPuppy)

	// then
	r.NoError(uerr, "Should be able to update store")
	updatedPuppy, err := s.store.ReadPuppy(11)
	r.NoError(err, "Should be able to read updated puppy")
	assert.Equal(targetPuppy, updatedPuppy, "Updated puppy should be equal to puppy2")
}

func (s *storerSuite) TestReadPuppy() {
	// given
	assert := tassert.New(s.T())
	testPuppy := puppy1()
	cerr := s.store.CreatePuppy(&testPuppy)
	r := require.New(s.T())
	r.NoError(cerr, "Create should not produce an error")

	// when
	newPuppy, err := s.store.ReadPuppy(11)

	// then
	r.NoError(err, "Should be able to read a newly added puppy")
	assert.Equal(testPuppy, newPuppy, "Newly added puppy should be equal to test puppy")
}

func (s *storerSuite) TestReadNonExistingPuppy() {
	// given
	assert := tassert.New(s.T())
	r := require.New(s.T())

	// when
	_, err := s.store.ReadPuppy(12)

	// then
	r.Error(err, "Should produce an error when puppy is not found")
	serr, ok := err.(*Error)
	assert.True(ok)
	assert.Equal(uint16(0x2), serr.Code)
}

func (s *storerSuite) TestDeletePuppy() {
	// given
	assert := tassert.New(s.T())
	testPuppy := puppy1()
	cerr := s.store.CreatePuppy(&testPuppy)
	r := require.New(s.T())
	r.NoError(cerr, "Create should not produce an error")

	// when
	err := s.store.DeletePuppy(11)

	// then
	r.NoError(err, "Should be able to delete a newly added puppy")
	_, rerr := s.store.ReadPuppy(11)
	serr, ok := rerr.(*Error)
	assert.True(ok)
	assert.Equal(uint16(0x2), serr.Code)
}

func (s *storerSuite) TestDeleteNonExistingPuppy() {
	// given
	assert := tassert.New(s.T())
	r := require.New(s.T())

	// when
	err := s.store.DeletePuppy(11)

	// then
	r.Error(err, "Should not be able to delete a non existing puppy")
	serr, ok := err.(*Error)
	assert.True(ok)
	assert.Equal(uint16(0x2), serr.Code)
}

func (s *storerSuite) TestCreateExistingPuppy() {
	// given
	assert := tassert.New(s.T())
	testPuppy := puppy1()
	cerr := s.store.CreatePuppy(&testPuppy)
	r := require.New(s.T())
	r.NoError(cerr, "Create should not produce an error")

	// when
	err := s.store.CreatePuppy(&testPuppy)

	// then
	assert.Error(err, "Should not be able to create twice a the same puppy")
}

func (s *storerSuite) TestCreatePuppyWithInvalidValue() {
	// given
	assert := tassert.New(s.T())
	testPuppy := puppy3()
	r := require.New(s.T())

	// when
	createError := s.store.CreatePuppy(&testPuppy)

	// then
	r.Error(createError, "Should not allow to create a puppy with invalid value")
	serr, ok := createError.(*Error)
	assert.True(ok)
	assert.Equal(uint16(0x1), serr.Code)
}

func (s *storerSuite) TestUpdatePuppyWithInvalidValue() {
	// given
	assert := tassert.New(s.T())
	testPuppy := puppy1()
	updatePuppy := puppy3()
	createError := s.store.CreatePuppy(&testPuppy)
	r := require.New(s.T())
	r.NoError(createError, "Create should not produce an error")

	// when
	uerr := s.store.UpdatePuppy(11, &updatePuppy)

	// then
	r.Error(uerr, "Should not allow to update a puppy with invalid value")
	serr, ok := uerr.(*Error)
	assert.True(ok)
	assert.Equal(uint16(0x1), serr.Code)
}

func (s *storerSuite) SetupTest() {
	s.store = s.storerType()
}

func TestStorer(t *testing.T) {
	syncSuite := storerSuite{
		store:      NewSyncStore(),
		storerType: func() Storer { return NewSyncStore() },
	}
	mapSuite := storerSuite{
		store:      NewMapStore(),
		storerType: func() Storer { return NewMapStore() },
	}
	suite.Run(t, &syncSuite)
	suite.Run(t, &mapSuite)
}
