package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/anz-bank/go-course/10_rest/nickolee/pkg/puppy/store"

	"github.com/stretchr/testify/assert"
)

// that is, no args were provided at all
func TestMainHappyPath(t *testing.T) {
	args = []string{"-d", "../../puppy-data/puppies.json", "-p", "80000", "-s", "map"}
	assert.NotPanics(t, main)
}

func TestLongFlag(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	args = []string{"--data", "../../puppy-data/puppies.json", "-p", "80000"}

	// purposely passing an invalid port number to cause main to panic in order to pass test
	// but in order for test to not panic out, we need to recover()
	func() {
		defer func() {
			_ = recover()
		}()
		main()
	}()

	expected := `Puppy with ID 1 has been created
Retrieved puppy: &{1 Vulpix Red 2900}
Puppy with ID 2 has been created
Retrieved puppy: &{2 Eevee Light Brown 1290}
Puppy with ID 3 has been created
Retrieved puppy: &{3 Vaporeon Sea Blue 3290}
`
	actual := buf.String()
	assert.Equal(t, expected, actual)
}

func TestParseError(t *testing.T) {
	args = []string{"--wrongFlag"}
	assert.Panics(t, main)
}

func TestEmptyFileName(t *testing.T) {
	args = []string{"--data"}
	assert.Panics(t, main)
}

func TestWrongFileName(t *testing.T) {
	args = []string{"--data", "iDontExist.json"}
	assert.Panics(t, main)
}

func TestUnmarshalPuppiesTypeMismatch(t *testing.T) {
	args = []string{"--data", "../../puppy-data/type_mismatch.json"}
	assert.Panics(t, main)
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

func TestReadFileAndUnmarshalPuppiesFunc(t *testing.T) {
	f, _ := os.Open("../../puppy-data/invalid.json")
	puppies, err := readFileAndUnmarshalPuppies(f)
	assert.Nil(t, puppies)
	assert.Error(t, err)
}
