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
	maps  storerImpl = iota
	smaps storerImpl = iota
)

var (
	puppy1 = func() Puppy {
		return Puppy{
			ID:     0x1234,
			Breed:  "Sheep herder",
			Colour: "Brown",
			Value:  1000,
		}
	}
	puppy2 = func() Puppy {
		return Puppy{
			ID:     0x1236,
			Breed:  "Sheep herder",
			Colour: "Brown",
			Value:  1000,
		}
	}
)

type storerSuite struct {
	suite.Suite
	store Storer
	impl  storerImpl
}

func (s *storerSuite) SetupTest() {
	switch s.impl {
	case maps:
		// create a map store
		s.store = CreateStore()
		puppy := puppy1()
		s.store.CreatePuppy(&puppy)
	case smaps:
		// create a sync map store
		s.store = CreateSyncStore()
		puppy := puppy1()
		s.store.CreatePuppy(&puppy)
	default:
		panic("Unrecognised storer implementation")
	}
}

func (s *storerSuite) TestMapStoreReadPuppySuccessful() {
	// given
	assert := tassert.New(s.T())

	// when
	actual := *s.store.ReadPuppy(0x1234)

	// then
	expected := puppy1()
	assert.Equal(expected, actual)
}

func (s *storerSuite) TestMapStoreReadPuppy_IDDoesNotExist() {
	// given
	assert := tassert.New(s.T())

	// when
	actual := *s.store.ReadPuppy(0x123)

	// then
	assert.Equal(Puppy{0, "", "", 0}, actual)
}

func (s *storerSuite) TestMapStoreCreatePuppySuccessful() {
	// given
	assert := tassert.New(s.T())
	newPup := puppy2()

	// when
	s.store.CreatePuppy(&newPup)

	// then
	createdPup := *s.store.ReadPuppy(0x1236)
	assert.Equal(newPup, createdPup, "Created puppy should be identical to the one passed to Create")
}

func (s *storerSuite) TestMapStoreUpdatePuppySuccessful() {
	// given
	assert := tassert.New(s.T())
	newPup := puppy1()

	// when
	s.store.CreatePuppy(&newPup)
	createdPup := *s.store.ReadPuppy(0x1236)
	createdPup.Value = 12323
	s.store.UpdatePuppy(createdPup.ID, &createdPup)

	// then
	updatedPup := *s.store.ReadPuppy(createdPup.ID)
	assert.Equal(createdPup, updatedPup, "Updated puppy should be identical to the one passed to Update")
}

func (s *storerSuite) TestMapStoreDeletePuppySuccessful() {
	//given
	assert := tassert.New(s.T())

	// when
	deleted := s.store.DeletePuppy(0x1234)

	// then
	assert.True(deleted, "Delete should return true indicating a puppy was deleted")
	actual := *s.store.ReadPuppy(0x1234)
	assert.Equal(Puppy{0, "", "", 0}, actual)
}

func TestStorerImpls(t *testing.T) {
	mapSuite := storerSuite{
		store: CreateStore(),
		impl:  maps,
	}
	suite.Run(t, &mapSuite)

	syncMapSuite := storerSuite{
		store: CreateSyncStore(),
		impl:  smaps,
	}
	suite.Run(t, &syncMapSuite)
}

func TestLettersMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "{11 Sheep herder Brown 1000}\n{11 Sheep herder Brown 10000}\n{0   0}\n" +
		"{11 Sheep herder Brown 1000}\n{11 Sheep herder Brown 10000}\n{0   0}\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
