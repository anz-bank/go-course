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
	id := suite.store.Create(Puppy{Breed: "Wolf", Color: "Grey", Value: "450"})
	id2 := suite.store.Create(Puppy{Breed: "Boxer", Color: "Brown", Value: "300"})

	a.EqualValues(id+1, id2)
}

func (suite *StoreSuite) TestRead() {
	a := assert.New(suite.T())
	data := suite.store.Read(0)

	a.Equal(data, Puppy{ID: 0, Breed: "Wolf", Color: "Grey", Value: "450"})
}

func (suite *StoreSuite) TestReadNonExistent() {
	a := assert.New(suite.T())
	success := suite.store.Read(100)

	a.Equal(success, Puppy{})
}

func (suite *StoreSuite) TestUpdate() {
	a := assert.New(suite.T())
	success := suite.store.Update(0, Puppy{Breed: "Doberman", Color: "Black", Value: "500"})
	data := suite.store.Read(0)

	a.Equal(success, true)
	a.Equal(data, Puppy{ID: 0, Breed: "Doberman", Color: "Black", Value: "500"})
}

func (suite *StoreSuite) TestUpdateNonExistent() {
	a := assert.New(suite.T())
	success := suite.store.Update(100, Puppy{Breed: "Doberman", Color: "Black", Value: "500"})

	a.Equal(success, false)
}

func (suite *StoreSuite) TestDestroy() {
	a := assert.New(suite.T())
	success := suite.store.Destroy(1)
	data := suite.store.Read(1)

	a.Equal(success, true)
	a.Equal(data, Puppy{})
}

func (suite *StoreSuite) TestDestroyNonExistent() {
	a := assert.New(suite.T())
	success := suite.store.Destroy(100)

	a.Equal(success, false)
}

func (suite *StoreSuite) TestIdIncrementOnDelete() {
	a := assert.New(suite.T())
	id := suite.store.Create(Puppy{Breed: "Greyhound", Color: "Light Brown", Value: "700"})
	suite.store.Destroy(id)
	newID := suite.store.Create(Puppy{Breed: "Greyhound", Color: "Light Brown", Value: "700"})

	a.EqualValues(newID, 3)
}

func TestStore(t *testing.T) {
	suite.Run(t, &StoreSuite{store: NewMapStore()})
}
