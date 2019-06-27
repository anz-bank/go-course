package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type storerSuite struct {
	suite.Suite
	store        Storer
	storeFactory func() Storer
}

func (s *storerSuite) SetupTest() {
	s.store = s.storeFactory()
}

func (s *storerSuite) TestCreatePuppyHappyCase() {
	puppy := Puppy{Colour: "Black"}
	id := s.store.CreatePuppy(puppy)
	puppyRead, err := s.store.ReadPuppy(id)
	s.Require().NoError(err)

	puppy.ID = id
	s.Equal(puppy, puppyRead)
}

func (s *storerSuite) TestReadPuppyHappyCase() {
	id := s.store.CreatePuppy(Puppy{Colour: "Blue"})
	p, err := s.store.ReadPuppy(id)
	s.Require().NoError(err)
	s.Equal("Blue", p.Colour)
}

func (s *storerSuite) TestReadPuppyNonExisting() {
	_, err := s.store.ReadPuppy("id_that_does_not_exist")
	s.Error(err)
}

func (s *storerSuite) TestUpdatePuppyHappyCase() {
	id := s.store.CreatePuppy(Puppy{Colour: "Brown"})
	err := s.store.UpdatePuppy(Puppy{ID: id, Colour: "Green"})
	s.Require().NoError(err)
	p, err := s.store.ReadPuppy(id)
	s.Require().NoError(err)
	s.Equal("Green", p.Colour)
}

func (s *storerSuite) TestUpdatePuppyNonExisting() {
	err := s.store.UpdatePuppy(Puppy{ID: "id_that_does_not_exist_either"})
	s.Error(err)
}

func (s *storerSuite) TestDeletePuppyHappyCase() {
	id := s.store.CreatePuppy(Puppy{Colour: "Brown"})
	err := s.store.DeletePuppy(id)
	s.Require().NoError(err)

	_, err = s.store.ReadPuppy(id)
	s.Error(err, "Puppy should be gone after deletion")
}

func (s *storerSuite) TestDeletePuppyNonExisting() {
	err := s.store.DeletePuppy("id_that_does_not_exist_again")
	s.Require().Error(err)
}

func TestStorers(t *testing.T) {
	suite.Run(t, &storerSuite{
		storeFactory: NewMapStore,
	})
	suite.Run(t, &storerSuite{
		storeFactory: NewSyncStore,
	})
}
