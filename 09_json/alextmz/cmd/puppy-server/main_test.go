package main

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/anz-bank/go-course/09_json/alextmz/pkg/puppy/store"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var tests = map[string]struct {
		file     string
		expected string
	}{
		"Valid JSON": {
			file: "valid-formatted-json.json",
			expected: `Printing puppy id 1: puppy.Puppy{ID:1, Breed:"Dogo", Colour:"White", Value:500}
Printing puppy id 2: puppy.Puppy{ID:2, Breed:"Mastiff", Colour:"Brindle", Value:700}
Printing puppy id 3: puppy.Puppy{ID:3, Breed:"Fila", Colour:"Golden", Value:900}
`},
	}

	var buf bytes.Buffer
	out = &buf
	args = []string{"-d", "../../test/" + tests["Valid JSON"].file}

	main()

	actual := buf.String()
	assert.Equal(t, tests["Valid JSON"].expected, actual)
}

func TestNoFileGiven(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"-d", " "}

	main()

	actual := buf.String()
	assert.Equal(t, "error opening file: open  : no such file or directory\n", actual)
}

func TestInvalidFlag(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"-#"}

	main()

	actual := buf.String()
	assert.Equal(t, "error parsing command line: unknown short flag '-#'\n", actual)
}

func TestInvalidPuppyID(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"-d", "../../test/invalid-ids.json"}

	main()

	actual := buf.String()
	assert.Equal(t, "error storing puppies: puppy already initialized with ID -10\n", actual)
}

func TestInvalidJSON(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"-d", "../../test/empty-file.json"}

	main()

	actual := buf.String()
	assert.Equal(t, "error reading JSON file: unexpected end of JSON input\n", actual)
}

func TestPrintPuppiesErr(t *testing.T) {
	puppystore := store.NewMapStore()
	printpuppies(puppystore, 1)

	var buf bytes.Buffer
	out = &buf
	args = []string{"-d", "../../test/empty-file.json"}

	main()

	actual := buf.String()
	assert.Equal(t, "error reading JSON file: unexpected end of JSON input\n", actual)
}

func TestReadError(t *testing.T) {
	var r1 readerror
	_, err := readfile(r1)
	assert.Error(t, err)
}

func TestReadOk(t *testing.T) {
	r2 := newReadOk(`[{"breed":"Dogo","colour":"White","value":500}]`)
	_, err := readfile(r2)
	assert.NoError(t, err)
}

type readerror int

func (readerror) Read(p []byte) (n int, err error) {
	return 0, errors.New("always error")
}

type readok struct {
	s    string
	done bool
}

func (r *readok) Read(p []byte) (n int, err error) {
	if r.done {
		return 0, io.EOF
	}

	l := copy(p, r.s)
	r.done = true

	return l, nil
}

func newReadOk(s string) *readok {
	return &readok{s, false}
}
