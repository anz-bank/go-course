package main

import (
	"bytes"
	"testing"

	assert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type storerSuite struct {
	suite.Suite
	storer Storer
}

func TestStorer(t *testing.T) {
	suite.Run(t, &storerSuite{storer: NewMapStore()})
	suite.Run(t, &storerSuite{storer: NewSyncStore()})
}

func (s *storerSuite) TestCreateWithInvaliID() {
	// Given
	r := assert.New(s.T())
	invalidPuppy := &Puppy{Breed: "Beagle", Color: "Brown", Value: 132.23}

	// When
	err := s.storer.Create(invalidPuppy)

	// Then
	r.EqualError(err, "The id should be a positive number. [Error Code: 1]", "unexpected output")
}

func (s *storerSuite) TestCreateWhenIDExists() {
	// Given
	r := assert.New(s.T())
	puppy := &Puppy{2, "Beagle", "Brown", 132.23}
	puppyConflict := &Puppy{2, "Boxer", "Black", 2000}
	if err := s.storer.Create(puppy); err != nil {
		r.FailNow("Prepare data error")
	}

	// When
	err := s.storer.Create(puppyConflict)

	// Then
	r.EqualError(err, "The Puppy with ID `2` already exists. [Error Code: 2]", "unexpected output")
}

func (s *storerSuite) TestCreateSuccess() {
	// Given
	r := assert.New(s.T())
	puppy := &Puppy{3, "Beagle", "Brown", 132.23}

	// When
	err := s.storer.Create(puppy)

	// Then
	r.True(err == nil, "unexpected output")
}

func (s *storerSuite) TestReadWithInvalidID() {
	r := assert.New(s.T())

	// When
	read, err := s.storer.Read(0)

	// Then
	r.True(read == nil, "unexpected output")
	r.EqualError(err, "The id should be a positive number. [Error Code: 1]", "unexpected output")
}

func (s *storerSuite) TestReadNotFound() {
	r := assert.New(s.T())

	// When
	read, err := s.storer.Read(4)

	// Then
	r.True(read == nil, "unexpected output")
	r.EqualError(err, "The puppy with ID `4` does not exist. [Error Code: 4]", "unexpected output")
}

func (s *storerSuite) TestReadSuccess() {
	// Given
	r := assert.New(s.T())
	puppy := &Puppy{5, "Beagle", "White", 612}
	if err := s.storer.Create(puppy); err != nil {
		r.FailNow("Prepare data error")
	}

	// When
	read, err := s.storer.Read(5)

	// Then
	r.True(err == nil, "unexpected output")
	r.Equalf(puppy, read, "unexpected output")
}

func (s *storerSuite) TestUpdateWithInvalidID() {
	// Given
	r := assert.New(s.T())
	puppy := &Puppy{6, "Beagle", "White", 612}

	// When
	err := s.storer.Update(0, puppy)

	// Then
	r.EqualError(err, "The id should be a positive number. [Error Code: 1]", "unexpected output")
}

func (s *storerSuite) TestUpdateWithPuppyContainsInvalidID() {
	// Given
	r := assert.New(s.T())
	puppy := &Puppy{0, "Beagle", "White", 612}

	// When
	err := s.storer.Update(6, puppy)

	// Then
	r.EqualError(err,
		"The input Puppy[ID: 0, Breed: Beagle, Color: White, Value: 612.00] has invalid ID `0`. [Error Code: 1]",
		"unexpected output")
}

func (s *storerSuite) TestUpdateWithInconsistentID() {
	// Given
	r := assert.New(s.T())
	puppy := &Puppy{7, "Beagle", "White", 612}

	// When
	err := s.storer.Update(6, puppy)

	// Then
	r.EqualError(err, "The ID mismatch; The given id is `6` but the puppy.ID is `7`. [Error Code: 1]", "unexpected output")
}

func (s *storerSuite) TestUpdateNotFound() {
	// Given
	r := assert.New(s.T())
	puppy := &Puppy{6, "Beagle", "White", 612}

	// When
	err := s.storer.Update(6, puppy)

	// Then
	r.EqualError(err, "The puppy with ID `6` does not exist. [Error Code: 4]", "unexpected output")
}

func (s *storerSuite) TestUpdateSuccess() {
	// Given
	r := assert.New(s.T())
	puppy := &Puppy{7, "Beagle", "White", 612}
	updatePuppy := &Puppy{7, "Beagle", "White", 1000}
	if err := s.storer.Create(puppy); err != nil {
		r.FailNow("Prepare data failed")
	}

	// When
	err := s.storer.Update(7, updatePuppy)

	// Then
	r.True(err == nil, "unexpected output")
}

func (s *storerSuite) TestDeleteWithInvalidID() {
	r := assert.New(s.T())

	// When
	err := s.storer.Delete(0)

	// Then
	r.EqualError(err, "The id should be a positive number. [Error Code: 1]", "unexpected output")
}

func (s *storerSuite) TestDeleteNotFound() {
	r := assert.New(s.T())

	// When
	err := s.storer.Delete(8)

	// Then
	r.EqualError(err, "The puppy with ID `8` does not exist. [Error Code: 4]", "unexpected output")
}

func (s *storerSuite) TestDeleteSuccess() {
	// Given
	r := assert.New(s.T())
	puppy := &Puppy{9, "Beagle", "White", 612}
	if err := s.storer.Create(puppy); err != nil {
		r.FailNow("Prepare data failed")
	}

	// When
	err := s.storer.Delete(9)

	// Then
	r.True(err == nil, "unexpected output")
}

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := `Map Store
Create Puppy[ID: 1, Breed: Beagle, Color: Brown, Value: 132.23] succeed
Read Puppy[ID: 1, Breed: Beagle, Color: Brown, Value: 132.23] with ID 1 succeed
Update Puppy[ID: 1, Breed: Boxer, Color: Black, Value: 2000.00] succeed
Read Puppy[ID: 1, Breed: Boxer, Color: Black, Value: 2000.00] with ID 1 succeed
Delete Puppy with ID 1 succeed
SyncMap Store
Create Puppy[ID: 1, Breed: Beagle, Color: Brown, Value: 132.23] succeed
Read Puppy[ID: 1, Breed: Beagle, Color: Brown, Value: 132.23] with ID 1 succeed
Update Puppy[ID: 1, Breed: Boxer, Color: Black, Value: 2000.00] succeed
Read Puppy[ID: 1, Breed: Boxer, Color: Black, Value: 2000.00] with ID 1 succeed
Delete Puppy with ID 1 succeed
`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
