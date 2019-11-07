package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type storerSuite struct {
	suite.Suite
	st Storer
}

func (su *storerSuite) TestCreatePuppy() {
	// can we create without error?
	p1 := Puppy{Breed: "Dogo", Colour: "White", Value: 500}
	expected := p1
	err := su.st.CreatePuppy(&p1)
	su.NoError(err)

	// do we error when creating something that already exists?
	su.Run("NoErrorOnCreate", func() {
		err = su.st.CreatePuppy(&p1)
		su.Error(err)
	})

	// what we create and what we read back match?
	su.Run("MatchCreatedAndRead", func() {
		actual, _ := su.st.ReadPuppy(p1.ID)
		actual.ID = 0
		su.Equal(expected, actual)
	})

	// do we error when trying to create a puppy from a nil pointer?
	su.Run("ErrorNilPuppy", func() {
		var p4 *Puppy
		err = su.st.CreatePuppy(p4)
		su.Error(err)
	})

	// do we error when trying to create an already identified Puppy?
	su.Run("ErrorAlreadyIdentifiedPuppy", func() {
		p2 := Puppy{Breed: "Mastiff", Colour: "Brindle", Value: 700}
		p2.ID = 99999
		err = su.st.CreatePuppy(&p2)
		su.Error(err)
	})

	// do we error with the right code when trying to create a puppy with Value < 0?
	su.Run("ErrorNegativeValue", func() {
		p3 := Puppy{Breed: "Fila", Colour: "Golden", Value: -900}
		err = su.st.CreatePuppy(&p3)
		su.Error(err)
		su.Require().IsType(Error{}, err)
		su.Equal(ErrNegativeValue, err.(Error).Code)
	})

	// cleanup
	err = su.st.DeletePuppy(p1.ID)
	su.NoError(err)
}

func (su *storerSuite) TestReadPuppy() {
	// setup
	p1 := Puppy{Breed: "Dogo", Colour: "White", Value: 500}
	expected := p1
	err := su.st.CreatePuppy(&p1)
	su.NoError(err)

	// can we read without error?
	su.Run("NoErrorRead", func() {
		_, err = su.st.ReadPuppy(p1.ID)
		su.NoError(err)
	})
	// do we error when reading what doesn't exist?
	su.Run("ErrorPuppyDoesNotExist", func() {
		_, err = su.st.ReadPuppy(99999)
		su.Error(err)
	})

	// do the read contents match what we expect?
	su.Run("NoErrorReadPuppyMatches", func() {
		actual, err := su.st.ReadPuppy(p1.ID)
		su.NoError(err)
		actual.ID = 0
		su.Equal(expected, actual)
	})

	// cleanup
	err = su.st.DeletePuppy(p1.ID)
	su.NoError(err)
}

func (su *storerSuite) TestUpdatePuppy() {
	// setup
	p1 := Puppy{Breed: "Dogo", Colour: "White", Value: 500}
	expected := p1
	err := su.st.CreatePuppy(&p1)
	su.NoError(err)
	p2 := Puppy{Breed: "Mastiff", Colour: "Brindle", Value: 700}
	err = su.st.CreatePuppy(&p2)
	su.NoError(err)

	// we can update without error?
	su.Run("NoErrorOnUpdate", func() {
		p1.Colour = "Black"
		err = su.st.UpdatePuppy(p1)
		su.NoError(err)
	})

	// updated content matches what we expect?
	su.Run("MatchUpdatedPuppy", func() {
		actual, err := su.st.ReadPuppy(p1.ID)
		su.NoError(err)
		expected.Colour = "Black"
		actual.ID = 0
		su.Equal(expected, actual)
	})
	// do we error with the right code when trying to update a puppy with Value < 0?
	su.Run("ErrorUpdateNegativeValue", func() {
		p2.Value = -10
		err = su.st.UpdatePuppy(p2)
		su.Error(err)
		su.Require().IsType(Error{}, err)
		su.Equal(ErrNegativeValue, err.(Error).Code)
	})

	// do we error when trying to update what doesn't exist?
	su.Run("ErrorUpdateNonexistentPuppy", func() {
		p3 := Puppy{Breed: "Mastiff", Colour: "Brindle", Value: 700}
		p3.ID = 99999
		err = su.st.UpdatePuppy(p3)
		su.Error(err)
	})

	//cleanup
	err = su.st.DeletePuppy(p1.ID)
	su.NoError(err)
}

func (su *storerSuite) TestDeletePuppy() {
	// setup
	p1 := Puppy{Breed: "Dogo", Colour: "White", Value: 500}
	err := su.st.CreatePuppy(&p1)
	su.NoError(err)

	// can we delete without error?
	su.Run("DeleteWithNoError", func() {
		err = su.st.DeletePuppy(p1.ID)
		su.NoError(err)
	})

	// after we delete, can we read the data back?
	su.Run("ErrorReadDeletedPuppy", func() {
		p, err := su.st.ReadPuppy(p1.ID)
		su.Error(err)
		su.Equal(p, Puppy{ID: 0, Breed: "", Colour: "", Value: 0})
	})

	// do we err when trying to delete what doesn't exist?
	su.Run("ErrorDeleteNonexistentPuppy", func() {
		err = su.st.DeletePuppy(99999)
		su.Error(err)
	})

	// no cleanup needed: all data taken care of already.
}

func Test_Suite(t *testing.T) {

	t.Run("SyncStore", func(t *testing.T) {
		suite.Run(t, &storerSuite{st: NewSyncStore()})
	})
	// suite.Run(t, &storerSuite{st: NewSyncStore()})
	t.Run("MapStore", func(t *testing.T) {
		suite.Run(t, &storerSuite{st: NewMapStore()})
	})

}
