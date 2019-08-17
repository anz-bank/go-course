package puppy

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	out io.Writer = os.Stdout
)

type StorerSuite struct {
	suite.Suite
	storer Storer
}

func (s *StorerSuite) TestCreatePuppySuccessfully() {
	m := assert.New(s.T())
	puppy1 := Puppy{ID: 1, Breed: "Jack Russell Terrier", Color: "White and Brown", Value: "1500"}
	createdID, _ := s.storer.CreatePuppy(&puppy1)
	readPuppy1, _ := s.storer.ReadPuppy(createdID)
	m.Equal(&puppy1, readPuppy1)
}

func (s *StorerSuite) TestCreatePuppyWithValueLessThanZero() {
	InvalidPuppy := Puppy{ID: 22, Breed: "Pomeranian", Color: "White", Value: "-100"}
	m := assert.New(s.T())
	var buf bytes.Buffer
	out = &buf
	createdID, err := s.storer.CreatePuppy(&InvalidPuppy)
	fmt.Fprintln(out, err)
	m.Error(err, "Negative ID should give an error")
	m.Equal(uint32(0), createdID)
}

func (s *StorerSuite) TestCreatePuppyWithWrongValue() {
	m := assert.New(s.T())
	WrongValuePuppy := Puppy{ID: 88, Breed: "Spaniel", Color: "Black", Value: "$$$"}
	_, err := s.storer.CreatePuppy(&WrongValuePuppy)
	m.Error(err, "Puppy value should be numerical")
}

func (s *StorerSuite) TestCreateNextPuppyWithCorrectID() {
	m := assert.New(s.T())
	puppy2 := Puppy{ID: 1234, Breed: "Fox Terrier", Color: "Black", Value: "1300"}
	puppy3 := Puppy{ID: 100, Breed: "German Shepperd", Color: "Brown", Value: "2000"}
	createID1, creatError1 := s.storer.CreatePuppy(&puppy2)
	if createID1 != uint32(1) && creatError1 != nil {
		panic("Error creating puppy")
	}
	// create a second puppy should give an id of 2
	createdID2, _ := s.storer.CreatePuppy(&puppy3)
	m.Equal(uint32(2), createdID2)
}

func (s *StorerSuite) TestReadPuppySuccessfully() {
	m := assert.New(s.T())
	puppy4 := Puppy{ID: 120, Breed: "Golden Retriever", Color: "Golden", Value: "2500"}
	returnedPuppyID4, _ := s.storer.CreatePuppy(&puppy4)
	returnPuppy4, _ := s.storer.ReadPuppy(returnedPuppyID4)
	m.Equal(&puppy4, returnPuppy4)
}

func (s *StorerSuite) TestReadNonExistentPuppy() {
	m := assert.New(s.T())
	returnPuppy, err := s.storer.ReadPuppy(100)
	m.Nil(returnPuppy)
	m.Error(err, "Non existent puppy ID should give an error")
}

func (s *StorerSuite) TestUpdatePuppySuccessfully() {
	m := assert.New(s.T())
	puppy5 := Puppy{ID: 200, Breed: "Chihuahua", Color: "White", Value: "500"}
	puppy6 := Puppy{ID: 300, Breed: "Husky", Color: "White", Value: "3500"}
	returnedPuppyID, _ := s.storer.CreatePuppy(&puppy5)
	err := s.storer.UpdatePuppy(returnedPuppyID, &puppy6)
	m.NoError(err)
}

func (s *StorerSuite) TestUpdatePuppyWithNegativeValue() {
	m := assert.New(s.T())
	puppy1 := Puppy{ID: 200, Breed: "Chihuahua", Color: "White", Value: "500"}
	puppy2 := Puppy{ID: 300, Breed: "Husky", Color: "White", Value: "-3500"}
	returnedPuppyID, _ := s.storer.CreatePuppy(&puppy1)
	err := s.storer.UpdatePuppy(returnedPuppyID, &puppy2)
	m.Error(err, "Update puppy with an negative value should give an error")
}

func (s *StorerSuite) TestUpdatePuppyWithNonIntValue() {
	m := assert.New(s.T())
	puppy1 := Puppy{ID: 200, Breed: "Chihuahua", Color: "White", Value: "500"}
	puppy2 := Puppy{ID: 300, Breed: "Husky", Color: "White", Value: "blah"}
	returnedPuppyID, _ := s.storer.CreatePuppy(&puppy1)
	err := s.storer.UpdatePuppy(returnedPuppyID, &puppy2)
	m.Error(err, "Update puppy with an non int should give an error")
}

func (s *StorerSuite) TestUpdateNonExistentID() {
	m := assert.New(s.T())
	puppy1 := Puppy{ID: 1, Breed: "Jack Russell Terrier", Color: "White and Brown", Value: "1500"}
	err := s.storer.UpdatePuppy(888, &puppy1)
	m.Error(err, "Non existent puppy ID should give an error")
}

func (s *StorerSuite) TestDeletePuppy() {
	m := assert.New(s.T())
	puppy7 := Puppy{ID: 700, Breed: "Pomeranian", Color: "White", Value: "700"}
	createdPuppyID, _ := s.storer.CreatePuppy(&puppy7)
	deleteErr := s.storer.DeletePuppy(createdPuppyID)
	returnPuppy, _ := s.storer.ReadPuppy(createdPuppyID)
	m.NoError(deleteErr)
	s.Nil(returnPuppy)
}

func (s *StorerSuite) TestDeleteNonExistentPuppy() {
	m := assert.New(s.T())
	err := s.storer.DeletePuppy(uint32(10000))
	m.Error(err, "Non existent puppy ID should give an error")
}

func TestStorer(t *testing.T) {
	suite.Run(t, &StorerSuite{storer: NewMapStore()})
	suite.Run(t, &StorerSuite{storer: NewSyncStore()})
}
