package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type storerSuite struct {
	suite.Suite
	storeFactory func() Storer
}

func (s *storerSuite) TestCreatePuppyHappyCase() {
	store := s.storeFactory()
	err := store.CreatePuppy(Puppy{ID: "1", Colour: "Black"})
	s.NoError(err, "Happy case puppy creation")

	p, _ := store.ReadPuppy("1")
	s.Equal("Black", p.Colour)
}

func (s *storerSuite) TestCreatePuppyAlreadyExists() {
	store := s.storeFactory()
	_ = store.CreatePuppy(Puppy{ID: "2"})
	err := store.CreatePuppy(Puppy{ID: "2", Colour: "Red"})
	s.Error(err, "Puppy creation fails if ID already exists")
}

func (s *storerSuite) TestReadPuppyHappyCase() {
	store := s.storeFactory()
	_ = store.CreatePuppy(Puppy{ID: "3", Colour: "Blue"})
	p, err := store.ReadPuppy("3")
	require.NoError(s.T(), err)
	s.Equal("Blue", p.Colour)
}

func (s *storerSuite) TestReadPuppyNonExisting() {
	store := s.storeFactory()
	_, err := store.ReadPuppy("4")
	s.Error(err)
}

func (s *storerSuite) TestUpdatePuppyHappyCase() {
	store := s.storeFactory()
	_ = store.CreatePuppy(Puppy{ID: "5", Colour: "Brown"})
	err := store.UpdatePuppy(Puppy{ID: "5", Colour: "Green"})
	require.NoError(s.T(), err)
	p, _ := store.ReadPuppy("5")
	s.Equal("Green", p.Colour)
}

func (s *storerSuite) TestUpdatePuppyNonExisting() {
	store := s.storeFactory()
	err := store.UpdatePuppy(Puppy{ID: "6"})
	s.Error(err)
}

func (s *storerSuite) TestDeletePuppyHappyCase() {
	store := s.storeFactory()
	_ = store.CreatePuppy(Puppy{ID: "7", Colour: "Brown"})
	deleted, err := store.DeletePuppy("7")
	require.NoError(s.T(), err)
	require.Equal(s.T(), true, deleted)

	_, err = store.ReadPuppy("7")
	s.Error(err, "Puppy gone after deletion")
}

func (s *storerSuite) TestDeletePuppyNonExisting() {
	store := s.storeFactory()
	deleted, err := store.DeletePuppy("8")
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
