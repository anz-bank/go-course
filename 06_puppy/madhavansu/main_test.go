package main

import (
	"bytes"
	"testing"

	tassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type storerImpl int

var mapID uint

const (
	smem storerImpl = iota
	mmem storerImpl = iota
)

var (
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
	puppyEmpty = func() Puppy {
		return Puppy{}
	}
)

type storerSuite struct {
	suite.Suite
	storer Storer
	impl   storerImpl
}

func (s *storerSuite) SetupTest() {
	switch s.impl {
	case smem:
		// create a sync store
		s.storer = newSyncStore()
		syncpup := puppyOut1()
		mapID = s.storer.createPuppy(syncpup)
	case mmem:
		// create a sync store
		s.storer = newMapStore()
		syncpup := puppyOut1()
		mapID = s.storer.createPuppy(syncpup)
	default:
		panic("Unrecognised storer implementation")
	}
}

func TestStorer(t *testing.T) {
	memSuite := storerSuite{impl: smem}
	suite.Run(t, &memSuite)
	smemSuite := storerSuite{impl: mmem}
	suite.Run(t, &smemSuite)
}

// Test suite
func (s *storerSuite) TestReadPuppySuccessful() {
	// given
	assert := tassert.New(s.T())

	// when
	syncPuppy := s.storer.readPuppy(mapID)

	// then
	expectedSyncPuppy := puppyOut1()
	assert.Equal(expectedSyncPuppy, syncPuppy, "store should return a puppy identical to puppyIN")
}

func (s *storerSuite) TestReadPuppy_IDDoesNotExist() {
	// given
	assert := tassert.New(s.T())

	// when
	syncPuppy := s.storer.readPuppy(300)

	// then
	assert.Equal(puppyEmpty(), syncPuppy, "should get an empty puppy when attempting to read an non-existing puppy")
}

func (s *storerSuite) TestCreatePuppySuccessful() {
	// given
	assert := tassert.New(s.T())
	puppyIN := puppyOut2()

	// when
	mapID = s.storer.createPuppy(puppyIN)

	// then
	createdSyncPuppy := s.storer.readPuppy(mapID)
	assert.Equal(puppyIN, createdSyncPuppy, "Created puppy should be identical to the one passed to Create")
}

func (s *storerSuite) TestUpdatePuppySuccessful() {
	// given
	assert := tassert.New(s.T())
	testModifiedPuppy := puppyOut2()

	// when
	s.storer.updatePuppy(2, testModifiedPuppy)

	// then
	storedSyncPuppy := s.storer.readPuppy(2)
	assert.Equal(testModifiedPuppy, storedSyncPuppy, "Stored pet should be equal to the modified pet")
}

func (s *storerSuite) TestDeletePuppySuccessful() {
	// given
	assert := tassert.New(s.T())

	// when
	s.storer.deletePuppy(2)

	// then
	storedSyncPuppy := s.storer.readPuppy(2)
	assert.Equal(puppyEmpty(), storedSyncPuppy, "Stored pet should be equal to the modified pet")
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
