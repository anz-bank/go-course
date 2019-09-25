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

func TestStorer(t *testing.T) {
	m := NewPuppyStorer()
	ms := &SyncMapStore{}
	suite.Run(t, &storesSuite{store: m, mapper: m})
	suite.Run(t, &storesSuite{store: ms, mapper: ms})
}

// TestMapStoreWithoutContructor may not be required as it may not be
// possible to create an instance of that type without using the constructor
func TestMapStoreWithoutContructor(t *testing.T) {
	var puppyStore MapStore
	pup := Puppy{1, "kelpie", "brown", "indispensable"}

	_, err := puppyStore.CreatePuppy(&pup)
	f := func(t *testing.T) {
		assert.NotNil(t, err)
		assert.Equal(t, "store not created, call initializing constructor first", err.Error())
	}
	t.Run("CreatePuppy test", f)

	err = puppyStore.UpdatePuppy(1, &pup)
	t.Run("UpdatePuppy test", f)
	_, err = puppyStore.ReadPuppy(1)
	t.Run("ReadPuppy test", f)
	puppyStore.DeletePuppy(1)
}

// TestCreateSuccess add to the store and verify
// by reading that it is in the store and then remove it
// to leave the store in the original state
func (s *storesSuite) TestCreateSuccess() {
	asserter := assert.New(s.T())
	if success, pup := create(s); success {
		// now check by reading the value back and compare
		pup2, err2 := s.store.ReadPuppy(pup.ID)
		if asserter.Nil(err2, "Read store should work") {
			asserter.Equal(*pup, *pup2)
		}
		s.store.DeletePuppy(pup.ID)
	}
}

func create(s *storesSuite) (bool, *Puppy) {
	asserter := assert.New(s.T())
	pup := Puppy{1, "kelpie", "brown", "indispensable"}
	id, err := s.store.CreatePuppy(&pup)
	if asserter.Nil(err, "Create on the store should have worked") {
		asserter.NotEqual(pup.ID, uint32(1))
		asserter.Equal(id, pup.ID, "Pup id must be set to actual id")
		return true, &pup
	}
	return false, nil
}

func (s *storesSuite) TestUpdateSuccess() {
	asserter := assert.New(s.T())
	if success, pup := create(s); success {
		pup2 := Puppy{pup.ID, "kelpie", "black", "indispensable"}
		err := s.store.UpdatePuppy(pup.ID, &pup2)
		asserter.Nil(err)
		// now check by reading the updated value back and compare
		pup3, err2 := s.store.ReadPuppy(pup.ID)
		if asserter.Nil(err2, "Reading back updated value should work") {
			asserter.Equal(pup2, *pup3)
		}
		s.store.DeletePuppy(pup.ID)
	}
}

//TestUpdateFailure checks the error returned when updating with an invalid id
func (s *storesSuite) TestUpdateFailure() {
	asserter := assert.New(s.T())
	if success, pup := create(s); success {
		pup2 := Puppy{1, "kelpie", "black", "indispensable"}
		err := s.store.UpdatePuppy(1, &pup2)
		if asserter.NotNil(err, "Update on id 1 should have failed") {
			st := fmt.Sprintf("no puppy with ID %v found", 1)
			asserter.Equal(st, err.Error())
		}
		s.store.DeletePuppy(pup.ID)
	}
}

func (s *storesSuite) TestDeleteSuccess() {
	asserter := assert.New(s.T())
	if success, pup := create(s); success {
		s.store.DeletePuppy(pup.ID)
		_, err := s.store.ReadPuppy(pup.ID)
		asserter.NotNil(err)
	}
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
	asserter.Nil(err)
	asserter.Equal(1, s.mapper.length())
	pup2 := Puppy{id, "kelpie", "black", "low"}
	err = s.store.UpdatePuppy(id, &pup2)
	asserter.Nil(err)
	asserter.Equal(1, s.mapper.length())
	s.store.DeletePuppy(id)
	asserter.Equal(0, s.mapper.length())
}
