package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/anz-bank/go-course/10_rest/willshen8/pkg/puppy"

	"github.com/stretchr/testify/assert"
)

//Unit test for parseCmdArgs() with long flags
func TestParseCmdArgs(t *testing.T) {
	args = []string{"--data", "./../../data/puppy.json", "--port", "1234", "--store", "map"}
	returnedArgs := parseCmdArgs(args)
	assert.Equal(t, "./../../data/puppy.json", returnedArgs.file)
	assert.Equal(t, 1234, returnedArgs.port)
	assert.Equal(t, "map", returnedArgs.storer)
}

//Unit test for parseCmdArgs() with short flags
func TestShortFlag(t *testing.T) {
	args = []string{"-d", "./../../data/puppy.json", "-p", "1234", "-s", "map"}
	returnedArgs := parseCmdArgs(args)
	assert.Equal(t, "./../../data/puppy.json", returnedArgs.file)
	assert.Equal(t, 1234, returnedArgs.port)
	assert.Equal(t, "map", returnedArgs.storer)
}

//Unit test (positive test) for parsePortFlag()
func TestParseCorrectPort(t *testing.T) {
	// setup the test server with arguments
	args = []string{"--data", "./../../data/puppy.json", "--port", "1234"}
	parsedArgs := parseCmdArgs(args)
	flagReturned, _ := parsePortFlag(parsedArgs)
	var expectedFlag = "1234"
	assert.Equal(t, expectedFlag, flagReturned)
}

//Unit test (negative test) for parsePortFlag()
func TestParseWrongPort(t *testing.T) {
	// setup the test server with arguments
	args = []string{"--data", "./../../data/puppy.json", "--port", "123456789"}
	parsedArgs := parseCmdArgs(args)
	flagReturned, flagError := parsePortFlag(parsedArgs)
	var expectedFlag = "8888"
	var expectedError = string(puppy.InvalidPort)
	assert.Equal(t, expectedFlag, flagReturned)
	assert.Equal(t, expectedError, flagError.Error())

}

// Unit test for createPuppies()
func TestSavePuppiesToStore(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	var puppies []puppy.Puppy
	puppies = append(puppies, puppy.Puppy{ID: 1, Breed: "Jack Russell Terrier", Color: "White and Brown", Value: "1500"})
	storer := puppy.NewMapStore()
	returnedPuppyStorer, _ := createPuppies(storer, puppies)

	expected := "&{map[1:{1 Jack Russell Terrier White and Brown 1500}] 1}\n"
	fmt.Fprintln(out, returnedPuppyStorer)
	actual := buf.String()
	assert.Equal(t, expected, actual)
}

func TestMain(m *testing.M) {
	// setup the test server with arguments
	args = []string{"--data", "./../../data/puppy.json", "--port", "1234", "--store", "map"}
	os.Exit(m.Run())
}

func TestParseCmdFlagsError(t *testing.T) {
	args = []string{"--wrongFlag"}
	assert.NotPanics(t, main)
	os.Exit(0)
}

func TestBadFileName(t *testing.T) {
	args = []string{"--data=helloWorld.json"}
	assert.Panics(t, main)
}

func TestInvalidJSON(t *testing.T) {
	args = []string{"--data=invalid_puppies.json"}
	assert.Panics(t, main)
}

func TestUnmarshalJSON(t *testing.T) {
	args = []string{"--data=type_mismatch.json"}
	assert.Panics(t, main)
}
