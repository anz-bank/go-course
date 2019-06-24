package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StorerTestSuite struct {
	suite.Suite
	storer Storer
}

func (suite StorerTestSuite) SetupTest() {
	suite.storer.CreatePuppy(&Puppy{1, "catdog", "brown + spots", "9.99"})
	suite.storer.CreatePuppy(&Puppy{2, "non-cube", "yellowy-blue", "3.50"})
}

func (suite StorerTestSuite) TestCreate() {
	suite.storer.CreatePuppy(&Puppy{3, "boxer", "yellowish-yellow", "99999.01"})
	pup := suite.storer.ReadPuppy(3)
	suite.Equal(3, pup.pid)
	suite.Equal("boxer", pup.breed)
	suite.Equal("yellowish-yellow", pup.colour)
	suite.Equal("99999.01", pup.value)
}

func (suite StorerTestSuite) TestRead() {
	pup := suite.storer.ReadPuppy(1)
	suite.Equal(1, pup.pid)
	suite.Equal("catdog", pup.breed)
	suite.Equal("brown + spots", pup.colour)
	suite.Equal("9.99", pup.value)
}

func (suite StorerTestSuite) TestUpdate() {
	suite.storer.UpdatePuppy(&Puppy{2, "cube", "bluey-yellow", "1.75"})
	pup := suite.storer.ReadPuppy(2)
	suite.Equal(2, pup.pid)
	suite.Equal("cube", pup.breed)
	suite.Equal("bluey-yellow", pup.colour)
	suite.Equal("1.75", pup.value)
}

func (suite StorerTestSuite) TestDelete() {
	suite.storer.DeletePuppy(1)
	pup1 := suite.storer.ReadPuppy(1)
	pup2 := suite.storer.ReadPuppy(2)
	suite.Nil(pup1)
	suite.Equal(2, pup2.pid)
	suite.Equal("non-cube", pup2.breed)
	suite.Equal("yellowy-blue", pup2.colour)
	suite.Equal("3.50", pup2.value)
}

func TestMapStorer(t *testing.T) {
	store := StorerTestSuite{
		storer: &MapStore{store: make(map[int]Puppy)},
	}
	suite.Run(t, &store)
}

func TestSyncStorer(t *testing.T) {
	store := StorerTestSuite{storer: &SyncStore{}}
	suite.Run(t, &store)
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()
	expected := `{100 lab orange $9.99}
&{map[99:{99 poodle red $10.99} 101:{101 cat striped $99.99}]}
{100 lab orange $9.99}
{99 poodle red $10.99}
{101 cat striped $99.99}
`
	actual := buf.String()

	assert.Equalf(t, expected, actual,
		"Unexpected output in main()\nexpected: %q\nactual: %q",
		expected, actual)
}

func TestNormalMap(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	pups := []Puppy{
		{99, "poodle", "blue", "$10.99"},
		{100, "lab", "orange", "$9.99"},
		{101, "cat", "striped", "$99.99"},
	}
	usingNormMap(pups)

	expected := `{100 lab orange $9.99}
&{map[99:{99 poodle red $10.99} 101:{101 cat striped $99.99}]}
`
	actual := buf.String()

	assert.Equalf(t, expected, actual,
		"Unexpected output in main()\nexpected: %q\nactual: %q",
		expected, actual)
}

func TestSyncMap(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	pups := []Puppy{
		{99, "poodle", "blue", "$10.99"},
		{100, "lab", "orange", "$9.99"},
		{101, "cat", "striped", "$99.99"},
	}
	usingSyncMap(pups)
	expected := `{100 lab orange $9.99}
{99 poodle red $10.99}
{101 cat striped $99.99}
`
	actual := buf.String()

	assert.Equalf(t, expected, actual,
		"Unexpected output in main()\nexpected: %q\nactual: %q",
		expected, actual)
}
