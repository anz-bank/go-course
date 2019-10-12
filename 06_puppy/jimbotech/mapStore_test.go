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
	mapper mapCheck
}

func TestSuite(t *testing.T) {
	var ms Storer = NewMapStore()
	var sms Storer = &SyncMapStore{}
	suite.Run(t, &storesSuite{store: ms, mapper: ms.(mapCheck)})
	suite.Run(t, &storesSuite{store: sms, mapper: sms.(mapCheck)})
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
	s.mapper = s.store.(mapCheck)
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
	asserter := assert.New(s.T())
	success, pup := create(s)
	if !asserter.True(success, "Create failed") {
		return
	}
	// now check by reading the value back and compare
	pup2, err2 := s.store.ReadPuppy(pup.ID)
	if !asserter.Nil(err2, "Read store should work") {
		return
	}
	asserter.Equal("kelpie", pup2.Breed)
	asserter.Equal("brown", pup2.Colour)
	asserter.Equal("indispensable", pup2.Value)
	asserter.Equal(*pup, *pup2)
}

func create(s *storesSuite) (bool, *Puppy) {
	asserter := assert.New(s.T())
	pup := Puppy{1, "kelpie", "brown", "indispensable"}
	id, err := s.store.CreatePuppy(&pup)
	if !asserter.Nil(err, "Create on the store should have worked") {
		return false, nil
	}
	asserter.NotEqual(pup.ID, uint32(1))
	asserter.Equal(id, pup.ID, "Pup id must be set to actual id")
	return true, &pup
}

func (s *storesSuite) TestUpdateSuccess() {
	asserter := assert.New(s.T())
	success, pup := create(s)
	if !success {
		asserter.Fail("Update failed")
		return
	}
	pup2 := Puppy{pup.ID, "kelpie", "black", "indispensable"}
	err := s.store.UpdatePuppy(pup.ID, &pup2)
	asserter.Nil(err)
	// now check by reading the updated value back and compare
	pup3, err2 := s.store.ReadPuppy(pup.ID)
	if asserter.Nil(err2, "Reading back updated value should work") {
		asserter.Equal(pup2, *pup3)
	}
}

//TestUpdateFailure checks the error returned when updating with an invalid id
func (s *storesSuite) TestUpdateFailure() {
	asserter := assert.New(s.T())
	success, _ := create(s)
	if !success {
		asserter.Fail("Creating first puppy failed")
		return
	}
	pup2 := Puppy{1, "kelpie", "black", "indispensable"}
	err := s.store.UpdatePuppy(1, &pup2)
	success = asserter.NotNil(err, "Update on id 1 should have failed")
	if !success {
		return
	}
	st := fmt.Sprintf("no puppy with ID %v found", 1)
	asserter.Equal(st, err.Error())
}

func (s *storesSuite) TestDeleteSuccess() {
	asserter := assert.New(s.T())
	success, pup := create(s)
	if !success {
		asserter.Fail("Creating puppy failed for delete test")
		return
	}
	err := s.store.DeletePuppy(pup.ID)
	if asserter.Nil(err, "Create puppy failed") {
		return
	}
	_, err = s.store.ReadPuppy(pup.ID)
	asserter.NotNil(err)
}

func (s *storesSuite) TestReadFailure() {
	asserter := assert.New(s.T())
	pup2, err := s.store.ReadPuppy(1)
	asserter.Nil(pup2)
	asserter.NotNil(err)
	st := fmt.Sprintf("no puppy with ID %v found", 1)
	asserter.Equal(st, err.Error())
}

func (s *storesSuite) TestMapChanges() {
	asserter := assert.New(s.T())
	asserter.Equal(0, s.mapper.length())
	pup := Puppy{1, "kelpie", "brown", "high"}
	id, err := s.store.CreatePuppy(&pup)
	if !asserter.Nil(err, "Create puppy failed") {
		return
	}
	asserter.Equal(1, s.mapper.length())
	pup2 := Puppy{id, "kelpie", "black", "low"}
	err = s.store.UpdatePuppy(id, &pup2)
	if !asserter.Nil(err, "Update puppy failed") {
		return
	}
	asserter.Equal(1, s.mapper.length())
	err = s.store.DeletePuppy(id)
	if !asserter.Nil(err, "Delete puppy failed") {
		return
	}
	asserter.Equal(0, s.mapper.length())
}
