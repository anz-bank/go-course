package main

import (
	"bytes"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StorerTestSuite struct {
	suite.Suite
	storer Storer
}

func (suite *StorerTestSuite) SetupTest() {
	suite.storer.CreatePuppy(Puppy{1, "Jack Russel", "white", 500.0})
}

func (suite *StorerTestSuite) TestCreate() {
	suite.storer.CreatePuppy(Puppy{12, "Komondor", "white", 1000.0})
	puppy := suite.storer.RetrievePuppy(12)
	suite.Equal(uint64(12), puppy.ID)
	suite.Equal("Komondor", puppy.Breed)
	suite.Equal("white", puppy.Colour)
	suite.Equal(1000.00, puppy.Value)
}

func (suite *StorerTestSuite) TestRead() {
	puppy := suite.storer.RetrievePuppy(1)
	suite.Equal(uint64(1), puppy.ID)
	suite.Equal("Jack Russel", puppy.Breed)
	suite.Equal("white", puppy.Colour)
	suite.Equal(500.00, puppy.Value)
}

func (suite *StorerTestSuite) TestUpdate() {
	suite.storer.UpdatePuppy(1, Puppy{1, "Jack Russel Terrier", "white and brown", 550.0})
	puppy := suite.storer.RetrievePuppy(1)
	suite.Equal(uint64(1), puppy.ID)
	suite.Equal("Jack Russel Terrier", puppy.Breed)
	suite.Equal("white and brown", puppy.Colour)
	suite.Equal(550.00, puppy.Value)
}

func (suite *StorerTestSuite) TestDelete() {
	suite.storer.DeletePuppy(1)
	puppy := suite.storer.RetrievePuppy(1)
	suite.Nil(puppy)
}

func TestMapStorer(t *testing.T) {
	s := StorerTestSuite{
		storer: MapStore{},
	}
	suite.Run(t, &s)
}

func TestSyncStorer(t *testing.T) {
	store := sync.Map{}
	s := StorerTestSuite{
		storer: SyncStore{
			store: &store,
		},
	}
	suite.Run(t, &s)
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("1 - Jack Russel Terrier [white and brown]: 550.000000")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output. Expected: %q - Actual: %q", expected, actual)
	}
}
