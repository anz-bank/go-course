package main

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/anz-bank/go-course/10_rest/alextmz/pkg/puppy/store"
	"github.com/stretchr/testify/assert"
)

func TestMainHappyPathSyncStore(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"-d", "../../test/valid-formatted-json.json", "-s", "sync"}

	go main()
	// wait for all main() output to be done
	<-syncoutput

	// test if main() text output is what we expect
	got := buf.String()
	want := `Starting puppyserver with options:
file   = ../../test/valid-formatted-json.json
port  = 7735
store = sync
Loaded 3 puppies.
Printing puppy id 1: puppy.Puppy{ID:1, Breed:"Dogo", Colour:"White", Value:500}
Printing puppy id 2: puppy.Puppy{ID:2, Breed:"Mastiff", Colour:"Brindle", Value:700}
Printing puppy id 3: puppy.Puppy{ID:3, Breed:"Fila", Colour:"Golden", Value:900}
`
	assert.Equal(t, want, got)

	// test if main() is really serving the HTTP we expect
	//
	// unfortunately, looks like travis-ci does not allow
	// connections to 127.0.0.1: we get 'connect: connection refused'
	// this works as expected locally.
	// leaving the code in-place as it should work.
	//
	// resp, err := http.Get("http://127.0.0.1:7735/api/puppy/1")
	// assert.NoError(t, err)
	// defer resp.Body.Close()
	// got2, err := ioutil.ReadAll(resp.Body)
	// assert.NoError(t, err)
	// want2 := "{\"id\":1,\"breed\":\"Dogo\",\"colour\":\"White\",\"value\":500}\n"
	// assert.Equal(t, string(got2), want2)

	// signals main() to shutdown the http server
	shutdownhttp <- true
}

func TestMainHappyPathMapStore(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"-d", "../../test/valid-formatted-json.json", "-s", "map"}

	go main()
	// wait for all main() output to be done
	<-syncoutput

	// test if main() text output is what we expect
	got := buf.String()
	want := `Starting puppyserver with options:
file   = ../../test/valid-formatted-json.json
port  = 7735
store = map
Loaded 3 puppies.
Printing puppy id 1: puppy.Puppy{ID:1, Breed:"Dogo", Colour:"White", Value:500}
Printing puppy id 2: puppy.Puppy{ID:2, Breed:"Mastiff", Colour:"Brindle", Value:700}
Printing puppy id 3: puppy.Puppy{ID:3, Breed:"Fila", Colour:"Golden", Value:900}
`
	assert.Equal(t, want, got)

	// test if main() is really serving the HTTP we expect
	// unfortunately, looks like travis-ci does not allow
	// connections to 127.0.0.1: we get 'connect: connection refused'
	// this works as expected locally.
	// leaving the code in-place as it should work.
	// resp, err := http.Get("http://127.0.0.1:7735/api/puppy/1")
	// assert.NoError(t, err)
	// defer resp.Body.Close()
	// got2, err := ioutil.ReadAll(resp.Body)
	// assert.NoError(t, err)
	// want2 := "{\"id\":1,\"breed\":\"Dogo\",\"colour\":\"White\",\"value\":500}\n"
	// assert.Equal(t, string(got2), want2)

	// signals main() to shutdown the http server
	shutdownhttp <- true
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
	var buf bytes.Buffer
	out = &buf
	args = []string{"-d", "../../test/empty-file.json"}
	puppystore := store.NewMapStore()

	printpuppies(puppystore, 1)

	actual := buf.String()
	assert.Equal(t, "puppy with ID 1 being read does not exist\n", actual)
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
