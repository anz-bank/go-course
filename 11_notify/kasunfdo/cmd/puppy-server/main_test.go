package main

import (
	"errors"
	"os"
	"testing"

	"github.com/anz-bank/go-course/11_notify/kasunfdo/pkg/puppy/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	os.Args = []string{""}
	os.Exit(m.Run())
}

func TestParseArgsLong(t *testing.T) {
	args := []string{
		"--data", "../../data/data.json",
		"--port", "1234",
		"--store", "sync",
		"--lostsvc", "lostSvc_end-point",
	}
	config, err := parseArgs(args)
	require.NoError(t, err)
	assert.NotNil(t, config.dataFile)
	assert.Equal(t, 1234, config.port)
	assert.Equal(t, "sync", config.storeType)
	assert.Equal(t, "lostSvc_end-point", config.lostSvcURL)
}

func TestParseArgsShort(t *testing.T) {
	args := []string{"-d", "../../data/data.json", "-p", "1234", "-s", "sync", "-l", "lostSvc_end-point"}
	config, err := parseArgs(args)
	require.NoError(t, err)
	assert.NotNil(t, config.dataFile)
	assert.Equal(t, 1234, config.port)
	assert.Equal(t, "sync", config.storeType)
	assert.Equal(t, "lostSvc_end-point", config.lostSvcURL)
}

func TestParseArgsWrong(t *testing.T) {
	args := []string{"-f", "../../data/data.json", "-r", "1234", "-t", "sync"}
	_, err := parseArgs(args)
	assert.Error(t, err)
}

func TestParseArgsDefault(t *testing.T) {
	args := []string{"-d", "../../data/data.json", "-l", "lostSvc_end-point"}
	config, err := parseArgs(args)
	require.NoError(t, err)
	assert.NotNil(t, config.dataFile)
	assert.Equal(t, 8080, config.port)
	assert.Equal(t, "map", config.storeType)
}

func TestCreateStore(t *testing.T) {
	storer, err := createStore("map")
	require.NoError(t, err)
	_, ok := storer.(*store.MapStore)
	assert.True(t, ok)

	storer, err = createStore("sync")
	require.NoError(t, err)
	_, ok = storer.(*store.SyncStore)
	assert.True(t, ok)

	storer, err = createStore("foo")
	require.Error(t, err)
	require.Nil(t, storer)
}

func TestCreatePuppies(t *testing.T) {
	s := store.NewMapStore()
	dataFile, err := os.Open("../../data/data.json")
	require.NoError(t, err)
	err = createPuppies(dataFile, s)
	assert.NoError(t, err)
}

type mockInvalidReader struct{}

func (mockInvalidReader) Read(_ []byte) (n int, err error) {
	return 0, errors.New("mockInvalidReader:error")
}

func TestCreatePuppiesReadFail(t *testing.T) {
	dataFile := mockInvalidReader{}
	err := createPuppies(dataFile, nil)
	assert.Error(t, err)
}
