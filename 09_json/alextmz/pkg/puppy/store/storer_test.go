package store

import (
	"testing"

	. "github.com/anz-bank/go-course/08_project/alextmz/pkg/puppy"
	"github.com/stretchr/testify/suite"
)

type storerSuite struct {
	suite.Suite
	p  []Puppy
	st Storer
}

func (su *storerSuite) SetupTest() {
	su.p = nil
	var p = []Puppy{
		{ID: 256, Breed: "Dogo", Colour: "White", Value: 500},
		{ID: 512, Breed: "Mastiff", Colour: "Brindle", Value: 700},
		{ID: 768, Breed: "Fila", Colour: "Golden", Value: 900},
		{ID: 1024, Breed: "Wolfhound", Colour: "Gray", Value: -50},
	}
	su.p = append(su.p, p...)
	err := su.st.CreatePuppy(&su.p[0])
	su.Nil(err)
}

// TearDownTest cleans everything after each test runs.
// It impacts test performance as things get re-populated soon afterwards,
// but we're looking for reliable here not fast
func (su *storerSuite) TearDownTest() {
	for _, v := range su.p {
		if _, err := su.st.ReadPuppy(v.ID); err == nil {
			err := su.st.DeletePuppy(v.ID)
			if err != nil {
				su.T().Fatalf("internal error when deleting, this should never happen: %#v\n", err)
			}
		}
	}
	su.p = nil
}

// TestCreatePuppy tests:
// - If we can create without error
// - If we error when creating something that already exists
// - If we can create and what we read back matches what we expect
func (su *storerSuite) TestCreatePuppy() {
	err := su.st.CreatePuppy(&su.p[1])
	su.Nil(err)
	err = su.st.CreatePuppy(&su.p[1])
	su.NotNil(err)
	err = su.st.CreatePuppy(&su.p[3])
	su.NotNil(err)
	pa, _ := su.st.ReadPuppy(su.p[1].ID)
	su.Equalf(su.p[1], *pa, "expected = %#v, actual = %#v\n", su.p[1], pa)
}

// TestReadPuppy tests:
// - If we can read without error
// - If we error when reading what doesn't exist
// - If contents of something read match what we expect
func (su *storerSuite) TestReadPuppy() {
	_, err := su.st.ReadPuppy(su.p[0].ID)
	su.Nil(err)
	_, err = su.st.ReadPuppy(su.p[1].ID)
	su.NotNil(err)
}

// TestUpdatePuppy tests:
// - If we can update without error
// - If we error when trying to update what doesn't exist
// - If the updated content matches what we expect
func (su *storerSuite) TestUpdatePuppy() {
	err := su.st.UpdatePuppy(su.p[0].ID, &su.p[1])
	su.Nil(err)
	err = su.st.UpdatePuppy(su.p[2].ID, &su.p[2])
	su.NotNil(err)
	err = su.st.UpdatePuppy(su.p[0].ID, &su.p[3])
	su.NotNil(err)
	pa, _ := su.st.ReadPuppy(su.p[0].ID)
	su.Equalf(su.p[1], *pa, "expected = %#v, actual = %#v\n", su.p[1], pa)
}

// TestDeletePuppy tests:
// - If we can delete without error
// - If we error when trying to delete what doesn't exist
// - If after we delete, we cannot read back the data
func (su *storerSuite) TestDeletePuppy() {
	err := su.st.DeletePuppy(su.p[0].ID)
	su.Nil(err)
	err = su.st.DeletePuppy(su.p[1].ID)
	su.NotNil(err)
	pa, err := su.st.ReadPuppy(su.p[0].ID)
	su.Nilf(pa, "expected = %#v, actual = %#v\n", nil, pa)
	su.NotNil(err)
}

func Test_Suite(t *testing.T) {
	suite.Run(t, &storerSuite{st: NewSyncStore()})
	suite.Run(t, &storerSuite{st: NewmapStore()})
}
