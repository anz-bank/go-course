package main

import (
	"testing"

	tassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type storerImpl int

const (
	syncImpl storerImpl = iota
	mapImpl
)

const firstPuppyID uint32 = 1

var (
	firstPuppy = func() Puppy {
		return Puppy{
			Breed:  "Retriever",
			Colour: "Golden",
			Value:  9999.99,
		}
	}
	modifiedPuppy = func() Puppy {
		return Puppy{
			Breed:  "Bitsa",
			Colour: "Mixed",
			Value:  1.99,
		}
	}
	anotherPuppy = func() Puppy {
		return Puppy{
			Breed:  "Labrador",
			Colour: "Black",
			Value:  999.99,
		}
	}
)

type storerSuite struct {
	suite.Suite
	store Storer
	impl  storerImpl
}

func (s *storerSuite) SetupTest() {
	var err error
	switch s.impl {
	case syncImpl:
		// create a new syncstore and add the first puppy
		s.store = NewSyncStore()
		puppySyncImpl := firstPuppy()
		_, err = s.store.CreatePuppy(&puppySyncImpl)
	case mapImpl:
		// create a new mapstore and add the first puppy
		s.store = NewMapStore()
		puppyMapImpl := firstPuppy()
		_, err = s.store.CreatePuppy(&puppyMapImpl)
	default:
		panic("Unrecognised storer implementation")
	}
	if err != nil {
		panic("Failed to setup puppy test")
	}
}

func (s *storerSuite) TestCreate() {
	// given
	assert := tassert.New(s.T())
	newPuppy := anotherPuppy()

	// when
	createdPuppyID, err := s.store.CreatePuppy(&newPuppy)
	newPuppy.ID = createdPuppyID

	// then
	assert.NoError(err, "Should not get an error creating a puppy")
	foundPuppy, err := s.store.ReadPuppy(createdPuppyID)
	if assert.NoError(err, "Should be able to read an newly created puppy") {
		assert.Equal(&newPuppy, foundPuppy, "Created puppy should be identical to the one passed to create")
	}
}

func (s *storerSuite) TestRead() {
	// given
	assert := tassert.New(s.T())
	expectedPuppy := firstPuppy()
	expectedPuppy.ID = firstPuppyID

	// when
	actualPuppy, err := s.store.ReadPuppy(firstPuppyID)

	// then
	if assert.NoError(err, "Should not get an error reading puppy from store") {
		assert.Equal(&expectedPuppy, actualPuppy, "Store should return a puppy identical to firstPuppy")
	}
}

func (s *storerSuite) TestReadFail() {
	// given
	assert := tassert.New(s.T())

	// when
	_, err := s.store.ReadPuppy(99)

	// then
	assert.Error(err, "Should get an error when attempting to read a non-existent puppy")
}

func (s *storerSuite) TestUpdate() {
	// given
	assert := tassert.New(s.T())
	updatePuppy := modifiedPuppy()

	// when
	puppyID, err := s.store.UpdatePuppy(firstPuppyID, &updatePuppy)

	// then
	if assert.NoError(err, "Should not get an error updating a puppy") {
		foundPuppy, err := s.store.ReadPuppy(puppyID)
		if assert.NoError(err, "Should not get an error reading the updated puppy") {
			assert.Equal(&updatePuppy, foundPuppy, "Found puppy should be equal to updated puppy")
		}
	}
}

func (s *storerSuite) TestUpdateFail() {
	// given
	assert := tassert.New(s.T())
	updatePuppy := anotherPuppy()

	// when
	puppyID, err := s.store.UpdatePuppy(99, &updatePuppy)

	// then
	if assert.NoError(err, "Should not get an error updating a puppy") {
		foundPuppy, err := s.store.ReadPuppy(puppyID)
		if assert.NoError(err, "Should not get an error reading the updated puppy") {
			assert.Equal(&updatePuppy, foundPuppy, "Found puppy should be equal to updated puppy")
		}
	}
}

func (s *storerSuite) TestDeleteExisting() {
	// when
	assert := tassert.New(s.T())
	err := s.store.DeletePuppy(firstPuppyID)

	// then
	assert.NoError(err, "Should not get an error deleting an existing puppy")
	_, err = s.store.ReadPuppy(firstPuppyID)
	assert.Error(err, "Should not be able to read a deleted puppy")
}

func (s *storerSuite) TestDeleteNotExisting() {
	// when
	assert := tassert.New(s.T())
	err := s.store.DeletePuppy(99)

	// then
	assert.NoError(err, "Should not get an error deleting a non-existent puppy")
}

func TestStorer(t *testing.T) {
	syncSuite := storerSuite{
		store: NewSyncStore(),
		impl:  syncImpl,
	}
	suite.Run(t, &syncSuite)
	mapSuite := storerSuite{
		store: NewMapStore(),
		impl:  mapImpl,
	}
	suite.Run(t, &mapSuite)
}
