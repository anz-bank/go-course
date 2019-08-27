package main

import (
	"testing"

	tassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	puppy0 = func() Puppy {
		return Puppy{
			Breed:  "Type A",
			Colour: "Grey",
			Value:  300,
		}
	}
	puppy1 = func() Puppy {
		return Puppy{
			Breed:  "Type B",
			Colour: "Brown",
			Value:  400,
		}
	}
	puppyNegativeValue = func() Puppy {
		return Puppy{
			Breed:  "Type C",
			Colour: "Red",
			Value:  -1,
		}
	}
	puppyNegativeID = func() Puppy {
		return Puppy{
			Breed:  "Type D",
			Colour: "Blue",
			Value:  100,
			ID:     -1,
		}
	}
)

type storerSuite struct {
	suite.Suite
	store Storer
}

func (s *storerSuite) SetupTest() {
	switch s.store.(type) {
	case *SyncStore:
		s.store = NewSyncStore()
	case *MemStore:
		s.store = NewMemStore()
	default:
		panic("Unrecognised storer implementation")
	}
	puppy := puppy0()
	_, err := s.store.CreatePuppy(&puppy)
	if err != nil {
		panic(err)
	}
}

func TestStorer(t *testing.T) {
	suite.Run(t, &storerSuite{store: NewMemStore()})
	suite.Run(t, &storerSuite{store: NewSyncStore()})
}

func (s *storerSuite) TestCreatePuppySuccessful() {
	assert := tassert.New(s.T())
	newPuppy0, newPuppy1 := puppy0(), puppy1()
	id0, err := s.store.CreatePuppy(&newPuppy0)
	assert.NoError(err, "Creating puppy should be ok")
	id1, err := s.store.CreatePuppy(&newPuppy1)
	assert.NoError(err, "Creating puppy should be ok")
	assert.Equal(id0, id1-1, "2nd id should be 1st +1, got %v and %v", id0, id1)
}

func (s *storerSuite) TestCreatePuppyNegativeValue() {
	assert := tassert.New(s.T())
	newPuppy := puppyNegativeValue()
	_, err := s.store.CreatePuppy(&newPuppy)
	assert.Error(err, "negative ID should cause an error")
}

func (s *storerSuite) TestReadPuppySuccessful() {
	for i := 0; i < 5; i++ {
		assert := tassert.New(s.T())
		newPuppy := puppy0()
		id, err := s.store.CreatePuppy(&newPuppy)
		assert.NoError(err, "Creating puppy should be ok")

		readPuppy, err := s.store.ReadPuppy(id)
		if assert.NoError(err, "Should be able to read puppy0 from store") {
			assert.Equal(&newPuppy, readPuppy, "store should return identic puppy")
		}
		assert.Equal(id, readPuppy.ID)
	}
}

func (s *storerSuite) TestCreateDeleteCreate() {
	assert := tassert.New(s.T())
	newPuppy0, newPuppy1 := puppy0(), puppy1()
	_, err := s.store.CreatePuppy(&newPuppy0)
	assert.NoError(err)
	_, err = s.store.CreatePuppy(&newPuppy1)
	assert.NoError(err)
	_, err = s.store.DeletePuppy(newPuppy0.ID)
	assert.NoError(err)
	newerPuppy0 := puppy0()
	_, err = s.store.CreatePuppy(&newerPuppy0)
	assert.NoError(err)
	p, _ := s.store.ReadPuppy(newPuppy1.ID)
	assert.Equal(newPuppy1, *p)
}

func (s *storerSuite) TestReadPuppyIDDoesNotExist() {
	assert := tassert.New(s.T())
	_, err := s.store.ReadPuppy(1000)
	assert.Error(err, "Should get an error when attempting to read an non-existing puppy")
}

func (s *storerSuite) TestReadPuppyNegativeID() {
	assert := tassert.New(s.T())
	_, err := s.store.ReadPuppy(-1)
	assert.Error(err, "negative ID should cause an error")
}

func (s *storerSuite) TestUpdatePuppy() {
	assert := tassert.New(s.T())
	existingPuppy, err := s.store.ReadPuppy(0)
	assert.NoError(err, "Reading puppy should not return error")
	existingPuppy.Colour = "Purple"
	err = s.store.UpdatePuppy(0, existingPuppy)
	assert.NoError(err, "Update should not return any error")
	puppy, err := s.store.ReadPuppy(0)
	assert.NoError(err, "Reading puppy should not return error")
	assert.Equal(existingPuppy.Colour, puppy.Colour, "Updated colour missmatch")
}

func (s *storerSuite) TestUpdatePuppyCorruptedID() {
	assert := tassert.New(s.T())
	existingPuppy, err := s.store.ReadPuppy(0)
	assert.NoError(err, "Reading puppy should be ok")
	err = s.store.UpdatePuppy(1000, existingPuppy)
	assert.Error(err, "Should get an error when attempting to update with corrupted id")
}

func (s *storerSuite) TestUpdatePuppyIDDoesNotExist() {
	assert := tassert.New(s.T())
	newPuppy := puppy0()
	err := s.store.UpdatePuppy(1000, &newPuppy)
	assert.Error(err, "Should get an error when attempting to update an non-existing puppy")
}

func (s *storerSuite) TestUpdatePuppyNegativeID() {
	assert := tassert.New(s.T())
	newPuppy := puppyNegativeID()
	err := s.store.UpdatePuppy(newPuppy.ID, &newPuppy)
	assert.Error(err, "negative ID should cause an error")
}

func (s *storerSuite) TestUpdatePuppyNegativeValue() {
	assert := tassert.New(s.T())
	newPuppy := puppyNegativeValue()
	err := s.store.UpdatePuppy(newPuppy.ID, &newPuppy)
	assert.Error(err, "negative ID should cause an error")
}

func (s *storerSuite) TestDeletePuppySuccessful() {
	assert := tassert.New(s.T())
	existingPuppy := puppy0()
	deleted, err := s.store.DeletePuppy(0)
	assert.NoError(err, "Delete should successfully delete a puppy")
	assert.True(deleted, "Delete should return true indicating a puppy was deleted")
	_, err = s.store.ReadPuppy(existingPuppy.ID)
	assert.Error(err, "Should not be able to read a deleted ID")
}

func (s *storerSuite) TestDeletePuppyIDDoesNotExist() {
	assert := tassert.New(s.T())
	_, err := s.store.DeletePuppy(1000)
	assert.Error(err, "Should not be able to delete puppy with non existing ID")
}

func (s *storerSuite) TestDeletePuppyNegativeID() {
	assert := tassert.New(s.T())
	_, err := s.store.DeletePuppy(-1)
	assert.Error(err, "negative ID should cause an error")
}
