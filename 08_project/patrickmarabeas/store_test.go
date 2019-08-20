package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StoreSuite struct {
	suite.Suite
	store Storer
}

func (suite *StoreSuite) TestCreate() {
	a := assert.New(suite.T())

	id, error := suite.store.Create(Puppy{Breed: "Wolf", Color: "Grey", Value: 450})
	a.Equal(id, 0)
	a.Equal(error, nil)
}

func (suite *StoreSuite) TestCreateSecond() {
	a := assert.New(suite.T())

	id, error := suite.store.Create(Puppy{Breed: "Boxer", Color: "Brown", Value: 300})
	a.Equal(id, 1)
	a.Equal(error, nil)
}

func (suite *StoreSuite) TestCreateNegativeNumber() {
	a := assert.New(suite.T())

	id, error := suite.store.Create(Puppy{Breed: "Wolf", Color: "Grey", Value: -100})
	a.Equal(id, -1)
	a.Equal(error, NewError(NegativeValue))
}

func (suite *StoreSuite) TestRead() {
	a := assert.New(suite.T())

	data, error := suite.store.Read(0)
	a.Equal(data, Puppy{ID: 0, Breed: "Wolf", Color: "Grey", Value: 450})
	a.Equal(error, nil)
}

func (suite *StoreSuite) TestReadNonExistent() {
	a := assert.New(suite.T())

	success, error := suite.store.Read(100)
	a.Equal(success, Puppy{})
	a.Equal(error, NewError(IDNotFound))
}

func (suite *StoreSuite) TestUpdate() {
	a := assert.New(suite.T())

	success, error := suite.store.Update(0, Puppy{Breed: "Doberman", Color: "Black", Value: 500})
	a.Equal(success, true)
	a.Equal(error, nil)

	data, error := suite.store.Read(0)
	a.Equal(data, Puppy{ID: 0, Breed: "Doberman", Color: "Black", Value: 500})
	a.Equal(error, nil)
}

func (suite *StoreSuite) TestUpdateNonExistent() {
	a := assert.New(suite.T())

	success, error := suite.store.Update(100, Puppy{Breed: "Doberman", Color: "Black", Value: 500})
	a.Equal(success, false)
	a.Equal(error, NewError(IDNotFound))
}

func (suite *StoreSuite) TestUpdateNegativeNumber() {
	a := assert.New(suite.T())

	success, error := suite.store.Update(0, Puppy{Breed: "Doberman", Color: "Black", Value: -500})
	a.Equal(success, false)
	a.Equal(error, NewError(NegativeValue))
}

func (suite *StoreSuite) TestDestroy() {
	a := assert.New(suite.T())

	success, error := suite.store.Destroy(1)
	a.Equal(success, true)
	a.Equal(error, nil)

	data, error := suite.store.Read(1)
	a.Equal(data, Puppy{})
	a.Equal(error, NewError(IDNotFound))
}

func (suite *StoreSuite) TestDestroyNonExistent() {
	a := assert.New(suite.T())

	success, error := suite.store.Destroy(100)
	a.Equal(success, false)
	a.Equal(error, NewError(IDNotFound))
}

func (suite *StoreSuite) TestIdIncrementOnDelete() {
	a := assert.New(suite.T())
	id, _ := suite.store.Create(Puppy{Breed: "Greyhound", Color: "Light Brown", Value: 700})
	success, _ := suite.store.Destroy(id)
	a.Equal(success, true)

	newID, _ := suite.store.Create(Puppy{Breed: "Greyhound", Color: "Light Brown", Value: 700})
	a.Equal(newID, 3)
}

func TestStore(t *testing.T) {
	suite.Run(t, &StoreSuite{store: NewMapStore()})
	suite.Run(t, &StoreSuite{store: NewSyncStore()})
}
