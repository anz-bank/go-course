package store

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	puppy "github.com/anz-bank/go-course/08_project/n0npax/pkg/puppy"
	"github.com/stretchr/testify/assert"
	tassert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/syndtr/goleveldb/leveldb"
)

type storerImpl int

const (
	syncStorer storerImpl = iota
	memStorer
	leveldbStorer
)

var (
	puppy0 = func() puppy.Puppy {
		return puppy.Puppy{
			Breed:  "Type A",
			Colour: "Grey",
			Value:  300,
		}
	}
	puppy1 = func() puppy.Puppy {
		return puppy.Puppy{
			Breed:  "Type B",
			Colour: "Brown",
			Value:  400,
		}
	}
	puppyNegativeValue = func() puppy.Puppy {
		return puppy.Puppy{
			Breed:  "Type C",
			Colour: "Red",
			Value:  -1,
		}
	}
	puppyNegativeID = func() puppy.Puppy {
		return puppy.Puppy{
			Breed:  "Type D",
			Colour: "Blue",
			Value:  100,
			ID:     -1,
		}
	}
)

type storerSuite struct {
	suite.Suite
	store puppy.Storer
	impl  storerImpl
}

func (s *storerSuite) SetupSuite() {
	// Remove old db if exists
	os.RemoveAll(levelDBPath)
}

func (s *storerSuite) TearDownTest() {
	if ldbs, ok := s.store.(*LevelDBStore); ok {
		ldbs.ldb.Close()
	}
}

func (s *storerSuite) SetupTest() {
	switch s.impl {
	case syncStorer:
		s.store = NewSyncStore()
	case memStorer:
		s.store = NewMemStore()
	case leveldbStorer:
		s.store = NewLevelDBStorer()
	default:
		panic("Unrecognised storer implementation")
	}
	p := puppy0()
	_, err := s.store.CreatePuppy(&p)
	if err != nil {
		panic(err)
	}
}

func (s *storerSuite) TestCreatePuppySuccessful() {
	assert := tassert.New(s.T())
	newPuppy0, newPuppy1 := puppy0(), puppy1()
	id0, err := s.store.CreatePuppy(&newPuppy0)
	assert.NoError(err, "Creating p should be ok")
	id1, err := s.store.CreatePuppy(&newPuppy1)
	assert.NoError(err, "Creating p should be ok")
	assert.Equal(id0, id1-1, "2nd id should be 1st +1, got %v and %v", id0, id1)
}

func (s *storerSuite) TestCreatePuppyNegativeValue() {
	assert := tassert.New(s.T())
	newPuppy := puppyNegativeValue()
	_, err := s.store.CreatePuppy(&newPuppy)
	assert.Error(err, "negative ID should cause an error")
}

func (s *storerSuite) TestReadPuppySuccessful() {
	assert := tassert.New(s.T())
	newPuppy := puppy0()
	id, err := s.store.CreatePuppy(&newPuppy)
	assert.NoError(err, "Creating p should be ok")
	readPuppy, err := s.store.ReadPuppy(id)
	if assert.NoError(err, "Should be able to read puppy0 from store") {
		assert.Equal(&newPuppy, readPuppy, "store should return identic puppy")
	}
}

func (s *storerSuite) TestReadPuppyIDDoesNotExist() {
	assert := tassert.New(s.T())
	_, err := s.store.ReadPuppy(1000)
	assert.Error(err, "Should get an error when attempting to read an non-existing puppy")
}

func (s *storerSuite) TestReadPuppyNegativeID() {
	assert := tassert.New(s.T())
	_, err := s.store.ReadPuppy(-1)
	assert.Error(err, "negative ID should cause an error")
}

func (s *storerSuite) TestUpdatePuppy() {
	assert := tassert.New(s.T())
	existingPuppy, err := s.store.ReadPuppy(0)
	assert.NoError(err, "Reading p should not return error")
	existingPuppy.Colour = "Purple"
	err = s.store.UpdatePuppy(0, existingPuppy)
	assert.NoError(err, "Update should not return any error")
	p, err := s.store.ReadPuppy(0)
	assert.NoError(err, "Reading p should not return error")
	assert.Equal(existingPuppy.Colour, p.Colour, "Updated colour missmatch")
}

func (s *storerSuite) TestUpdatePuppyCorruptedID() {
	assert := tassert.New(s.T())
	existingPuppy, err := s.store.ReadPuppy(0)
	assert.NoError(err, "Reading p should be ok")
	err = s.store.UpdatePuppy(1000, existingPuppy)
	assert.Error(err, "Should get an error when attempting to update with corrupted id")
}

func (s *storerSuite) TestUpdatePuppyIDDoesNotExist() {
	assert := tassert.New(s.T())
	newPuppy := puppy0()
	err := s.store.UpdatePuppy(1000, &newPuppy)
	assert.Error(err, "Should get an error when attempting to update an non-existing puppy")
}

func (s *storerSuite) TestUpdatePuppyNegativeID() {
	assert := tassert.New(s.T())
	newPuppy := puppyNegativeID()
	err := s.store.UpdatePuppy(newPuppy.ID, &newPuppy)
	assert.Error(err, "negative ID should cause an error")
}

func (s *storerSuite) TestUpdatePuppyNegativeValue() {
	assert := tassert.New(s.T())
	newPuppy := puppyNegativeValue()
	err := s.store.UpdatePuppy(newPuppy.ID, &newPuppy)
	assert.Error(err, "negative ID should cause an error")
}

func (s *storerSuite) TestDeletePuppySuccessful() {
	assert := tassert.New(s.T())
	existingPuppy := puppy0()
	deleted, err := s.store.DeletePuppy(0)
	assert.NoError(err, "Delete should successfully delete a puppy")
	assert.True(deleted, "Delete should return true indicating a p was deleted")
	_, err = s.store.ReadPuppy(existingPuppy.ID)
	assert.Error(err, "Should not be able to read a deleted ID")
}

func (s *storerSuite) TestDeletePuppyIDDoesNotExist() {
	assert := tassert.New(s.T())
	_, err := s.store.DeletePuppy(1000)
	assert.Error(err, "Should not be able to delete p with non existing ID")
}

func (s *storerSuite) TestDeletePuppyNegativeID() {
	assert := tassert.New(s.T())
	_, err := s.store.DeletePuppy(-1)
	assert.Error(err, "negative ID should cause an error")
}

func TestStorer(t *testing.T) {
	suite.Run(t, &storerSuite{impl: syncStorer})
	suite.Run(t, &storerSuite{impl: memStorer})
	suite.Run(t, &storerSuite{impl: leveldbStorer})
}

func TestBrokenDataInLevelDB(t *testing.T) {
	// Prepare corrupted data to cause internal error
	func() {
		assert := tassert.New(t)
		db, _ := leveldb.OpenFile(levelDBPath, nil)
		defer db.Close()
		puppyByte, err := json.Marshal("this string cannot be casted to puppy")
		assert.NoError(err, "no error expected during marshaling string")
		byteID := []byte(strconv.Itoa(999))
		err = db.Put(byteID, puppyByte, nil)
		assert.NoError(err, "no error expected during preparing corrupted data in db")
	}()
	s := NewLevelDBStorer()
	defer s.ldb.Close()
	_, err := s.ReadPuppy(999)
	assert.Error(t, err)
}
