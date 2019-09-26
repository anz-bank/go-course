package main

import (
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"

	puppy "github.com/anz-bank/go-course/10_rest/runnerdave/pkg/puppy"
	store "github.com/anz-bank/go-course/10_rest/runnerdave/pkg/puppy/store"
	"github.com/stretchr/testify/assert"
)

var (
	puppy1 = func() puppy.Puppy {
		return puppy.Puppy{ID: 1, Breed: "Chihuahua", Colour: "Brown", Value: 12.30}
	}
	puppy2 = func() puppy.Puppy {
		return puppy.Puppy{ID: 2, Breed: "Cacri", Colour: "Undefined", Value: 1.30}
	}
	puppy3 = func() puppy.Puppy {
		return puppy.Puppy{ID: 12, Breed: "Imaginary", Colour: "Undefined", Value: -1.30}
	}
)

func TestMain(m *testing.M) {
	args = []string{}
	os.Exit(m.Run())
}

func TestLoad(t *testing.T) {
	s := store.NewSyncStore()

	puppies := []puppy.Puppy{puppy1(), puppy2()}
	err := load(s, puppies)
	if err != nil {
		panic(err)
	}

	readable := strconv.Quote(`{"breed":"Chihuahua","color":"Brown","id":1,"value":12.3} 
	{"breed":"Cacri","color":"Undefined","id":2,"value":1.3} `)
	expected := strings.ReplaceAll(readable, "\\n\\t", "")
	actual := strconv.Quote(s.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestLoadInvalidValuePuppy(t *testing.T) {
	s := store.NewSyncStore()

	puppies := []puppy.Puppy{puppy1(), puppy3()}
	err := load(s, puppies)

	expected := puppy.Error{Code: 0x1, Message: "puppy has invalid value (-1.300000)"}
	assert.Equal(t, &expected, err, "Not invalid error")
}

func TestNewStoreInvalidStorage(t *testing.T) {
	s := newStore("blah")
	assert.Equal(t, nil, s, "Not invalid storage error")
}

func TestNewStoreSyncStorage(t *testing.T) {
	s := newStore("sync")
	expected := store.NewSyncStore()
	assert.Equal(t, expected, s, "Not a sync storage")
}

func TestMainOutputNoValueForFlag(t *testing.T) {
	args = []string{"-d", " "}
	assert.Panics(t, func() { main() }, "Main with no value for flag did not panic")
}

func TestMainOutputInvalidData(t *testing.T) {
	args = []string{"--data", "../../puppydata/invalid-data.json"}
	assert.Panics(t, func() { main() }, "Main with invalid data did not panic")
}

func TestMainInvalidStorage(t *testing.T) {
	args = []string{"-s", "bad", "--data", "../../puppydata/data.json"}
	assert.Panics(t, func() { main() }, "Main with invalid storage did not panic")
}

func TestMainInvalidStorageLoad(t *testing.T) {
	args = []string{"--data", "../../puppydata/invalid-puppies-negative-value.json"}
	assert.Panics(t, func() { main() }, "Main with invalid storage did not panic")
}

func TestMainOutputInvalidPuppy(t *testing.T) {
	args = []string{"--data", "../../puppydata/invalid-puppies.json"}
	assert.Panics(t, func() { main() }, "Main with invalid data did not panic")
}

func TestInvalidPort(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	args = []string{"--data", "../../puppydata/data.json", "-s", "map",
		"-p", "165535"}

	main()
	readable := strconv.Quote(`Store of puppies:
	&{map[1:{ULTRASURE green 1 2732.81} 
	2:{EPLODE brown 2 3889.92} 
	3:{SENTIA blue 3 1472.29} 
	4:{MAKINGWAY brown 4 2254.7} 
	5:{VISUALIX brown 5 3250.57} 
	6:{SPLINX brown 6 2467.04} 
	7:{POWERNET blue 7 3970.65}] 8}
	Could not start puppy server: listen tcp: 
	address 165535: invalid port`)
	expected := strings.ReplaceAll(readable, "\\n\\t", "")
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}
