package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StorerSuite struct {
	suite.Suite
	storer        Storer
	Puppy1        Puppy
	Puppy2        Puppy
	Puppy3        Puppy
	Puppy4        Puppy
	Puppy5        Puppy
	Puppy6        Puppy
	Puppy7        Puppy
	InvalidPuppy  Puppy
	ErrorCreateID uint32
	ErrorCreate   string
	ErrorRead     string
	ErrorUpdate   string
	ErrorDelete   string
}

func (s *StorerSuite) SetupTest() {
	s.Puppy1 = Puppy{1, "Jack Russell Terrier", "White and Brown", "1500"}
	s.Puppy2 = Puppy{1234, "Fox Terrier", "Black", "1300"}
	s.Puppy3 = Puppy{100, "German Shepperd", "Brown", "2000"}
	s.Puppy4 = Puppy{120, "Golden Retriever", "Golden", "2500"}
	s.Puppy5 = Puppy{200, "Chihuahua", "White", "500"}
	s.Puppy6 = Puppy{300, "Husky", "White", "3500"}
	s.Puppy7 = Puppy{700, "Pomeranian", "White", "700"}
	s.InvalidPuppy = Puppy{22, "Pomeranian", "White", "-100"}

	s.ErrorCreate = "Error code 1: Puppy value can't be less than 0.\n"
	s.ErrorRead = "Error code 2: Puppy ID can not be found, read operation failed.\n"
	s.ErrorUpdate = "Error code 2: Puppy ID can not be found, update operation failed.\n"
	s.ErrorDelete = "Error code 2: Puppy ID can not be found, delete operation failed.\n"
	s.ErrorCreateID = 0
}

func (s *StorerSuite) TestCreatePuppySuccessfully() {
	m := assert.New(s.T())
	createdID, _ := s.storer.CreatePuppy(&s.Puppy1)

	readPuppy1, _ := s.storer.ReadPuppy(createdID)
	m.Equal(&s.Puppy1, readPuppy1)
}

func (s *StorerSuite) TestCreatePuppyWithValueThanZero() {
	m := assert.New(s.T())
	var buf bytes.Buffer
	out = &buf
	createdID, createError := s.storer.CreatePuppy(&s.InvalidPuppy)
	fmt.Fprintln(out, createError)
	actualError := buf.String()
	m.Equal(s.ErrorCreate, actualError)
	m.Equal(s.ErrorCreateID, createdID)
}

func (s *StorerSuite) TestCreateNextPuppyWithCorrectID() {
	m := assert.New(s.T())
	createID1, creatError1 := s.storer.CreatePuppy(&s.Puppy2)
	if createID1 != uint32(1) && creatError1 != nil {
		panic("Error creating puppy")
	}
	// create a second puppy should give an id of 2
	createdID2, _ := s.storer.CreatePuppy(&s.Puppy3)
	m.Equal(uint32(2), createdID2)
}

func (s *StorerSuite) TestReadPuppySuccessfully() {
	m := assert.New(s.T())
	returnedPuppyID4, _ := s.storer.CreatePuppy(&s.Puppy4)
	returnPuppy4, _ := s.storer.ReadPuppy(returnedPuppyID4)
	m.Equal(&s.Puppy4, returnPuppy4)
}

func (s *StorerSuite) TestReadNonExistentPuppy() {
	m := assert.New(s.T())
	var buf bytes.Buffer
	out = &buf
	returnPuppy, readError := s.storer.ReadPuppy(100)
	m.Nil(returnPuppy)
	fmt.Fprintln(out, readError)
	actualError := buf.String()
	m.Equal(s.ErrorRead, actualError)
}

func (s *StorerSuite) TestUpdatePuppySuccessfully() {
	returnedPuppyID, _ := s.storer.CreatePuppy(&s.Puppy5)
	updateResult, _ := s.storer.UpdatePuppy(returnedPuppyID, &s.Puppy6)
	s.Equal(true, updateResult)
}

func (s *StorerSuite) TestUpdateNonExistentID() {
	var buf bytes.Buffer
	out = &buf
	updateResult, updateError := s.storer.UpdatePuppy(888, &s.Puppy1)
	fmt.Fprintln(out, updateError)
	actualError := buf.String()
	s.Equal(s.ErrorUpdate, actualError)
	s.Equal(false, updateResult)
}

func (s *StorerSuite) TestDeletePuppy() {
	createdPuppyID, _ := s.storer.CreatePuppy(&s.Puppy7)
	deleteResult, _ := s.storer.DeletePuppy(createdPuppyID)
	s.Equal(true, deleteResult)
	returnPuppy, _ := s.storer.ReadPuppy(createdPuppyID)
	s.Nil(returnPuppy)
}

func (s *StorerSuite) TestDeleteNonExistentPuppy() {
	var buf bytes.Buffer
	out = &buf
	deleteResult, deleteError := s.storer.DeletePuppy(uint32(10000))
	fmt.Fprintln(out, deleteError)
	actualError := buf.String()
	s.Equal(false, deleteResult)
	s.Equal(s.ErrorDelete, actualError)
}

func TestStorer(t *testing.T) {
	suite.Run(t, &StorerSuite{storer: NewMapStore()})
	suite.Run(t, &StorerSuite{storer: NewSyncStore()})
}
