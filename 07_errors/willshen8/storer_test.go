package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	Puppy1        = Puppy{1, "Jack Russell Terrier", "White and Brown", "1500"}
	Puppy2        = Puppy{1234, "Fox Terrier", "Black", "1300"}
	Puppy3        = Puppy{100, "German Shepperd", "Brown", "2000"}
	Puppy4        = Puppy{120, "Golden Retriever", "Golden", "2500"}
	Puppy5        = Puppy{200, "Chihuahua", "White", "500"}
	Puppy6        = Puppy{300, "Husky", "White", "3500"}
	Puppy7        = Puppy{700, "Pomeranian", "White", "700"}
	InvalidPuppy  = Puppy{22, "Pomeranian", "White", "-100"}
	ReadError     = Error{Code: NonExistentPuppy, Message: "Puppy ID can not be found, read operation failed."}
	UpdateError   = Error{Code: NonExistentPuppy, Message: "Puppy ID can not be found, update operation failed."}
	DeleteError   = Error{Code: NonExistentPuppy, Message: "Puppy ID can not be found, delete operation failed."}
	ErrorCreateID = uint32(0)
)

type StorerSuite struct {
	suite.Suite
	storer Storer
}

func (s *StorerSuite) TestCreatePuppySuccessfully() {
	m := assert.New(s.T())
	createdID, _ := s.storer.CreatePuppy(&Puppy1)
	readPuppy1, _ := s.storer.ReadPuppy(createdID)
	m.Equal(&Puppy1, readPuppy1)
}

func (s *StorerSuite) TestCreatePuppyWithValueThanZero() {
	createError := "Error code 1001: Puppy value can't be less than 0.\n"
	m := assert.New(s.T())
	var buf bytes.Buffer
	out = &buf
	createdID, error := s.storer.CreatePuppy(&InvalidPuppy)
	fmt.Fprintln(out, error)
	actualError := buf.String()
	m.Equal(createError, actualError)
	m.Equal(ErrorCreateID, createdID)
}

func (s *StorerSuite) TestCreateNextPuppyWithCorrectID() {
	m := assert.New(s.T())
	createID1, creatError1 := s.storer.CreatePuppy(&Puppy2)
	if createID1 != uint32(1) && creatError1 != nil {
		panic("Error creating puppy")
	}
	// create a second puppy should give an id of 2
	createdID2, _ := s.storer.CreatePuppy(&Puppy3)
	m.Equal(uint32(2), createdID2)
}

func (s *StorerSuite) TestReadPuppySuccessfully() {
	m := assert.New(s.T())
	returnedPuppyID4, _ := s.storer.CreatePuppy(&Puppy4)
	returnPuppy4, _ := s.storer.ReadPuppy(returnedPuppyID4)
	m.Equal(&Puppy4, returnPuppy4)
}

func (s *StorerSuite) TestReadNonExistentPuppy() {
	m := assert.New(s.T())
	returnPuppy, error := s.storer.ReadPuppy(100)
	m.Nil(returnPuppy)
	m.Equal(&ReadError, error)
}

func (s *StorerSuite) TestUpdatePuppySuccessfully() {
	returnedPuppyID, _ := s.storer.CreatePuppy(&Puppy5)
	updateResult, _ := s.storer.UpdatePuppy(returnedPuppyID, &Puppy6)
	s.Equal(true, updateResult)
}

func (s *StorerSuite) TestUpdateNonExistentID() {
	updateResult, error := s.storer.UpdatePuppy(888, &Puppy1)
	s.Equal(&UpdateError, error)
	s.Equal(false, updateResult)
}

func (s *StorerSuite) TestDeletePuppy() {
	createdPuppyID, _ := s.storer.CreatePuppy(&Puppy7)
	deleteResult, _ := s.storer.DeletePuppy(createdPuppyID)
	s.Equal(true, deleteResult)
	returnPuppy, _ := s.storer.ReadPuppy(createdPuppyID)
	s.Nil(returnPuppy)
}

func (s *StorerSuite) TestDeleteNonExistentPuppy() {
	deleteResult, error := s.storer.DeletePuppy(uint32(10000))
	s.Equal(false, deleteResult)
	s.Equal(&DeleteError, error)
}

func TestStorer(t *testing.T) {
	suite.Run(t, &StorerSuite{storer: NewMapStore()})
	suite.Run(t, &StorerSuite{storer: NewSyncStore()})
}
