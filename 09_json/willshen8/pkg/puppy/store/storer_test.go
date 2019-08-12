package puppy

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	puppy "github.com/anz-bank/go-course/09_json/willshen8/pkg/puppy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	out       io.Writer = os.Stdout
	ReadError           = puppy.Error{Code: puppy.NonExistentPuppy,
		Message: "Puppy ID can not be found, read operation failed."}
	UpdateError = puppy.Error{Code: puppy.NonExistentPuppy,
		Message: "Puppy ID can not be found, update operation failed."}
	DeleteError = puppy.Error{Code: puppy.NonExistentPuppy,
		Message: "Puppy ID can not be found, delete operation failed."}
	WrongValueError = puppy.Error{Code: puppy.ErrorValueFormat,
		Message: "Unrecongised puppy value."}
	ErrorCreateID = uint32(0)
)

type StorerSuite struct {
	suite.Suite
	storer Storer
}

func (s *StorerSuite) TestCreatePuppySuccessfully() {
	m := assert.New(s.T())
	Puppy1 := puppy.Puppy{ID: 1, Breed: "Jack Russell Terrier", Color: "White and Brown", Value: "1500"}
	createdID, _ := s.storer.CreatePuppy(&Puppy1)
	readPuppy1, _ := s.storer.ReadPuppy(createdID)
	m.Equal(&Puppy1, readPuppy1)
}

func (s *StorerSuite) TestCreatePuppyWithValueLessThanZero() {
	InvalidPuppy := puppy.Puppy{ID: 22, Breed: "Pomeranian", Color: "White", Value: "-100"}
	createError := "Error code 1001: Puppy value can't be less than 0.\n"
	m := assert.New(s.T())
	var buf bytes.Buffer
	out = &buf
	createdID, error := s.storer.CreatePuppy(&InvalidPuppy)
	fmt.Fprintln(out, error)
	actualError := buf.String()
	m.Equal(createError, actualError)
	m.Equal(ErrorCreateID, createdID)
}

func (s *StorerSuite) TestCreatePuppyWithNoValue() {
	m := assert.New(s.T())
	NoValuePuppy := puppy.Puppy{ID: 88, Breed: "Spaniel", Color: "Black", Value: "$$$"}
	_, err := s.storer.CreatePuppy(&NoValuePuppy)
	m.Equal(&WrongValueError, err)
}

func (s *StorerSuite) TestCreateNextPuppyWithCorrectID() {
	m := assert.New(s.T())
	Puppy2 := puppy.Puppy{ID: 1234, Breed: "Fox Terrier", Color: "Black", Value: "1300"}
	Puppy3 := puppy.Puppy{ID: 100, Breed: "German Shepperd", Color: "Brown", Value: "2000"}
	createID1, creatError1 := s.storer.CreatePuppy(&Puppy2)
	if createID1 != uint32(1) && creatError1 != nil {
		panic("Error creating puppy")
	}
	// create a second puppy should give an id of 2
	createdID2, _ := s.storer.CreatePuppy(&Puppy3)
	m.Equal(uint32(2), createdID2)
}

func (s *StorerSuite) TestReadPuppySuccessfully() {
	m := assert.New(s.T())
	Puppy4 := puppy.Puppy{ID: 120, Breed: "Golden Retriever", Color: "Golden", Value: "2500"}
	returnedPuppyID4, _ := s.storer.CreatePuppy(&Puppy4)
	returnPuppy4, _ := s.storer.ReadPuppy(returnedPuppyID4)
	m.Equal(&Puppy4, returnPuppy4)
}

func (s *StorerSuite) TestReadNonExistentPuppy() {
	m := assert.New(s.T())
	returnPuppy, error := s.storer.ReadPuppy(100)
	m.Nil(returnPuppy)
	m.Equal(&ReadError, error)
}

func (s *StorerSuite) TestUpdatePuppySuccessfully() {
	Puppy5 := puppy.Puppy{ID: 200, Breed: "Chihuahua", Color: "White", Value: "500"}
	Puppy6 := puppy.Puppy{ID: 300, Breed: "Husky", Color: "White", Value: "3500"}
	returnedPuppyID, _ := s.storer.CreatePuppy(&Puppy5)
	updateResult, _ := s.storer.UpdatePuppy(returnedPuppyID, &Puppy6)
	s.Equal(true, updateResult)
}

func (s *StorerSuite) TestUpdateNonExistentID() {
	Puppy1 := puppy.Puppy{ID: 1, Breed: "Jack Russell Terrier", Color: "White and Brown", Value: "1500"}
	updateResult, error := s.storer.UpdatePuppy(888, &Puppy1)
	s.Equal(&UpdateError, error)
	s.Equal(false, updateResult)
}

func (s *StorerSuite) TestDeletePuppy() {
	Puppy7 := puppy.Puppy{ID: 700, Breed: "Pomeranian", Color: "White", Value: "700"}
	createdPuppyID, _ := s.storer.CreatePuppy(&Puppy7)
	deleteResult, _ := s.storer.DeletePuppy(createdPuppyID)
	s.Equal(true, deleteResult)
	returnPuppy, _ := s.storer.ReadPuppy(createdPuppyID)
	s.Nil(returnPuppy)
}

func (s *StorerSuite) TestDeleteNonExistentPuppy() {
	deleteResult, error := s.storer.DeletePuppy(uint32(10000))
	s.Equal(false, deleteResult)
	s.Equal(&DeleteError, error)
}

func TestStorer(t *testing.T) {
	suite.Run(t, &StorerSuite{storer: NewMapStore()})
	suite.Run(t, &StorerSuite{storer: NewSyncStore()})
}
