package main

import (
	"bytes"
	"testing"

	tassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type storerImpl int

const (
	mem storerImpl = iota
)

var (
	puppy107 = func() Puppy {
		return Puppy{
			id:    107,
			breed: "Slinky",
			color: "purple",
			value: 12999,
		}
	}
	modifiedPuppy107 = func() Puppy {
		return Puppy{
			id:    107,
			breed: "Labrador Retriever",
			color: "brown",
			value: 13999,
		}
	}
	puppy108 = func() Puppy {
		return Puppy{
			id:    108,
			breed: "Bulldog",
			color: "red",
			value: 842.50,
		}
	}
	puppyEmpty = func() Puppy {
		return Puppy{
			id:    0,
			breed: "",
			color: "",
			value: 0,
		}
	}
)

type storerSuite struct {
	suite.Suite
	syncstore, mapstore Storer
	impl                storerImpl
}

func (s *storerSuite) SetupTest() {
	switch s.impl {
	case mem:
		// create a sync store
		s.syncstore = newSyncStore()
		syncpup := puppy107()
		s.syncstore.createPuppy(syncpup)
		// create a map store
		s.mapstore = newMapStore()
		mappup := puppy107()
		s.mapstore.createPuppy(mappup)
	default:
		panic("Unrecognised storer implementation")
	}
}

func (s *storerSuite) TestReadPuppySuccessful() {
	// given
	assert := tassert.New(s.T())

	// when
	syncPuppy := s.syncstore.readPuppy(107)
	mapPuppy := s.mapstore.readPuppy(107)

	// then
	expectedpuppy := puppy107()
	assert.Equal(expectedpuppy, syncPuppy, "store should return a puppy identical to puppy107")
	assert.Equal(expectedpuppy, mapPuppy, "store should return a puppy identical to puppy107")
}

func (s *storerSuite) TestReadPuppy_IDDoesNotExist() {
	// given
	assert := tassert.New(s.T())

	// when
	syncPuppy := s.syncstore.readPuppy(117)
	mapPuppy := s.mapstore.readPuppy(117)

	// then
	assert.Equal(puppyEmpty(), syncPuppy, "should get an empty puppy when attempting to read an non-existing puppy")
	assert.Equal(puppyEmpty(), mapPuppy, "should get an empty puppy when attempting to read an non-existing puppy")
}

func (s *storerSuite) TestCreatePuppySuccessful() {
	// given
	assert := tassert.New(s.T())
	newPuppy := puppy108()

	// when
	s.syncstore.createPuppy(newPuppy)
	s.mapstore.createPuppy(newPuppy)

	// then
	createdSyncPuppy := s.syncstore.readPuppy(108)
	createdMapPuppy := s.mapstore.readPuppy(108)
	assert.Equal(newPuppy, createdSyncPuppy, "Created puppy should be identical to the one passed to Create")
	assert.Equal(newPuppy, createdMapPuppy, "Created puppy should be identical to the one passed to Create")
}

func (s *storerSuite) TestUpdatePuppySuccessful() {
	// given
	assert := tassert.New(s.T())
	testModifiedPuppy := modifiedPuppy107()

	// when
	s.syncstore.updatePuppy(107, testModifiedPuppy)
	s.mapstore.updatePuppy(107, testModifiedPuppy)

	// then
	storedSyncPuppy := s.syncstore.readPuppy(107)
	storedMapPuppy := s.mapstore.readPuppy(107)
	assert.Equal(testModifiedPuppy, storedSyncPuppy, "Stored pet should be equal to the modified pet")
	assert.Equal(testModifiedPuppy, storedMapPuppy, "Stored pet should be equal to the modified pet")
}

func (s *storerSuite) TestDeletePuppySuccessful() {
	// given
	assert := tassert.New(s.T())

	// when
	s.syncstore.deletePuppy(107)
	s.mapstore.deletePuppy(107)

	// then
	storedSyncPuppy := s.syncstore.readPuppy(107)
	storedMapPuppy := s.mapstore.readPuppy(107)
	assert.Equal(puppyEmpty(), storedSyncPuppy, "Stored pet should be equal to the modified pet")
	assert.Equal(puppyEmpty(), storedMapPuppy, "Stored pet should be equal to the modified pet")
}

func TestStorer(t *testing.T) {
	memSuite := storerSuite{
		syncstore: newSyncStore(),
		mapstore:  newMapStore(),
		impl:      mem,
	}
	suite.Run(t, &memSuite)
}

// Test main output
func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := `~~~~~~~~~~
Sync Store
~~~~~~~~~~
101 :  {101 Poodle red 18000}
102 :  {0   0}
103 :  {103 German Shepherd red 4533}
~~~~~~~~~
Map Store
~~~~~~~~~
104 :  {104 Pug brown 0.91}
105 :  {0   0}
106 :  {106 Beagle brown 0.91}`
	r.Equalf(expected, buf.String(), "Unexpected output in main()")
}
