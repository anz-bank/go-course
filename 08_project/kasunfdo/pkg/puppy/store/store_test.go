package store

import (
	"fmt"
	"os"
	"testing"

	"github.com/anz-bank/go-course/08_project/kasunfdo/pkg/puppy"
	"github.com/stretchr/testify/suite"
)

type StorerTest struct {
	suite.Suite
	store puppy.Storer
	id    uint64
}

func (suite *StorerTest) SetupTest() {
	os.RemoveAll("./test")
	suite.id, _ = suite.store.CreatePuppy(puppy.Puppy{Breed: "Labrador", Colour: "Cream", Value: 2999.99})
}

func (suite *StorerTest) TearDownSuite() {
	os.RemoveAll("./test")
}

func (suite *StorerTest) TestCreatePuppy() {
	id, err := suite.store.CreatePuppy(puppy.Puppy{Breed: "German Shepard", Colour: "Brown", Value: 3499.99})
	suite.True(id > 1)
	suite.Nil(err)

	id, err = suite.store.CreatePuppy(puppy.Puppy{Breed: "Terrier", Colour: "White", Value: -3499.99})
	suite.True(id == 0)
	suite.Require().IsType(&puppy.Error{}, err)
	customErr, _ := err.(*puppy.Error)
	suite.Equal(puppy.ErrInvalid, customErr.Code)
	suite.Equal("invalid input: value of puppy is negative", customErr.Message)
}

func (suite *StorerTest) TestReadPuppy() {
	p, err := suite.store.ReadPuppy(suite.id)

	suite.Nil(err)
	suite.Equal(p.ID, suite.id)
	suite.Equal(p.Breed, "Labrador")
	suite.Equal(p.Colour, "Cream")
	suite.Equal(p.Value, 2999.99)

	_, err = suite.store.ReadPuppy(100)
	suite.Require().IsType(&puppy.Error{}, err)
	customErr, _ := err.(*puppy.Error)
	suite.Equal(puppy.ErrNotFound, customErr.Code)
	suite.Equal("not found: puppy with id: 100 is not found", customErr.Message)
}

func (suite *StorerTest) TestUpdatePuppy() {
	err := suite.store.UpdatePuppy(puppy.Puppy{ID: suite.id, Breed: "Labrador Retriever", Colour: "Brown", Value: 3999.99})

	suite.Nil(err)
	p, err := suite.store.ReadPuppy(suite.id)

	suite.Nil(err)
	suite.Equal(p.ID, suite.id)
	suite.Equal(p.Breed, "Labrador Retriever")
	suite.Equal(p.Colour, "Brown")
	suite.Equal(p.Value, 3999.99)

	err = suite.store.UpdatePuppy(puppy.Puppy{ID: suite.id, Breed: "Poodle", Colour: "White", Value: -1999.99})
	suite.Require().IsType(&puppy.Error{}, err)
	customErr, _ := err.(*puppy.Error)
	suite.Equal(puppy.ErrInvalid, customErr.Code)
	suite.Equal("invalid input: value of puppy is negative", customErr.Message)

	err = suite.store.UpdatePuppy(puppy.Puppy{ID: 100, Breed: "Poodle", Colour: "White", Value: 1999.99})
	suite.Require().IsType(&puppy.Error{}, err)
	customErr, _ = err.(*puppy.Error)
	suite.Equal(puppy.ErrNotFound, customErr.Code)
	suite.Equal("not found: puppy with id: 100 is not found", customErr.Message)
}

func (suite *StorerTest) TestDeletePuppy() {
	err := suite.store.DeletePuppy(suite.id)
	suite.Nil(err)

	_, err = suite.store.ReadPuppy(suite.id)
	suite.Require().IsType(&puppy.Error{}, err)
	customErr, _ := err.(*puppy.Error)
	suite.Equal(puppy.ErrNotFound, customErr.Code)
	suite.Equal(fmt.Sprintf("not found: puppy with id: %v is not found", suite.id), customErr.Message)

	err = suite.store.DeletePuppy(suite.id)
	suite.Require().IsType(&puppy.Error{}, err)
	customErr, _ = err.(*puppy.Error)
	suite.Equal(puppy.ErrNotFound, customErr.Code)
	suite.Equal(fmt.Sprintf("not found: puppy with id: %v is not found", suite.id), customErr.Message)
}

func TestStore(t *testing.T) {
	suite.Run(t, &StorerTest{store: NewMapStore()})
	suite.Run(t, &StorerTest{store: NewSyncStore()})
}
