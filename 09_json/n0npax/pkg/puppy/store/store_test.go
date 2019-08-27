package store

import (
	"math/rand"
	"testing"

	puppy "github.com/anz-bank/go-course/09_json/n0npax/pkg/puppy"
	tassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	puppy0 = func() puppy.Puppy {
		return puppy.Puppy{
			Breed:  "Type A",
			Colour: "Grey",
			Value:  rand.Intn(100) + 100,
		}
	}
	puppy1 = func() puppy.Puppy {
		return puppy.Puppy{
			Breed:  "Type B",
			Colour: "Brown",
			Value:  rand.Intn(300) + 300,
		}
	}
	puppyNegativeValue = func() puppy.Puppy {
		return puppy.Puppy{
			Breed:  "Type C",
			Colour: "Red",
			Value:  -1,
		}
	}
	puppyNegativeID = func() puppy.Puppy {
		return puppy.Puppy{
			Breed:  "Type D",
			Colour: "Blue",
			Value:  100,
			ID:     -1,
		}
	}
)

type storerSuite struct {
	suite.Suite
	store puppy.Storer
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

// Create
func (s *storerSuite) CreatePuppySuccessful() {
	assert := tassert.New(s.T())
	puppies := []puppy.Puppy{puppy0(), puppy1()}
	for i, p := range puppies {
		p := p
		id, err := s.store.CreatePuppy(&p)
		assert.NoError(err, "Creating p should be ok")
		puppies[i].ID = id
	}
	assert.Equal(puppies[0].ID, puppies[1].ID-1, "2nd id should be 1st +1")
}

func (s *storerSuite) CreatePuppyNegativeValue() {
	assert := tassert.New(s.T())
	newPuppy := puppyNegativeValue()
	_, err := s.store.CreatePuppy(&newPuppy)
	assert.Error(err, "negative ID should cause an error")
}

func (s *storerSuite) TestReadPuppy() {
	s.Run("OK", s.CreatePuppySuccessful)
	s.Run("NegativeValue", s.CreatePuppyNegativeValue)
}

// Read
func (s *storerSuite) ReadPuppySuccessful() {
	assert := tassert.New(s.T())
	newPuppy := puppy0()
	id, err := s.store.CreatePuppy(&newPuppy)
	assert.NoError(err, "Creating p should be ok")
	readPuppy, err := s.store.ReadPuppy(id)
	if assert.NoError(err, "Should be able to read puppy0 from store") {
		assert.Equal(&newPuppy, readPuppy, "store should return identical puppy")
	}
}

func (s *storerSuite) ReadPuppyIDDoesNotExist() {
	assert := tassert.New(s.T())
	_, err := s.store.ReadPuppy(1000)
	assert.Error(err, "Should get an error when attempting to read an non-existing puppy")
}

func (s *storerSuite) ReadPuppyNegativeID() {
	assert := tassert.New(s.T())
	_, err := s.store.ReadPuppy(-1)
	assert.Error(err, "negative ID should cause an error")
}

func (s *storerSuite) TestReadReadPuppy() {
	s.Run("OK", s.ReadPuppySuccessful)
	s.Run("NoID", s.ReadPuppyIDDoesNotExist)
	s.Run("NegativeID", s.ReadPuppyNegativeID)
}

// Update
func (s *storerSuite) UpdatePuppy() {
	assert := tassert.New(s.T())
	existingPuppy, err := s.store.ReadPuppy(0)
	assert.NoError(err, "Reading p should not return error")
	existingPuppy.Colour = "Purple"
	err = s.store.UpdatePuppy(0, existingPuppy)
	assert.NoError(err, "Update should not return any error")
	p, err := s.store.ReadPuppy(0)
	assert.NoError(err, "Reading p should not return error")
	assert.Equal(existingPuppy.Colour, p.Colour, "Updated colour missmatch")
}

func (s *storerSuite) UpdatePuppyCorruptedID() {
	assert := tassert.New(s.T())
	existingPuppy, err := s.store.ReadPuppy(0)
	assert.NoError(err, "Reading p should be ok")
	err = s.store.UpdatePuppy(1000, existingPuppy)
	assert.Error(err, "Should get an error when attempting to update with corrupted id")
}

func (s *storerSuite) UpdatePuppyIDDoesNotExist() {
	assert := tassert.New(s.T())
	newPuppy := puppy0()
	err := s.store.UpdatePuppy(1000, &newPuppy)
	assert.Error(err, "Should get an error when attempting to update an non-existing puppy")
}

func (s *storerSuite) UpdatePuppyNegativeID() {
	assert := tassert.New(s.T())
	newPuppy := puppyNegativeID()
	err := s.store.UpdatePuppy(newPuppy.ID, &newPuppy)
	assert.Error(err, "negative ID should cause an error")
}

func (s *storerSuite) UpdatePuppyNegativeValue() {
	assert := tassert.New(s.T())
	newPuppy := puppyNegativeValue()
	err := s.store.UpdatePuppy(newPuppy.ID, &newPuppy)
	assert.Error(err, "negative ID should cause an error")
}

func (s *storerSuite) TestUpdateReadPuppy() {
	s.Run("OK", s.UpdatePuppy)
	s.Run("CorruptedID", s.UpdatePuppyCorruptedID)
	s.Run("NoID", s.UpdatePuppyIDDoesNotExist)
	s.Run("NegativeID", s.UpdatePuppyNegativeID)
	s.Run("NegativeValue", s.UpdatePuppyNegativeValue)
}

// Delete
func (s *storerSuite) DeletePuppySuccessful() {
	assert := tassert.New(s.T())
	existingPuppy := puppy0()
	deleted, err := s.store.DeletePuppy(0)
	assert.NoError(err, "Delete should successfully delete a puppy")
	assert.True(deleted, "Delete should return true indicating a p was deleted")
	_, err = s.store.ReadPuppy(existingPuppy.ID)
	assert.Error(err, "Should not be able to read a deleted ID")
}

func (s *storerSuite) DeletePuppyIDDoesNotExist() {
	assert := tassert.New(s.T())
	_, err := s.store.DeletePuppy(1000)
	assert.Error(err, "Should not be able to delete p with non existing ID")
}

func (s *storerSuite) DeletePuppyNegativeID() {
	assert := tassert.New(s.T())
	_, err := s.store.DeletePuppy(-1)
	assert.Error(err, "negative ID should cause an error")
}

func (s *storerSuite) TestDeletePuppy() {
	s.Run("OK", s.DeletePuppySuccessful)
	s.Run("NoID", s.DeletePuppyIDDoesNotExist)
	s.Run("NegativeID", s.DeletePuppyNegativeID)
}

// Run suite
func TestStorer(t *testing.T) {
	suite.Run(t, &storerSuite{store: NewMemStore()})
	suite.Run(t, &storerSuite{store: NewSyncStore()})
}
