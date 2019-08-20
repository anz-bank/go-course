package store

import (
	"testing"

	p "github.com/anz-bank/go-course/08_project/patrickmarabeas/pkg/puppy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StoreSuite struct {
	suite.Suite
	store Storer
}

func (suite *StoreSuite) TestCreate() {
	a := assert.New(suite.T())

	id, error := suite.store.Create(p.Puppy{Breed: "Wolf", Color: "Grey", Value: 450})
	a.Equal(id, 0)
	a.Equal(error, nil)
}

func (suite *StoreSuite) TestCreateSecond() {
	a := assert.New(suite.T())

	id, error := suite.store.Create(p.Puppy{Breed: "Boxer", Color: "Brown", Value: 300})
	a.Equal(id, 1)
	a.Equal(error, nil)
}

func (suite *StoreSuite) TestCreateNegativeNumber() {
	a := assert.New(suite.T())

	id, error := suite.store.Create(p.Puppy{Breed: "Wolf", Color: "Grey", Value: -100})
	a.Equal(id, -1)
	a.Equal(error, p.NewError(p.NegativeValue))
}

func (suite *StoreSuite) TestRead() {
	a := assert.New(suite.T())

	data, error := suite.store.Read(0)
	a.Equal(data, p.Puppy{ID: 0, Breed: "Wolf", Color: "Grey", Value: 450})
	a.Equal(error, nil)
}

func (suite *StoreSuite) TestReadNonExistent() {
	a := assert.New(suite.T())

	success, error := suite.store.Read(100)
	a.Equal(success, p.Puppy{})
	a.Equal(error, p.NewError(p.IDNotFound))
}

func (suite *StoreSuite) TestUpdate() {
	a := assert.New(suite.T())

	success, error := suite.store.Update(0, p.Puppy{Breed: "Doberman", Color: "Black", Value: 500})
	a.Equal(success, true)
	a.Equal(error, nil)

	data, error := suite.store.Read(0)
	a.Equal(data, p.Puppy{ID: 0, Breed: "Doberman", Color: "Black", Value: 500})
	a.Equal(error, nil)
}

func (suite *StoreSuite) TestUpdateNonExistent() {
	a := assert.New(suite.T())

	success, error := suite.store.Update(100, p.Puppy{Breed: "Doberman", Color: "Black", Value: 500})
	a.Equal(success, false)
	a.Equal(error, p.NewError(p.IDNotFound))
}

func (suite *StoreSuite) TestUpdateNegativeNumber() {
	a := assert.New(suite.T())

	success, error := suite.store.Update(0, p.Puppy{Breed: "Doberman", Color: "Black", Value: -500})
	a.Equal(success, false)
	a.Equal(error, p.NewError(p.NegativeValue))
}

func (suite *StoreSuite) TestDestroy() {
	a := assert.New(suite.T())

	success, error := suite.store.Destroy(1)
	a.Equal(success, true)
	a.Equal(error, nil)

	data, error := suite.store.Read(1)
	a.Equal(data, p.Puppy{})
	a.Equal(error, p.NewError(p.IDNotFound))
}

func (suite *StoreSuite) TestDestroyNonExistent() {
	a := assert.New(suite.T())

	success, error := suite.store.Destroy(100)
	a.Equal(success, false)
	a.Equal(error, p.NewError(p.IDNotFound))
}

func (suite *StoreSuite) TestIdIncrementOnDelete() {
	a := assert.New(suite.T())
	id, _ := suite.store.Create(p.Puppy{Breed: "Greyhound", Color: "Light Brown", Value: 700})
	success, _ := suite.store.Destroy(id)
	a.Equal(success, true)

	newID, _ := suite.store.Create(p.Puppy{Breed: "Greyhound", Color: "Light Brown", Value: 700})
	a.Equal(newID, 3)
}

func TestStore(t *testing.T) {
	suite.Run(t, &StoreSuite{store: NewMapStore()})
	suite.Run(t, &StoreSuite{store: NewSyncStore()})
}
