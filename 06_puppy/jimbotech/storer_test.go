package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type storesSuite struct {
	suite.Suite
	store  Storer
	mapper mapTest
}

const brown = "brown"
const black = "black"
const grey = "grey"

func TestSuite(t *testing.T) {
	suite.Run(t, &storesSuite{store: NewMapStore()})
	suite.Run(t, &storesSuite{store: &SyncMapStore{}})
}

//SetupTest creates the correct empty map for each test
func (s *storesSuite) SetupTest() {
	switch s.store.(type) {
	case MapStore:
		s.store = NewMapStore()
	case *SyncMapStore:
		s.store = &SyncMapStore{}
	default:
		s.Fail("Unknown Storer implementation")
	}
	s.mapper = s.store.(mapTest)
}

func (s *storesSuite) TestReadSuccess() {
	pup := create(s)
	// now check by reading the value back and compare
	pup2, err2 := s.store.ReadPuppy(pup.ID)
	s.Require().NoError(err2)
	s.Equal(brown, pup2.Colour)
	// modify the retured value to make sure the
	// value in the store does not change
	pup2.Colour = grey
	pup3, err2 := s.store.ReadPuppy(pup.ID)
	s.Require().NoError(err2)
	s.Equal(brown, pup3.Colour)
	s.NotEqual(pup2, pup3)
}

// TestCreateSuccess add to the store and verify
// by reading that it is in the store
func (s *storesSuite) TestCreateSuccess() {
	pup := create(s)
	// Now modify the original and make sure the
	// value in the store will not change
	pup.Colour = black
	// now check by reading the value back and compare
	pup2, err2 := s.store.ReadPuppy(pup.ID)
	s.Require().NoError(err2)
	s.Equal("kelpie", pup2.Breed)
	s.Equal(brown, pup2.Colour)
	s.Equal("indispensable", pup2.Value)
	s.True(pup2.Colour == brown)
	s.True(pup.Colour == black)
	s.NotEqual(pup, pup2)
}

func create(s *storesSuite) *Puppy {
	pup := Puppy{Breed: "kelpie", Colour: brown, Value: "indispensable"}
	id, err := s.store.CreatePuppy(&pup)
	s.Require().NoError(err)
	s.Require().NotEqual(pup.ID, uint32(1))
	s.Require().Equal(id, pup.ID, "Pup id must be set to actual id")
	return &pup
}

func (s *storesSuite) TestUpdateSuccess() {
	pup := create(s)
	pup2 := Puppy{Breed: "kelpie", Colour: black, Value: "indispensable"}
	err := s.store.UpdatePuppy(pup.ID, &pup2)
	s.Require().NoError(err)
	pup2.Colour = brown
	// now check by reading the updated value back and compare
	pup3, err2 := s.store.ReadPuppy(pup.ID)
	if s.Nil(err2, "Reading back updated value should work") {
		s.True(pup2.Colour == brown)
		s.True(pup3.Colour == black)
		s.NotEqual(pup2, *pup3)
	}
}

//TestUpdateFailure checks the error returned when updating with an invalid id
func (s *storesSuite) TestUpdateFailure() {
	create(s)
	pup2 := Puppy{Breed: "kelpie", Colour: black, Value: "indispensable"}
	err := s.store.UpdatePuppy(1, &pup2)
	success := s.NotNil(err, "Update on id 1 should have failed")
	if !success {
		return
	}
	st := fmt.Sprintf("no puppy with ID %v found", 1)
	s.Equal(st, err.Error())
}

func (s *storesSuite) TestDeleteSuccess() {
	pup := create(s)
	err := s.store.DeletePuppy(pup.ID)
	s.Require().NoError(err)
	_, err = s.store.ReadPuppy(pup.ID)
	s.NotNil(err)
}

func (s *storesSuite) TestReadFailure() {
	pup2, err := s.store.ReadPuppy(1)
	s.Require().Nil(pup2)
	s.Require().Error(err)
	st := fmt.Sprintf("no puppy with ID %v found", 1)
	s.Equal(st, err.Error())
}

func (s *storesSuite) TestMapChanges() {
	s.Equal(0, s.mapper.length())
	pup := Puppy{Breed: "kelpie", Colour: brown, Value: "high"}
	id, err := s.store.CreatePuppy(&pup)
	s.Require().Nil(err, "Create puppy failed")
	s.Equal(1, s.mapper.length())
	pup2 := Puppy{Breed: "kelpie", Colour: black, Value: "low"}
	err = s.store.UpdatePuppy(id, &pup2)
	s.Require().Nil(err, "Update puppy failed")
	s.Equal(1, s.mapper.length())
	err = s.store.DeletePuppy(id)
	s.Require().Nil(err, "Delete puppy failed")
	s.Equal(0, s.mapper.length())
}
