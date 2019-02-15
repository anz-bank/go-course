package main

import (
	"bytes"
	"testing"

	tassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type syncStorerImpl int
type mapStorerImpl int

var mapID, mapIDm uint

const (
	smem syncStorerImpl = iota
	mmem mapStorerImpl  = iota
)

var (
	puppyIN = func() Puppy {
		return Puppy{
			id:    107,
			breed: "Slinky",
			color: "purple",
			value: "12999",
		}
	}
	puppyOut1 = func() Puppy {
		return Puppy{
			id:    1,
			breed: "Slinky",
			color: "purple",
			value: "12999",
		}
	}
	puppyOut2 = func() Puppy {
		return Puppy{
			id:    2,
			breed: "Slinky",
			color: "purple",
			value: "12999",
		}
	}
	modifiedpuppyIN = func() Puppy {
		return Puppy{
			id:    107,
			breed: "Labrador Retriever",
			color: "brown",
			value: "13999",
		}
	}
	puppy108 = func() Puppy {
		return Puppy{
			id:    108,
			breed: "Bulldog",
			color: "red",
			value: "842.50",
		}
	}
	puppyEmpty = func() Puppy {
		return Puppy{}
	}
)

type storerSuite struct {
	suite.Suite
	syncstore, mapstore Storer
	simpl               syncStorerImpl
	mimpl               mapStorerImpl
}

func (s *storerSuite) SetupTest() {
	switch s.simpl {
	case smem:
		// create a sync store
		s.syncstore = newSyncStore()
		syncpup := puppyIN()
		mapID = s.syncstore.createPuppy(syncpup)
	default:
		panic("Unrecognised storer implementation")
	}
	switch s.mimpl {
	case mmem:
		// create a map store
		s.mapstore = newMapStore()
		mappup := puppyIN()
		mapIDm = s.mapstore.createPuppy(mappup)
	default:
		panic("Unrecognised storer implementation")
	}
}

// Sync map based test suit
func (s *storerSuite) TestReadPuppySuccessful() {
	// given
	assert := tassert.New(s.T())

	// when
	syncPuppy := s.syncstore.readPuppy(mapID)

	// then
	expectedSyncPuppy := puppyIN()
	assert.Equal(expectedSyncPuppy, syncPuppy, "store should return a puppy identical to puppyIN")
}

func (s *storerSuite) TestReadPuppy_IDDoesNotExist() {
	// given
	assert := tassert.New(s.T())

	// when
	syncPuppy := s.syncstore.readPuppy(28)

	// then
	assert.Equal(puppyEmpty(), syncPuppy, "should get an empty puppy when attempting to read an non-existing puppy")
}

func (s *storerSuite) TestCreatePuppySuccessful() {
	// given
	assert := tassert.New(s.T())
	puppyIN := puppyIN()

	// when
	mapID = s.syncstore.createPuppy(puppyIN)

	// then
	createdSyncPuppy := s.syncstore.readPuppy(mapID)
	assert.Equal(puppyIN, createdSyncPuppy, "Created puppy should be identical to the one passed to Create")
}

func (s *storerSuite) TestUpdatePuppySuccessful() {
	// given
	assert := tassert.New(s.T())
	testModifiedPuppy := modifiedpuppyIN()

	// when
	s.syncstore.updatePuppy(107, testModifiedPuppy)

	// then
	storedSyncPuppy := s.syncstore.readPuppy(107)
	assert.Equal(testModifiedPuppy, storedSyncPuppy, "Stored pet should be equal to the modified pet")
}

func (s *storerSuite) TestDeletePuppySuccessful() {
	// given
	assert := tassert.New(s.T())

	// when
	s.syncstore.deletePuppy(107)

	// then
	storedSyncPuppy := s.syncstore.readPuppy(107)
	assert.Equal(puppyEmpty(), storedSyncPuppy, "Stored pet should be equal to the modified pet")
}

// Map based test suit
func (s *storerSuite) TestMapReadPuppySuccessful() {
	// given
	assert := tassert.New(s.T())

	// when
	mapPuppy := s.mapstore.readPuppy(mapIDm)

	// then
	expectedMapPuppy := puppyOut1()
	assert.Equal(expectedMapPuppy, mapPuppy, "store should return a puppy identical to puppyIN")
}

func (s *storerSuite) TestMapReadPuppy_IDDoesNotExist() {
	// given
	assert := tassert.New(s.T())

	// when
	mapPuppy := s.mapstore.readPuppy(117)

	// then
	assert.Equal(puppyEmpty(), mapPuppy, "should get an empty puppy when attempting to read an non-existing puppy")
}

func (s *storerSuite) TestMapCreatePuppySuccessful() {
	// given
	assert := tassert.New(s.T())
	newPuppy := puppy108()

	// when
	s.syncstore.createPuppy(newPuppy)
	mapIDm = s.mapstore.createPuppy(puppyIN())

	// then
	createdMapPuppy := s.mapstore.readPuppy(mapIDm)
	assert.Equal(puppyOut2(), createdMapPuppy, "Created puppy should be identical to the one passed to Create")
}

func (s *storerSuite) TestMapUpdatePuppySuccessful() {
	// given
	assert := tassert.New(s.T())
	testModifiedPuppy := modifiedpuppyIN()

	// when
	s.mapstore.updatePuppy(107, testModifiedPuppy)

	// then
	storedMapPuppy := s.mapstore.readPuppy(107)
	assert.Equal(testModifiedPuppy, storedMapPuppy, "Stored pet should be equal to the modified pet")
}

func (s *storerSuite) TestMapDeletePuppySuccessful() {
	// given
	assert := tassert.New(s.T())

	// when
	s.mapstore.deletePuppy(107)

	// then
	storedMapPuppy := s.mapstore.readPuppy(107)
	assert.Equal(puppyEmpty(), storedMapPuppy, "Stored pet should be equal to the modified pet")
}

func TestStorer(t *testing.T) {
	memSuite := storerSuite{
		syncstore: newSyncStore(),
		mapstore:  newMapStore(),
		simpl:     smem,
		mimpl:     mmem,
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
102 :  {0   }
103 :  {103 German Shepherd red 4533}
~~~~~~~~~
Map Store
~~~~~~~~~
1  :  {1 Pug brown 0.91}
2  :  {0   }
3  :  {3 Beagle brown 0.91}` + "\n"
	r.Equalf(expected, buf.String(), "Unexpected output in main()")
}
