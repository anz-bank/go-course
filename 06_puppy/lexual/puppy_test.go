package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type storerSuite struct {
	suite.Suite
	storer Storer
}

func (s *storerSuite) TestCreateReadPuppy() {
	r := assert.New(s.T())
	p1 := Puppy{7, "bulldog", "black", 100}
	s.storer.CreatePuppy(&p1)
	p2 := s.storer.ReadPuppy(7)
	r.Equal(p1, *p2)
}

func (s *storerSuite) TestUpdatePuppy() {
	r := assert.New(s.T())
	p1 := Puppy{7, "bulldog", "black", 100}
	p2 := Puppy{7, "bulldog", "white", 100}
	s.storer.CreatePuppy(&p1)
	s.storer.UpdatePuppy(7, &p2)
	p3 := s.storer.ReadPuppy(7)
	r.NotEqual(p1, *p3)
	r.NotEqual(p1, p2)
	r.Equal(p2, *p3)
}

func (s *storerSuite) TestDeletePuppy() {
	r := assert.New(s.T())

	p1 := Puppy{7, "bulldog", "black", 100}
	s.storer.CreatePuppy(&p1)
	p2 := s.storer.ReadPuppy(7)
	r.Equal(p1, *p2)

	didDelete := s.storer.DeletePuppy(7)
	r.True(didDelete)
	p3 := s.storer.ReadPuppy(7)
	r.Nil(p3)

	didDelete = s.storer.DeletePuppy(7)
	r.False(didDelete)
}

func TestStorer(t *testing.T) {
	// ms := make(MapStore)
	// ms := MapStore{}
	suite.Run(t, &storerSuite{storer: &MapStore{}})
	ss := SyncStore{}
	suite.Run(t, &storerSuite{storer: &ss})
}
