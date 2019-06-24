package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type storerSuite struct {
	suite.Suite
	store Storer
}

func (s *storerSuite) TestCreatePuppyHappyCase() {
	t := s.T()
	err := s.store.CreatePuppy(Puppy{ID: "1", Colour: "Black"})
	assert.NoError(s.T(), err, "Happy case puppy creation")

	p, _ := s.store.ReadPuppy("1")
	assert.Equal(t, "Black", p.Colour)
}

func (s *storerSuite) TestCreatePuppyAlreadyExists() {
	_ = s.store.CreatePuppy(Puppy{ID: "2"})
	err := s.store.CreatePuppy(Puppy{ID: "2", Colour: "Red"})
	assert.Error(s.T(), err, "Puppy creation fails if ID already exists")
}

func (s *storerSuite) TestReadPuppyHappyCase() {
	t := s.T()
	_ = s.store.CreatePuppy(Puppy{ID: "3", Colour: "Blue"})
	p, err := s.store.ReadPuppy("3")
	require.NoError(t, err)
	assert.Equal(t, "Blue", p.Colour)
}

func (s *storerSuite) TestReadPuppyNonExisting() {
	t := s.T()
	_, err := s.store.ReadPuppy("4")
	assert.Error(t, err)
}

func (s *storerSuite) TestUpdatePuppyHappyCase() {
	t := s.T()
	_ = s.store.CreatePuppy(Puppy{ID: "5", Colour: "Brown"})
	err := s.store.UpdatePuppy(Puppy{ID: "5", Colour: "Green"})
	require.NoError(t, err)
	p, _ := s.store.ReadPuppy("5")
	require.Equal(t, "Green", p.Colour)
}

func (s *storerSuite) TestUpdatePuppyNonExisting() {
	t := s.T()
	err := s.store.UpdatePuppy(Puppy{ID: "6"})
	assert.Error(t, err)
}

func (s *storerSuite) TestDeletePuppyHappyCase() {
	t := s.T()
	_ = s.store.CreatePuppy(Puppy{ID: "7", Colour: "Brown"})
	deleted, err := s.store.DeletePuppy("7")
	require.NoError(t, err)
	require.Equal(t, true, deleted)

	_, err = s.store.ReadPuppy("7")
	assert.Error(t, err, "Puppy gone after deletion")
}

func (s *storerSuite) TestDeletePuppyNonExisting() {
	t := s.T()
	deleted, err := s.store.DeletePuppy("8")
	require.Error(t, err)
	assert.Equal(t, false, deleted)
}

func TestRest(t *testing.T) {
	suite.Run(t, &storerSuite{store: NewMapStore()})
	suite.Run(t, &storerSuite{store: NewSyncStore()})
}
