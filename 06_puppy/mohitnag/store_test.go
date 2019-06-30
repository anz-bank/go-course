package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type storer struct {
	suite.Suite
	store      Storer
	storerType func() Storer
}

func (s *storer) SetupTest() {
	s.store = s.storerType()
}

func (s *storer) TestCreatePuppy() {
	assert := assert.New(s.T())
	puppy := Puppy{ID: 1, Breed: "dog", Colour: "white", Value: "$2"}
	testCases := []struct {
		Scenario      string
		Input         Puppy
		ExpectedError string
	}{
		{"Create Puppy", puppy, "<nil>"},
		{"Creating already existing Puppy should fail", puppy, "puppy with Id 1 already exists"},
	}
	for _, tc := range testCases {
		tc := tc
		s.T().Run(tc.Scenario, func(t *testing.T) {
			err := s.store.CreatePuppy(tc.Input)
			assert.Equal(tc.ExpectedError, fmt.Sprint(err))

		})
	}
}

func (s *storer) TestReadPuppy() {
	assert := assert.New(s.T())
	puppy := Puppy{ID: 1, Breed: "dog", Colour: "white", Value: "$2"}
	err := s.store.CreatePuppy(puppy)
	assert.NoError(err)
	testCases := []struct {
		Scenario      string
		ID            uint32
		Expected      Puppy
		ExpectedError string
	}{
		{"Read already existing Puppy", 1, puppy, "<nil>"},
		{"Reading a non-existing Puppy should fail", 32, Puppy{}, "puppy with Id 32 does not exists"},
	}
	for _, tc := range testCases {
		tc := tc
		s.T().Run(tc.Scenario, func(t *testing.T) {
			puppy, err := s.store.ReadPuppy(tc.ID)
			assert.Equal(tc.Expected, puppy)
			assert.Equal(tc.ExpectedError, fmt.Sprint(err))
		})
	}

}

func (s *storer) TestUpdatePuppy() {
	assert := assert.New(s.T())
	puppy := Puppy{ID: 1, Breed: "dog", Colour: "white", Value: "$2"}
	err := s.store.CreatePuppy(puppy)
	assert.NoError(err)
	puppyUpdate := Puppy{ID: 1, Breed: "dog", Colour: "black", Value: "$2"}
	puppyNonExist := Puppy{ID: 32, Breed: "dog", Colour: "black", Value: "$2"}

	testCases := []struct {
		Scenario      string
		Puppy         Puppy
		ExpectedError string
	}{
		{"Update already existing Puppy", puppyUpdate, "<nil>"},
		{"Update a non-existing Puppy should fail", puppyNonExist, "puppy with Id 32 does not exists"},
	}
	for _, tc := range testCases {
		tc := tc
		s.T().Run(tc.Scenario, func(t *testing.T) {
			err := s.store.UpdatePuppy(tc.Puppy)
			if err == nil {
				puppy, _ := s.store.ReadPuppy(tc.Puppy.ID)
				assert.Equal("black", puppy.Colour)
			}
			assert.Equal(tc.ExpectedError, fmt.Sprint(err))
		})
	}

}

func (s *storer) TestDeletePuppy() {
	assert := assert.New(s.T())
	puppy := Puppy{ID: 1, Breed: "dog", Colour: "white", Value: "$2"}
	err := s.store.CreatePuppy(puppy)
	assert.NoError(err)
	testCases := []struct {
		Scenario string
		ID       uint32
		Expected bool
	}{
		{"Delete already existing Puppy", 1, true},
		{"Delete a non-existing Puppy should fail", 32, false},
	}
	for _, tc := range testCases {
		tc := tc
		s.T().Run(tc.Scenario, func(t *testing.T) {
			isDeleted := s.store.DeletePuppy(tc.ID)
			assert.Equal(tc.Expected, isDeleted)
		})
	}
}

func TestStorers(t *testing.T) {
	suite.Run(t, &storer{
		storerType: func() Storer { return &MemStore{} },
	})
	suite.Run(t, &storer{
		storerType: func() Storer { return &SyncStore{} },
	})
}
