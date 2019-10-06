package store

import (
	"fmt"
	"testing"

	"github.com/anz-bank/go-course/11_notify/kasunfdo/pkg/puppy"
	"github.com/stretchr/testify/suite"
)

type StorerTest struct {
	suite.Suite
	store puppy.Storer
	id    uint64
}

func TestStore(t *testing.T) {
	suite.Run(t, &StorerTest{store: NewMapStore()})
	suite.Run(t, &StorerTest{store: NewSyncStore()})
}

func (suite *StorerTest) SetupTest() {
	suite.id, _ = suite.store.CreatePuppy(puppy.Puppy{Breed: "Labrador", Colour: "Cream", Value: 2999.99})
}

func (suite *StorerTest) TestCreatePuppy() {
	id, err := suite.store.CreatePuppy(puppy.Puppy{Breed: "German Shepard", Colour: "Brown", Value: 3499.99})
	suite.Greater(id, suite.id)
	suite.Nil(err)

	id, err = suite.store.CreatePuppy(puppy.Puppy{Breed: "Terrier", Colour: "White", Value: -3499.99})
	suite.Zero(id)
	suite.Require().IsType(&puppy.Error{}, err)
	customErr, _ := err.(*puppy.Error)
	suite.Equal(puppy.ErrInvalid, customErr.Code)
	suite.Equal("value of puppy is negative", customErr.Message)
}

func (suite *StorerTest) TestReadPuppy() {
	p, err := suite.store.ReadPuppy(suite.id)

	suite.Nil(err)
	suite.Equal(suite.id, p.ID)
	suite.Equal("Labrador", p.Breed)
	suite.Equal("Cream", p.Colour)
	suite.Equal(2999.99, p.Value)

	_, err = suite.store.ReadPuppy(100)
	suite.Require().IsType(&puppy.Error{}, err)
	customErr, _ := err.(*puppy.Error)
	suite.Equal(puppy.ErrNotFound, customErr.Code)
	suite.Equal("puppy with id: 100 is not found", customErr.Message)
}

func (suite *StorerTest) TestUpdatePuppy() {
	modifiedPuppy := puppy.Puppy{ID: suite.id, Breed: "Labrador Retriever", Colour: "Brown", Value: 3999.99}
	err := suite.store.UpdatePuppy(modifiedPuppy)

	suite.Nil(err)
	p, err := suite.store.ReadPuppy(suite.id)

	suite.Nil(err)
	suite.Equal(modifiedPuppy, p)

	err = suite.store.UpdatePuppy(puppy.Puppy{ID: suite.id, Breed: "Poodle", Colour: "White", Value: -1999.99})
	suite.Require().IsType(&puppy.Error{}, err)
	customErr, _ := err.(*puppy.Error)
	suite.Equal(puppy.ErrInvalid, customErr.Code)
	suite.Equal("value of puppy is negative", customErr.Message)

	err = suite.store.UpdatePuppy(puppy.Puppy{ID: 100, Breed: "Poodle", Colour: "White", Value: 1999.99})
	suite.Require().IsType(&puppy.Error{}, err)
	customErr, _ = err.(*puppy.Error)
	suite.Equal(puppy.ErrNotFound, customErr.Code)
	suite.Equal("puppy with id: 100 is not found", customErr.Message)
}

func (suite *StorerTest) TestDeletePuppy() {
	err := suite.store.DeletePuppy(suite.id)
	suite.Nil(err)

	_, err = suite.store.ReadPuppy(suite.id)
	suite.Require().IsType(&puppy.Error{}, err)
	customErr, _ := err.(*puppy.Error)
	suite.Equal(puppy.ErrNotFound, customErr.Code)
	suite.Equal(fmt.Sprintf("puppy with id: %v is not found", suite.id), customErr.Message)

	err = suite.store.DeletePuppy(suite.id)
	suite.Require().IsType(&puppy.Error{}, err)
	customErr, _ = err.(*puppy.Error)
	suite.Equal(puppy.ErrNotFound, customErr.Code)
	suite.Equal(fmt.Sprintf("puppy with id: %v is not found", suite.id), customErr.Message)
}
