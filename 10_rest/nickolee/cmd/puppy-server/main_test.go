package main

import (
	"os"
	"testing"

	puppy "github.com/anz-bank/go-course/10_rest/nickolee/pkg/puppy"
	"github.com/anz-bank/go-course/10_rest/nickolee/pkg/puppy/store"

	"github.com/stretchr/testify/assert"
)

// testing.M allows you to pass coverage for all but main()
func TestMain(m *testing.M) {
	os.Args = []string{""} // this does the trick. Removes main and gives you the 100% coverage
	os.Exit(m.Run())
}

// This tests the enumerated arguments a user could provide to the store flag
func TestCreateStoreFunc(t *testing.T) {
	userInput := "sync"
	s, err := createStore(userInput)
	assert.IsType(t, store.NewSyncStore(), s)
	assert.NoError(t, err)

	userInput = "anythingElse"
	s, err = createStore(userInput)
	assert.Nil(t, s)
	assert.Error(t, err)
}

func TestReadFileAndUnmarshalPuppiesInvalid(t *testing.T) {
	f := "../../puppy-data/invalid.json"
	puppies, err := readFileAndUnmarshalPuppies(&f)
	assert.Nil(t, puppies)
	assert.Error(t, err)
}

func TestReadFileAndUnmarshalPuppiesValid(t *testing.T) {
	f := "../../puppy-data/puppies.json"
	expected := []puppy.Puppy{
		{ID: 17, Breed: "Vulpix", Colour: "Red", Value: 2900},
		{ID: 27, Breed: "Eevee", Colour: "Light Brown", Value: 1290},
		{ID: 37, Breed: "Vaporeon", Colour: "Sea Blue", Value: 3290}}

	puppies, err := readFileAndUnmarshalPuppies(&f)
	assert.NoError(t, err)
	assert.Equal(t, expected, puppies)
}

func TestParseFlagsLong(t *testing.T) {
	arguments := []string{"--data", "../../puppy-data/invalid.json", "--port", "7777", "--store", "sync"}
	c, err := parseFlags(arguments)

	assert.NoError(t, err)
	assert.Equal(t, "../../puppy-data/invalid.json", *c.file)
	assert.Equal(t, 7777, c.port)
	assert.Equal(t, "sync", c.storer)
}

func TestParseFlagsShort(t *testing.T) {
	arguments := []string{"-d", "../../puppy-data/invalid.json", "-p", "7777", "-s", "sync"}
	c, err := parseFlags(arguments)

	assert.NoError(t, err)
	assert.Equal(t, "../../puppy-data/invalid.json", *c.file)
	assert.Equal(t, 7777, c.port)
	assert.Equal(t, "sync", c.storer)
}

func TestParseFlagsDefault(t *testing.T) {
	fileName = nil // initialise global variable
	var expectedNilStrPtr *string
	arguments := []string{}
	c, err := parseFlags(arguments)
	assert.NoError(t, err)
	assert.Equal(t, expectedNilStrPtr, c.file)
	assert.Equal(t, 7777, c.port)
	assert.Equal(t, "map", c.storer)
}
