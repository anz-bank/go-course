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
	ok, err := suite.storer.CreatePuppy(Puppy{1, "Jack Russel", "white", 500.0})
	if !ok && !err.(*PupError).PupAlreadyExists() {
		suite.Fail("Failed to setup test suite")
	}
}

func (suite *StorerTestSuite) TestCreate() {
	ok, _ := suite.storer.CreatePuppy(Puppy{12, "Komondor", "white", 1000.0})
	puppy, _ := suite.storer.RetrievePuppy(12)
	suite.Equal(int64(12), puppy.ID)
	suite.Equal("Komondor", puppy.Breed)
	suite.Equal("white", puppy.Colour)
	suite.Equal(1000.00, puppy.Value)
	suite.Equal(true, ok)
}

func (suite *StorerTestSuite) TestCreateInvalid() {
	ok, err := suite.storer.CreatePuppy(Puppy{12, "Komondor", "white", -1000.0})
	suite.Equal(false, ok)
	suite.Equal("[1] Puppy value must be non negative: -1000.000000", err.Error())
	suite.Equal(true, err.(*PupError).InvalidPupValue())
}

func (suite *StorerTestSuite) TestRead() {
	puppy, _ := suite.storer.RetrievePuppy(1)
	suite.Equal(int64(1), puppy.ID)
	suite.Equal("Jack Russel", puppy.Breed)
	suite.Equal("white", puppy.Colour)
	suite.Equal(500.00, puppy.Value)
}

func (suite *StorerTestSuite) TestUpdate() {
	ok, _ := suite.storer.UpdatePuppy(1, Puppy{1, "Jack Russel Terrier", "white and brown", 550.0})
	puppy, _ := suite.storer.RetrievePuppy(1)
	suite.Equal(int64(1), puppy.ID)
	suite.Equal("Jack Russel Terrier", puppy.Breed)
	suite.Equal("white and brown", puppy.Colour)
	suite.Equal(550.00, puppy.Value)
	suite.Equal(true, ok)
}

func (suite *StorerTestSuite) TestUpdateMissing() {
	ok, err := suite.storer.UpdatePuppy(10, Puppy{10, "Jack Russel Terrier", "white and brown", 550.0})
	suite.Equal(false, ok)
	suite.Equal("[0] Puppy not found: 10", err.Error())
	suite.Equal(true, err.(*PupError).IsMissingPup())
}

func (suite *StorerTestSuite) TestUpdateInvalid() {
	ok, err := suite.storer.UpdatePuppy(1, Puppy{1, "Jack Russel Terrier", "white and brown", -550.0})
	suite.Equal(false, ok)
	suite.Equal("[1] Puppy value must be non negative: -550.000000", err.Error())
	suite.Equal(true, err.(*PupError).InvalidPupValue())
}

func (suite *StorerTestSuite) TestDelete() {
	ok, _ := suite.storer.DeletePuppy(1)
	puppy, _ := suite.storer.RetrievePuppy(1)
	suite.Nil(puppy)
	suite.Equal(true, ok)
}

func (suite *StorerTestSuite) TestDeleteMissing() {
	ok, err := suite.storer.DeletePuppy(10)
	suite.Equal(false, ok)
	suite.Equal("[0] Puppy not found: 10", err.Error())
	suite.Equal(true, err.(*PupError).IsMissingPup())
}

func TestMapStorer(t *testing.T) {
	s := StorerTestSuite{
		storer: MapStore{},
	}
	suite.Run(t, &s)
}

func TestSyncStorer(t *testing.T) {
	store := sync.Map{}
	lock := sync.Mutex{}

	s := StorerTestSuite{
		storer: SyncStore{
			store: &store,
			lock:  &lock,
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
