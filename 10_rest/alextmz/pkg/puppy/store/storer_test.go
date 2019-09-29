package store

import (
	"testing"

	"github.com/anz-bank/go-course/10_rest/alextmz/pkg/puppy"
	"github.com/stretchr/testify/suite"
)

type storerSuite struct {
	suite.Suite
	st puppy.Storer
}

const (
	IDPuppyDoesNotExist    = 99999
	IDPuppyInvalidNegative = -1
)

func (su *storerSuite) TestCreatePuppy() {
	// can we create without error?
	p1 := puppy.Puppy{Breed: "Dogo", Colour: "White", Value: 500}
	// this value copy is here, not below, to guarantee that it is not going
	// to be modified before being used on the relevant test
	expected := p1
	err := su.st.CreatePuppy(&p1)
	su.NoError(err)
	// do we error when creating something that already exists?
	err = su.st.CreatePuppy(&p1)
	su.NotNil(err)
	// what we create and what we read back match?
	actual, _ := su.st.ReadPuppy(p1.ID)
	actual.ID = 0
	su.Equal(expected, actual)
	// do we error when trying to create a puppy from a nil pointer?
	var p4 *puppy.Puppy
	err = su.st.CreatePuppy(p4)
	su.NotNil(err)
	// do we error when trying to create an already identified Puppy?
	p2 := puppy.Puppy{Breed: "Mastiff", Colour: "Brindle", Value: 700}
	p2.ID = IDPuppyDoesNotExist
	err = su.st.CreatePuppy(&p2)
	su.NotNil(err)
	// do we error when trying to create a puppy that has ID < 0?
	p3 := puppy.Puppy{ID: -1, Breed: "Fila", Colour: "Golden", Value: 900}
	err = su.st.CreatePuppy(&p3)
	su.NotNil(err)
	// cleanup
	err = su.st.DeletePuppy(p1.ID)
	su.NoError(err)
}

func (su *storerSuite) TestReadPuppy() {
	// setup
	p1 := puppy.Puppy{Breed: "Dogo", Colour: "White", Value: 500}
	expected := p1
	err := su.st.CreatePuppy(&p1)
	su.NoError(err)
	// can we read without error?
	_, err = su.st.ReadPuppy(p1.ID)
	su.NoError(err)
	// do we error when reading what doesn't exist?
	_, err = su.st.ReadPuppy(IDPuppyDoesNotExist)
	su.NotNil(err)
	// do the read contents match what we expect?
	actual, err := su.st.ReadPuppy(p1.ID)
	su.NoError(err)
	actual.ID = 0
	su.Equal(expected, actual)
	// do we error when trying to read a puppy with ID < 0?
	_, err = su.st.ReadPuppy(IDPuppyInvalidNegative)
	su.NotNil(err)
	// cleanup
	err = su.st.DeletePuppy(p1.ID)
	su.NoError(err)
}

func (su *storerSuite) TestUpdatePuppy() {
	// setup
	p1 := puppy.Puppy{Breed: "Dogo", Colour: "White", Value: 500}
	expected := p1
	err := su.st.CreatePuppy(&p1)
	su.NoError(err)
	// we can update without error?
	expectColour := "Black"
	p1.Colour = expectColour
	err = su.st.UpdatePuppy(p1)
	su.NoError(err)
	// updated content matches what we expect?
	actual, err := su.st.ReadPuppy(p1.ID)
	su.NoError(err)
	expected.Colour = expectColour
	actual.ID = 0
	su.Equal(expected, actual)
	// do we error when trying to update what doesn't exist?
	p2 := puppy.Puppy{Breed: "Mastiff", Colour: "Brindle", Value: 700}
	p2.ID = IDPuppyDoesNotExist
	err = su.st.UpdatePuppy(p2)
	su.NotNil(err)
	// do we error when trying to update a puppy with ID < 0?
	p2.ID = IDPuppyInvalidNegative
	err = su.st.UpdatePuppy(p2)
	su.NotNil(err)
	//cleanup
	err = su.st.DeletePuppy(p1.ID)
	su.NoError(err)
}

func (su *storerSuite) TestDeletePuppy() {
	// setup
	p1 := puppy.Puppy{Breed: "Dogo", Colour: "White", Value: 500}
	err := su.st.CreatePuppy(&p1)
	su.NoError(err)
	// can we delete without error?
	err = su.st.DeletePuppy(p1.ID)
	su.NoError(err)
	// after we delete, can we read the data back?
	p, err := su.st.ReadPuppy(p1.ID)
	su.NotNil(err)
	su.Equal(p, puppy.Puppy{ID: 0, Breed: "", Colour: "", Value: 0})
	// do we err when trying to delete what doesn't exist?
	err = su.st.DeletePuppy(IDPuppyDoesNotExist)
	su.Error(err)
	// do we error when trying to delete a puppy with ID < 0?
	err = su.st.DeletePuppy(IDPuppyInvalidNegative)
	su.NotNil(err)
}

func Test_Suite(t *testing.T) {
	suite.Run(t, &storerSuite{st: NewSyncStore()})
	suite.Run(t, &storerSuite{st: NewmapStore()})
}
