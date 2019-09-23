package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/anz-bank/go-course/11_notify/willshen8/pkg/puppy"

	"github.com/stretchr/testify/assert"
)

// Unit test for parseCmdArgs() with long flags
func TestParseCmdArgs(t *testing.T) {
	args = []string{"--data", "/dev/null", "--port", "1234", "--store", "map"}
	returnedArgs, _ := parseCmdArgs(args)
	assert.NotNil(t, returnedArgs.file)
	assert.Equal(t, 1234, returnedArgs.port)
	assert.Equal(t, "map", returnedArgs.storer)
}

func TestParseCmdArgsWithError(t *testing.T) {
	args = []string{"--wrongflag", "/dev/null", "-p", "1234", "-s", "map"}
	_, err := parseCmdArgs(args)
	assert.Error(t, err)
}

// Unit test for parseCmdArgs() with short flags
func TestShortFlag(t *testing.T) {
	args = []string{"-d", "/dev/null", "-p", "1234", "-s", "map"}
	returnedArgs, _ := parseCmdArgs(args)
	assert.NotNil(t, returnedArgs.file)
	assert.Equal(t, 1234, returnedArgs.port)
	assert.Equal(t, "map", returnedArgs.storer)
}

// Unit test (positive test) for validatePortFlag()
func TestParseCorrectPort(t *testing.T) {
	// setup the test server with arguments
	args = []string{"--data", "/dev/null", "--port", "1234"}
	parsedArgs, _ := parseCmdArgs(args)
	flagReturned, _ := validatePortFlag(parsedArgs)
	var expectedFlag = "1234"
	assert.Equal(t, expectedFlag, flagReturned)
}

// Unit test (negative test) for validatePortFlag()
func TestParseWrongPort(t *testing.T) {
	// setup the test server with arguments
	args = []string{"--data", "/dev/null", "--port", "123456789"}
	parsedArgs, _ := parseCmdArgs(args)
	flagReturned, flagError := validatePortFlag(parsedArgs)
	var expectedFlag = "0"
	var expectedError = "invalid port number entered"
	assert.Equal(t, expectedFlag, flagReturned)
	assert.Equal(t, expectedError, flagError.Error())
}

// Unit tests for createStorer()
func TestCreateStorerWithCorrectArguments(t *testing.T) {
	tests := []struct {
		testName string
		storer   string
		err      error
	}{
		{testName: "map storer", storer: "map", err: nil},
		{testName: "sync storer", storer: "sync", err: nil},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			_, actualErr := createStorer(tc.storer)
			assert.Equal(t, tc.err, actualErr)
		})
	}
}

func TestCreateStorerWithWrongArguments(t *testing.T) {
	tests := []struct {
		testName string
		storer   string
	}{
		{testName: "wrong storer", storer: "blah"},
		{testName: "empty storer", storer: ""},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			_, actualErr := createStorer(tc.storer)
			assert.Error(t, actualErr)
		})
	}
}

// Unit test for createPuppies()
func TestSavePuppiesToStore(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	var puppies []puppy.Puppy
	puppies = append(puppies, puppy.Puppy{ID: 1, Breed: "Jack Russell Terrier", Color: "White and Brown", Value: "1500"})
	storer := puppy.NewMapStore()
	returnedPuppyStorer, _ := createPuppies(storer, puppies)

	expected := "&{map[1:{1 Jack Russell Terrier White and Brown 1500}] 2 {0 0}}\n"
	fmt.Fprintln(out, returnedPuppyStorer)
	actual := buf.String()
	assert.Equal(t, expected, actual)
}

func TestSavePuppiesToStoreWithErr(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	var puppies []puppy.Puppy
	puppies = append(puppies, puppy.Puppy{ID: 1, Breed: "Jack Russell Terrier", Color: "White and Brown", Value: "-1500"})
	storer := puppy.NewMapStore()
	returnedPuppyStorer, _ := createPuppies(storer, puppies)

	expected := "<nil>\n"
	fmt.Fprintln(out, returnedPuppyStorer)
	actual := buf.String()
	assert.Equal(t, expected, actual)
}

// Unit test for parseJSONPuppies()
func TestParseJSONPuppiesSuccessfully(t *testing.T) {
	input := strings.NewReader(
		`[{"ID": 1, "Breed": "Jack Russell Terrier","Color": "White and Brown","Value": "1500"}]`)
	puppies, err := parseJSONPuppies(input)
	expected := []puppy.Puppy{{
		ID: 0x1, Breed: "Jack Russell Terrier", Color: "White and Brown", Value: "1500"}}
	assert.Equal(t, expected, puppies)
	assert.Equal(t, nil, err)
}

func TestParseBadJSONPuppies(t *testing.T) {
	input := strings.NewReader(`[[{"ID": "1", "Breed": 1, "Color": false, "Value": 2}]`)
	puppies, err := parseJSONPuppies(input)
	expected := []puppy.Puppy(nil)
	assert.Equal(t, expected, puppies)
	assert.Error(t, err)
}

func TestParseEmptyJSONPuppies(t *testing.T) {
	input := strings.NewReader(`dfsdfdsfsfs`)
	puppies, err := parseJSONPuppies(input)
	expected := []puppy.Puppy(nil)
	assert.Equal(t, expected, puppies)
	assert.Error(t, err)
}

func TestMain(m *testing.M) {
	// setup the test server with arguments
	args = []string{"--data", "/dev/null", "--port", "8888", "--store", "map"}
	os.Exit(m.Run())
}

func TestMainHappyPath(t *testing.T) {
	args = []string{""}
	assert.NotPanics(t, main)
	os.Exit(0)
}
