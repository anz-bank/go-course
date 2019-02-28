package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	assert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	_ = iota
	MAP
	SYNCMAP
	LEVELDB
)

var dbPath string

type storerSuite struct {
	suite.Suite
	storer     Storer
	storerType int
	initPuppy  *Puppy
}

func TestStorer(t *testing.T) {
	suite.Run(t, &storerSuite{storerType: MAP})
	suite.Run(t, &storerSuite{storerType: SYNCMAP})
	suite.Run(t, &storerSuite{storerType: LEVELDB})
}

func (s *storerSuite) SetupTest() {
	dbPath = getDbPath()
	switch {
	case s.storerType == MAP:
		s.storer = NewMapStore()
	case s.storerType == SYNCMAP:
		s.storer = NewSyncStore()
	case s.storerType == LEVELDB:
		s.storer = NewLevelDbStore(dbPath)
	default:
		panic("Unknow storer implementation")
	}

	s.initPuppy = &Puppy{1, "Beagle", "Brown", 132.23}
	if err := s.storer.Create(s.initPuppy); err != nil {
		panic("Setup test data failed")
	}
}

func (s *storerSuite) TearDownTest() {
	deleteTestDb(dbPath)
	s.initPuppy = nil
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
	puppy := &Puppy{1, "Beagle", "Black", 132.23}

	// When
	err := s.storer.Create(puppy)

	// Then
	r.EqualError(err, "The Puppy with ID `1` already exists. [Error Code: 2]", "unexpected output")
}

func (s *storerSuite) TestCreateSuccess() {
	// Given
	r := assert.New(s.T())
	puppy := &Puppy{2, "Beagle", "Brown", 1800}

	// When
	err := s.storer.Create(puppy)

	// Then
	r.Nil(err, "unexpected output")
}

func (s *storerSuite) TestReadWithInvalidID() {
	r := assert.New(s.T())

	// When
	read, err := s.storer.Read(0)

	// Then
	r.Nil(read, "unexpected output")
	r.EqualError(err, "The id should be a positive number. [Error Code: 1]", "unexpected output")
}

func (s *storerSuite) TestReadNotFound() {
	r := assert.New(s.T())

	// When
	read, err := s.storer.Read(2)

	// Then
	r.Nil(read, "unexpected output")
	r.EqualError(err, "The puppy with ID `2` does not exist. [Error Code: 4]", "unexpected output")
}

func (s *storerSuite) TestReadSuccess() {
	r := assert.New(s.T())

	// When
	read, err := s.storer.Read(1)

	// Then
	r.Nil(err, "unexpected output")
	r.Equalf(s.initPuppy, read, "unexpected output")
}

func (s *storerSuite) TestUpdateWithInvalidID() {
	// Given
	r := assert.New(s.T())
	puppy := &Puppy{1, "Beagle", "White", 612}

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
	err := s.storer.Update(1, puppy)

	// Then
	r.EqualError(err,
		"The input Puppy[ID: 0, Breed: Beagle, Color: White, Value: 612.00] has invalid ID `0`. [Error Code: 1]",
		"unexpected output")
}

func (s *storerSuite) TestUpdateWithInconsistentID() {
	// Given
	r := assert.New(s.T())
	puppy := &Puppy{1, "Beagle", "White", 612}

	// When
	err := s.storer.Update(2, puppy)

	// Then
	r.EqualError(err, "The ID mismatch; The given id is `2` but the puppy.ID is `1`. [Error Code: 1]", "unexpected output")
}

func (s *storerSuite) TestUpdateNotFound() {
	// Given
	r := assert.New(s.T())
	puppy := &Puppy{2, "Beagle", "White", 612}

	// When
	err := s.storer.Update(2, puppy)

	// Then
	r.EqualError(err, "The puppy with ID `2` does not exist. [Error Code: 4]", "unexpected output")
}

func (s *storerSuite) TestUpdateSuccess() {
	// Given
	r := assert.New(s.T())
	puppy := &Puppy{1, "Beagle", "White", 612}

	// When
	err := s.storer.Update(1, puppy)

	// Then
	r.Nil(err, "unexpected output")
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
	r := assert.New(s.T())

	// When
	err := s.storer.Delete(1)

	// Then
	r.Nil(err, "unexpected output")
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
levelDB Store
Create Puppy[ID: 1, Breed: Beagle, Color: Brown, Value: 132.23] succeed
Read Puppy[ID: 1, Breed: Beagle, Color: Brown, Value: 132.23] with ID 1 succeed
Update Puppy[ID: 1, Breed: Boxer, Color: Black, Value: 2000.00] succeed
Read Puppy[ID: 1, Breed: Boxer, Color: Black, Value: 2000.00] with ID 1 succeed
Delete Puppy with ID 1 succeed
`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func getDbPath() string {
	p := filepath.Join(os.TempDir(), "leveldb-puppy")
	deleteTestDb(p)

	return p
}

func deleteTestDb(path string) {
	if err := os.RemoveAll(path); err != nil {
		panic("Fail to clean the envrionment")
	}
}
