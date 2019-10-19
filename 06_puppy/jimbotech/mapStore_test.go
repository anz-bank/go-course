package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type storesSuite struct {
	suite.Suite
	store  Storer
	mapper mapTest
}

func TestSuite(t *testing.T) {
	var ms Storer = NewMapStore()
	var sms Storer = &SyncMapStore{}
	suite.Run(t, &storesSuite{store: ms, mapper: ms.(mapTest)})
	suite.Run(t, &storesSuite{store: sms, mapper: sms.(mapTest)})
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

// TestMapStoreWithoutContructor may not be required as it may not be
// possible to create an instance of that type without using the constructor
func TestMapStoreWithoutContructor(t *testing.T) {
	var puppyStore MapStore
	pup := Puppy{1, "kelpie", "brown", "indispensable"}

	_, err := puppyStore.CreatePuppy(&pup)
	assert.Equal(t, ErrNotConstructed, err)
	err = puppyStore.UpdatePuppy(1, &pup)
	assert.Equal(t, ErrNotConstructed, err)
	_, err = puppyStore.ReadPuppy(1)
	assert.Equal(t, ErrNotConstructed, err)
	err = puppyStore.DeletePuppy(1)
	assert.Equal(t, ErrNotConstructed, err)
}

// TestCreateSuccess add to the store and verify
// by reading that it is in the store
func (s *storesSuite) TestCreateSuccess() {
	success, pup := create(s)
	if !s.True(success, "Create failed") {
		return
	}
	// now check by reading the value back and compare
	pup2, err2 := s.store.ReadPuppy(pup.ID)
	if !s.Nil(err2, "Read store should work") {
		return
	}
	s.Equal("kelpie", pup2.Breed)
	s.Equal("brown", pup2.Colour)
	s.Equal("indispensable", pup2.Value)
	s.Equal(*pup, *pup2)
}

func create(s *storesSuite) (bool, *Puppy) {
	pup := Puppy{1, "kelpie", "brown", "indispensable"}
	id, err := s.store.CreatePuppy(&pup)
	if !s.Nil(err, "Create on the store should have worked") {
		return false, nil
	}
	s.NotEqual(pup.ID, uint32(1))
	s.Equal(id, pup.ID, "Pup id must be set to actual id")
	return true, &pup
}

func (s *storesSuite) TestUpdateSuccess() {
	success, pup := create(s)
	if !success {
		s.Fail("Update failed")
		return
	}
	pup2 := Puppy{pup.ID, "kelpie", "black", "indispensable"}
	err := s.store.UpdatePuppy(pup.ID, &pup2)
	s.Nil(err)
	// now check by reading the updated value back and compare
	pup3, err2 := s.store.ReadPuppy(pup.ID)
	if s.Nil(err2, "Reading back updated value should work") {
		s.Equal(pup2, *pup3)
	}
}

//TestUpdateFailure checks the error returned when updating with an invalid id
func (s *storesSuite) TestUpdateFailure() {
	success, _ := create(s)
	if !success {
		s.Fail("Creating first puppy failed")
		return
	}
	pup2 := Puppy{1, "kelpie", "black", "indispensable"}
	err := s.store.UpdatePuppy(1, &pup2)
	success = s.NotNil(err, "Update on id 1 should have failed")
	if !success {
		return
	}
	st := fmt.Sprintf("no puppy with ID %v found", 1)
	s.Equal(st, err.Error())
}

func (s *storesSuite) TestDeleteSuccess() {
	success, pup := create(s)
	if !success {
		s.Fail("Creating puppy failed for delete test")
		return
	}
	err := s.store.DeletePuppy(pup.ID)
	if s.Nil(err, "Create puppy failed") {
		return
	}
	_, err = s.store.ReadPuppy(pup.ID)
	s.NotNil(err)
}

func (s *storesSuite) TestReadFailure() {
	pup2, err := s.store.ReadPuppy(1)
	s.Nil(pup2)
	s.NotNil(err)
	st := fmt.Sprintf("no puppy with ID %v found", 1)
	s.Equal(st, err.Error())
}

func (s *storesSuite) TestMapChanges() {
	s.Equal(0, s.mapper.length())
	pup := Puppy{1, "kelpie", "brown", "high"}
	id, err := s.store.CreatePuppy(&pup)
	if !s.Nil(err, "Create puppy failed") {
		return
	}
	s.Equal(1, s.mapper.length())
	pup2 := Puppy{id, "kelpie", "black", "low"}
	err = s.store.UpdatePuppy(id, &pup2)
	if !s.Nil(err, "Update puppy failed") {
		return
	}
	s.Equal(1, s.mapper.length())
	err = s.store.DeletePuppy(id)
	if !s.Nil(err, "Delete puppy failed") {
		return
	}
	s.Equal(0, s.mapper.length())
}
