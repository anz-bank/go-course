package main

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"

	"bou.ke/monkey"

	puppy "github.com/anz-bank/go-course/10_rest/n0npax/pkg/puppy"
	store "github.com/anz-bank/go-course/10_rest/n0npax/pkg/puppy/store"
	"github.com/stretchr/testify/assert"
)

const (
	testData              = `[{"id":0,"value":4,"breed":"Type: D","colour":"Red"},{"id":0,"value":3,"breed":"Type: U","colour":"White"},{"id":0,"value":2,"breed":"Type: P","colour":"Green"},{"id":0,"value":1,"breed":"Type: A","colour":"Blue"}]`
	corruptedTestData     = `[{"id":0,"value":"LLAMA","breed":"Type: U","colour":"White"}]`
	negativeValueTestData = `[{"id":0,"value":-44,"breed":"Type: U","colour":"White"}]`
)

func TestMain(t *testing.T) {
	fakeExit := func(int) {
		panic("foo-arg")
	}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	os.Args = []string{"", "-s", "map", "-d", "/dev/null"}
	assert.PanicsWithValuef(t, "foo-arg", main, "foo-arg")
}

func TestMainFakeArgs(t *testing.T) {
	parser = func([]string) (config, error) {
		return config{}, errors.New("test")
	}
	logFatalf = func(...interface{}) {
		panic("test")
	}
	assert.Panics(t, main)
}

func TestRunPuppyServerBadFile(t *testing.T) {
	pf, err := os.Open("/dev/null")
	assert.NoError(t, err)
	c := config{port: 8888, sType: "map", puppyFile: pf}
	err = runPuppyServer(c)
	assert.Error(t, err)
}

func TestRunPuppyServerBadStorerType(t *testing.T) {
	c := config{port: 8888, sType: "foo"}
	err := runPuppyServer(c)
	assert.Error(t, err)
}

func TestRunPuppyServer(t *testing.T) {

	fakeRun := func(addr ...string) (err error) {
		panic("foo")
	}
	patch := monkey.Patch(puppy.RestBackend(store.NewMemStore()).Run, fakeRun)
	defer patch.Unpatch()
	file := strings.NewReader(testData)
	c := config{port: 8888, sType: "map", puppyFile: file, storer: store.NewMemStore()}
	err := runPuppyServer(c)
	assert.NoError(t, err)
}

func TestCreateStorer(t *testing.T) {
	stores := []string{"map", "sync"}
	for _, v := range stores {
		v := v
		t.Run(v, func(t *testing.T) {
			s, err := createStorer(v)
			assert.NoError(t, err)
			_, ok := s.(puppy.Storer)
			assert.True(t, ok)
		})
	}
}

func TestCreateNotSupportedStorer(t *testing.T) {
	stores := []string{"foo", "bar"}
	for _, v := range stores {
		v := v
		t.Run(v, func(t *testing.T) {
			_, err := createStorer(v)
			assert.Error(t, err)
		})
	}
}

func TestExitOnError(t *testing.T) {
	fakeExit := func(int) {}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	var buf bytes.Buffer
	out = &buf

	exitOnError(errors.New("test"))
	actual := buf.String()
	assert.Equal(t, "test\n", actual)
}

func TestFeedStore(t *testing.T) {
	file := strings.NewReader(testData)
	s := store.NewMemStore()
	c := config{storer: s, puppyFile: file}
	err := feedStorer(c)
	assert.NoError(t, err)
}

func TestFeedStoreCorruptedData(t *testing.T) {
	file := strings.NewReader(corruptedTestData)
	s := store.NewMemStore()
	c := config{storer: s, puppyFile: file}
	err := feedStorer(c)
	assert.Error(t, err)
}

func TestFeedStoreNegativeVal(t *testing.T) {
	file := strings.NewReader(negativeValueTestData)
	s := store.NewMemStore()
	c := config{storer: s, puppyFile: file}
	err := feedStorer(c)
	assert.Error(t, err)
}

func TestReadPuppiesDevNull(t *testing.T) {
	file, err := os.Open("/dev/null")
	assert.NoError(t, err)
	_, err = readPuppies(file)
	assert.Error(t, err)

}

func TestReadPuppiesFromEtcHosts(t *testing.T) {
	file, err := os.Open("/etc/hosts")
	assert.NoError(t, err)
	_, err = readPuppies(file)
	assert.Error(t, err)
}

func TestReadPuppiesNil(t *testing.T) {
	_, err := readPuppies((*os.File)(nil))
	assert.NoError(t, err)
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}
func TestReadPuppiesBrokenRead(t *testing.T) {
	r := errReader(1)
	_, err := readPuppies(r)
	assert.Error(t, err)
}

func TestMainReadPuppies(t *testing.T) {
	file, err := os.Open("../../test-data/data.json")
	assert.NoError(t, err)
	puppies, err := readPuppies(file)
	assert.NoError(t, err)
	for _, p := range puppies {
		assert.Equal(t, 0, p.ID)
	}
}

func TestParseArgs(t *testing.T) {
	args := []string{"-s", "map", "-d", "/dev/null", "--port", "1234"}
	config, err := parseArgs(args)
	assert.NoError(t, err)
	assert.Equal(t, config.port, 1234)
	assert.Equal(t, config.sType, "map")
}
