package store

import (
	"os"
	"testing"

	types "github.com/anz-bank/go-training/08_project/mohankrishna/pkg/mohankrishna-puppy"
	tassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type storerImpl int

const (
	maps storerImpl = iota
	smaps
	levelDB
)

var (
	puppy1 = func() types.Puppy {
		return types.Puppy{
			ID:     0x1234,
			Breed:  "Sheep herder",
			Colour: "Brown",
			Value:  1000,
		}
	}
	puppy2 = func() types.Puppy {
		return types.Puppy{
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

func (s *storerSuite) SetupSuite() {
	switch s.impl {
	case maps:
		// create a map store
		s.store = NewMapStore()
	case smaps:
		// create a sync map store
		s.store = NewSyncStore()
	case levelDB:
		// create a level db store
		s.store = NewLevelDBStore(os.TempDir() + "/level_store")
	default:
		panic("Unrecognised storer implementation")
	}
	puppy := puppy1()
	err := s.store.CreatePuppy(&puppy)
	if err != nil {
		panic("Could not initialise tests")
	}
}

func (s *storerSuite) TearDownSuite() {
	os.RemoveAll(os.TempDir() + "/level_store")
}

func TestStorerImpls(t *testing.T) {
	mapSuite := storerSuite{impl: maps}
	suite.Run(t, &mapSuite)

	syncMapSuite := storerSuite{impl: smaps}
	suite.Run(t, &syncMapSuite)

	levelDBSuite := storerSuite{impl: levelDB}
	suite.Run(t, &levelDBSuite)
}

func (s *storerSuite) TestMapStoreReadPuppySuccessful() {
	// given
	assert := tassert.New(s.T())
	pup1 := puppy1()
	err := s.store.CreatePuppy(&pup1)
	if err != nil {
		panic("Could not initialise test data")
	}

	// when
	actual, err := s.store.ReadPuppy(0x1234)

	// then
	expected := puppy1()
	if assert.NoError(err, "Should be able to read puppy from store") {
		assert.Equal(expected, *actual)
	}
}

func (s *storerSuite) TestMapStoreReadPuppyIDDoesNotExist() {
	// given
	assert := tassert.New(s.T())

	// when
	_, err := s.store.ReadPuppy(0x123)

	// then
	assert.Error(err, "Should get an error when attempting to read an non-existing puppy")
}

func (s *storerSuite) TestMapStoreCreatePuppySuccessful() {
	// given
	assert := tassert.New(s.T())
	newPup := puppy2()

	// when
	err := s.store.CreatePuppy(&newPup)

	// then
	assert.NoError(err, "Should not get an error creating a puppy to a free ID")
	createdPup, err := s.store.ReadPuppy(0x1236)
	if assert.NoError(err, "Should be able to read an newly created puppy") {
		assert.Equal(&newPup, createdPup, "Created puppy should be identical to the one passed to Create")
	}
}

func (s *storerSuite) TestMapStoreCreatePuppyIdAlreadyExists() {
	// given
	assert := tassert.New(s.T())
	oldPup := puppy1()
	newPup := puppy2()
	newPup.ID = 0x1234

	// when
	err := s.store.CreatePuppy(&newPup)

	// then
	assert.Error(err, "Should get an error creating a puppy with existing ID")
	currentPet, err := s.store.ReadPuppy(0x1234)
	if assert.NoError(err, "Should be able to read the old existing puppy") {
		assert.Equal(&oldPup, currentPet, "Created puppy should be identical to the one passed to Create")
	}
}

func (s *storerSuite) TestMapStoreUpdatePuppySuccessful() {
	// given
	assert := tassert.New(s.T())
	oldPup := puppy1()
	updatedPup := puppy2()
	updatedPup.ID = 0x1234

	// when
	err := s.store.UpdatePuppy(oldPup.ID, &updatedPup)
	assert.NoError(err, "Should be able to update the puppy")
	// then
	pup, err := s.store.ReadPuppy(oldPup.ID)
	if assert.NoError(err, "Should be able to read the updated puppy") {
		assert.Equal(updatedPup, *pup, "Updated puppy should be identical to the one passed to Update")
	}
}

func (s *storerSuite) TestMapStoreUpdatePuppyIDDoesnotExist() {
	// given
	assert := tassert.New(s.T())
	updatedPup := puppy2()
	updatedPup.ID = 0x123

	// when
	err := s.store.UpdatePuppy(updatedPup.ID, &updatedPup)

	// then
	assert.NoError(err, "Updating existing puppy should not throw an error")
	pup, err := s.store.ReadPuppy(updatedPup.ID)
	if assert.NoError(err, "Updating to non-existing puppy ID is not an error") {
		assert.Equal(updatedPup, *pup, "Updated puppy should be identical to the one passed to Update")
	}
}

func (s *storerSuite) TestMapStoreUpdatePuppyInvalidIDs() {
	// given
	assert := tassert.New(s.T())
	updatedPup := puppy1()

	// when
	err := s.store.UpdatePuppy(0x123, &updatedPup)

	// then
	assert.Error(err, "Updating with un matching id and pet.ID is an error")
}

func (s *storerSuite) TestMapStoreDeletePuppySuccessful() {
	//given
	assert := tassert.New(s.T())

	// when
	err := s.store.DeletePuppy(0x1234)

	// then
	assert.NoError(err, "Delete should successfully delete a puppy")

	_, err = s.store.ReadPuppy(0x1234)
	assert.Error(err, "Should not be able to read a deleted puppy")
}

func (s *storerSuite) TestMapStoreDeletePuppyIDDoesNotExist() {
	//given
	assert := tassert.New(s.T())

	// when
	err := s.store.DeletePuppy(0x123234)

	// then
	assert.Error(err, "Should get an error deleting a non existant ID")
}
