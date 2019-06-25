package main

import (
	"testing"

	"github.com/stretchr/testify/require"
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
	err := s.store.CreatePuppy(Puppy{ID: "1", Colour: "Black"})
	s.NoError(err, "Happy case puppy creation")

	p, _ := s.store.ReadPuppy("1")
	s.Equal("Black", p.Colour)
}

func (s *storerSuite) TestCreatePuppyAlreadyExists() {
	_ = s.store.CreatePuppy(Puppy{ID: "2"})
	err := s.store.CreatePuppy(Puppy{ID: "2", Colour: "Red"})
	s.Error(err, "Puppy creation fails if ID already exists")
}

func (s *storerSuite) TestReadPuppyHappyCase() {
	_ = s.store.CreatePuppy(Puppy{ID: "3", Colour: "Blue"})
	p, err := s.store.ReadPuppy("3")
	require.NoError(s.T(), err)
	s.Equal("Blue", p.Colour)
}

func (s *storerSuite) TestReadPuppyNonExisting() {
	_, err := s.store.ReadPuppy("4")
	s.Error(err)
}

func (s *storerSuite) TestUpdatePuppyHappyCase() {
	_ = s.store.CreatePuppy(Puppy{ID: "5", Colour: "Brown"})
	err := s.store.UpdatePuppy(Puppy{ID: "5", Colour: "Green"})
	require.NoError(s.T(), err)
	p, _ := s.store.ReadPuppy("5")
	s.Equal("Green", p.Colour)
}

func (s *storerSuite) TestUpdatePuppyNonExisting() {
	err := s.store.UpdatePuppy(Puppy{ID: "6"})
	s.Error(err)
}

func (s *storerSuite) TestDeletePuppyHappyCase() {
	_ = s.store.CreatePuppy(Puppy{ID: "7", Colour: "Brown"})
	deleted, err := s.store.DeletePuppy("7")
	require.NoError(s.T(), err)
	require.Equal(s.T(), true, deleted)

	_, err = s.store.ReadPuppy("7")
	s.Error(err, "Puppy gone after deletion")
}

func (s *storerSuite) TestDeletePuppyNonExisting() {
	deleted, err := s.store.DeletePuppy("8")
	require.Error(s.T(), err)
	s.Equal(false, deleted)
}

func TestStorers(t *testing.T) {
	suite.Run(t, &storerSuite{
		storeFactory: NewMapStore,
	})
	suite.Run(t, &storerSuite{
		storeFactory: NewSyncStore,
	})
}
