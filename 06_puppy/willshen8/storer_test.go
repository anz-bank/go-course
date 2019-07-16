package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StorerSuite struct {
	suite.Suite
	storer Storer
	Puppy1 Puppy
	Puppy2 Puppy
	Puppy3 Puppy
	Puppy4 Puppy
	Puppy5 Puppy
	Puppy6 Puppy
	Puppy7 Puppy
}

func (s *StorerSuite) SetupTest() {
	s.Puppy1 = Puppy{1, "Jack Russell Terrier", "White and Brown", "1500"}
	s.Puppy2 = Puppy{1234, "Fox Terrier", "Black", "1300"}
	s.Puppy3 = Puppy{100, "German Shepperd", "Brown", "2000"}
	s.Puppy4 = Puppy{120, "Golden Retriever", "Golden", "2500"}
	s.Puppy5 = Puppy{200, "Chihuahua", "White", "500"}
	s.Puppy6 = Puppy{300, "Husky", "White", "3500"}
	s.Puppy7 = Puppy{700, "Pomeranian", "White", "700"}
}

func (s *StorerSuite) TestCreatePuppySuccessfully() {
	m := assert.New(s.T())
	createdID := s.storer.CreatePuppy(&s.Puppy1)

	readPuppy1 := s.storer.ReadPuppy(createdID)
	m.Equal(&s.Puppy1, readPuppy1)
}

func (s *StorerSuite) TestCreateNextPuppyWithCorrectID() {
	m := assert.New(s.T())
	s.storer.CreatePuppy(&s.Puppy2)
	// create a second puppy should give an id of 2
	createdID := s.storer.CreatePuppy(&s.Puppy3)
	m.Equal(uint32(2), createdID)
}

func (s *StorerSuite) TestReadPuppySuccessfully() {
	m := assert.New(s.T())
	returnedPuppyID4 := s.storer.CreatePuppy(&s.Puppy4)
	returnPuppy4 := s.storer.ReadPuppy(returnedPuppyID4)
	m.Equal(&s.Puppy4, returnPuppy4)
}

func (s *StorerSuite) TestReadNonExistentPuppy() {
	m := assert.New(s.T())
	returnPuppy := s.storer.ReadPuppy(100)
	m.Nil(returnPuppy)
}

func (s *StorerSuite) TestUpdatePuppySuccessfully() {
	returnedPuppyID := s.storer.CreatePuppy(&s.Puppy5)
	updateResult := s.storer.UpdatePuppy(returnedPuppyID, &s.Puppy6)
	s.Equal(true, updateResult)
}

func (s *StorerSuite) TestUpdateNonExistentID() {
	updateResult := s.storer.UpdatePuppy(888, &s.Puppy1)
	s.Equal(false, updateResult)
}

func (s *StorerSuite) TestDeletePuppy() {
	createdPuppyID := s.storer.CreatePuppy(&s.Puppy7)
	deleteResult := s.storer.DeletePuppy(createdPuppyID)
	s.Equal(true, deleteResult)
	returnPuppy := s.storer.ReadPuppy(createdPuppyID)
	s.Nil(returnPuppy)

	deleteResult2 := s.storer.DeletePuppy(createdPuppyID)
	s.Equal(false, deleteResult2)
}

func TestStorer(t *testing.T) {
	suite.Run(t, &StorerSuite{storer: NewMapStore()})
	suite.Run(t, &StorerSuite{storer: NewSyncStore()})
}
