package main

import (
	"bytes"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StorerTestSuite struct {
	suite.Suite
	storer Storer
}

func (suite *StorerTestSuite) SetupTest() {
	tempPuppy := Puppy{id: 1, breed: "cat", color: "Black", value: "1"}
	suite.storer.CreatePuppy(&tempPuppy)
}

func (suite *StorerTestSuite) TestCreate() {
	crePup := Puppy{id: 1, breed: "cat", color: "Black", value: "1"}
	suite.storer.CreatePuppy(&crePup)
	returnPup := suite.storer.ReadPuppy(crePup.id)
	suite.Equal(returnPup.id, crePup.id)
	suite.Equal(returnPup.breed, crePup.breed)
}

func (suite *StorerTestSuite) TestRead() {
	existingPup := *suite.storer.ReadPuppy(1)
	tempPuppy := Puppy{id: 1, breed: "cat", color: "Black", value: "1"}
	suite.Equal(existingPup, tempPuppy)
	readNonExisting := suite.storer.ReadPuppy(2)
	suite.Nil(readNonExisting)
}

func (suite *StorerTestSuite) TestUpdate() {
	upPup := Puppy{id: 1, breed: "dog", color: "Brown", value: "2"}
	suite.storer.UpdatePuppy(&upPup)
	returnPup := suite.storer.ReadPuppy(upPup.id)
	suite.Equal(returnPup.id, upPup.id)
	suite.Equal(returnPup.breed, upPup.breed)
	suite.Equal(returnPup.breed, upPup.breed)
	suite.Equal(returnPup.value, upPup.value)
}

func (suite *StorerTestSuite) TestDelete() {
	deleteExisting := suite.storer.DeletePuppy(1)
	suite.Equal(deleteExisting, true)
	deleteNonExisting := suite.storer.DeletePuppy(2)
	suite.Equal(deleteNonExisting, false)
}

func (suite *StorerTestSuite) TestHasPuppy() {
	returnPup := suite.storer.HasPuppy(1)
	suite.Equal(returnPup, true)
}

func TestMapStorer(t *testing.T) {
	ms := StorerTestSuite{
		storer: MapStore{},
	}
	suite.Run(t, &ms)
}

func TestSyncStorer(t *testing.T) {
	store := sync.Map{}
	ss := StorerTestSuite{
		storer: SynStore{
			store: &store,
		},
	}
	suite.Run(t, &ss)
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	expected := `{1 dog Black 1}
{1 cat Brown 1}
`
	actual := buf.String()
	if expected != actual {
		t.Errorf("Unexpected output. Expected: %q - Actual: %q", expected, actual)
	}
}
