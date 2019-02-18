package main

import (
	"testing"

	tassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type storerSuite struct {
	suite.Suite
	store Storer
}

//TestStorerImpls will run the test suite for all types of storers
func TestStorerImpls(t *testing.T) {
	syncMapSuite := storerSuite{
		store: GetSyncStore()}
	suite.Run(t, &syncMapSuite)
	mapSuite := storerSuite{
		store: GetMapStore()}
	suite.Run(t, &mapSuite)
}

var (
	labrodar          = Puppy{"Labrodar", "Red", 34343.43}
	poddle            = Puppy{"Poddle", "White", 3343.43}
	chihuahua         = Puppy{"Chihuahua", "White", 3343.43}
	modifiedChihuahua = Puppy{"Chihuahua", "Red", 3343.43}
)

func (s *storerSuite) TestCreatePuppy() {
	assert := tassert.New(s.T())
	id := s.store.CreatePuppy(&labrodar)
	assert.Equal(1, id, "Creation failed")
	assert.NotEmpty(labrodar, "Creation failed")
}

func (s *storerSuite) TestReadPuppy() {
	assert := tassert.New(s.T())
	id := s.store.CreatePuppy(&poddle)
	myNewPuppy, _ := s.store.ReadPuppy(id)
	assert.Equal(*myNewPuppy, poddle, "Creation failed")
}

func (s *storerSuite) TestUpdatePuppy() {
	assert := tassert.New(s.T())
	id := s.store.CreatePuppy(&chihuahua)
	myNewPuppy, _ := s.store.ReadPuppy(id)
	_ = s.store.UpdatePuppy(id, &modifiedChihuahua)
	udpatedPuppy, _ := s.store.ReadPuppy(id)
	assert.NotEqual(myNewPuppy, udpatedPuppy, "Update failed")
	assert.Equal(udpatedPuppy.colour, "Red", "Update failed")
}

func (s *storerSuite) TestDeletePuppy() {
	assert := tassert.New(s.T())
	_ = s.store.DeletePuppy(1)
	actual, _ := s.store.ReadPuppy(1)
	assert.Nil(actual, "Deletion of puppy failed")
}
