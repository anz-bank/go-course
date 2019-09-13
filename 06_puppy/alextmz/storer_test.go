package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type storerSuite struct {
	suite.Suite
	st Storer
}

const IDPuppyDoesNotExist = 99999

func (su *storerSuite) TestCreatePuppy() {
	p1 := Puppy{Breed: "Dogo", Colour: "White", Value: 500}
	p2 := Puppy{Breed: "Mastiff", Colour: "Brindle", Value: 700}
	expected := p1

	// can we create without error?
	err := su.st.CreatePuppy(&p1)
	su.Nil(err)

	// do we error when creating something that already exists?
	err = su.st.CreatePuppy(&p1)
	su.NotNil(err)

	// what we create and what we read back match?
	actual, _ := su.st.ReadPuppy(p1.ID)
	actual.ID = 0
	su.Equal(expected, actual)

	// do we error when trying to create an already identified Puppy?
	p2.ID = IDPuppyDoesNotExist
	err = su.st.CreatePuppy(&p2)
	su.NotNil(err)

	// cleanup
	err = su.st.DeletePuppy(p1.ID)
	su.Nil(err)
}

func (su *storerSuite) TestReadPuppy() {
	p1 := Puppy{Breed: "Dogo", Colour: "White", Value: 500}
	expected := p1
	err := su.st.CreatePuppy(&p1)
	su.Nil(err)

	// can we read without error?
	_, err = su.st.ReadPuppy(p1.ID)
	su.Nil(err)

	// do we error when reading what doesn't exist?
	_, err = su.st.ReadPuppy(IDPuppyDoesNotExist)
	su.NotNil(err)

	// do the read contents match what we expect?
	actual, err := su.st.ReadPuppy(p1.ID)
	su.Nil(err)
	actual.ID = 0
	su.Equal(expected, actual)

	// cleanup
	err = su.st.DeletePuppy(p1.ID)
	su.Nil(err)
}

func (su *storerSuite) TestUpdatePuppy() {
	p1 := Puppy{Breed: "Dogo", Colour: "White", Value: 500}
	p2 := Puppy{Breed: "Mastiff", Colour: "Brindle", Value: 700}
	expected := p1
	err := su.st.CreatePuppy(&p1)
	su.Nil(err)

	// we can update without error?
	expectColour := "Black"
	p1.Colour = expectColour
	err = su.st.UpdatePuppy(p1)
	su.Nil(err)

	// updated content matches what we expect?
	actual, err := su.st.ReadPuppy(p1.ID)
	su.Nil(err)
	expected.Colour = expectColour
	actual.ID = 0
	su.Equal(expected, actual)

	// do we error when trying to update what doesn't exist?
	p2.ID = IDPuppyDoesNotExist
	err = su.st.UpdatePuppy(p2)
	su.NotNil(err)

	//cleanup
	err = su.st.DeletePuppy(p1.ID)
	su.Nil(err)
}

func (su *storerSuite) TestDeletePuppy() {
	// setup
	p1 := Puppy{Breed: "Dogo", Colour: "White", Value: 500}
	err := su.st.CreatePuppy(&p1)
	su.Nil(err)

	// can we delete without error?
	err = su.st.DeletePuppy(p1.ID)
	su.Nil(err)

	// after we delete, can we read the data back?
	p, err := su.st.ReadPuppy(p1.ID)
	su.NotNil(err)
	su.Equal(p, Puppy{ID: 0, Breed: "", Colour: "", Value: 0})

	// do we err when trying to delete what doesn't exist?
	err = su.st.DeletePuppy(IDPuppyDoesNotExist)
	su.NotNil(err)
}

func Test_Suite(t *testing.T) {
	suite.Run(t, &storerSuite{st: NewSyncStore()})
	suite.Run(t, &storerSuite{st: NewmapStore()})
}
