package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/anz-bank/go-course/09_json/mohitnag/pkg/puppy"
	"github.com/anz-bank/go-course/09_json/mohitnag/pkg/puppy/store"
	"github.com/stretchr/testify/assert"
)

const (
	defaultPuppy     = "./../../pkg/puppy/store/testdata/default-puppy.json"
	invalidPuppyJSON = "./../../pkg/puppy/store/testdata/invalid-puppy.json"
	duplicatePuppies = "./../../pkg/puppy/store/testdata/duplicate-puppies.json"
	manyPuppies      = "./../../pkg/puppy/store/testdata/many-puppies.json"
)

func TestMain(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	out = &buf
	args = []string{"--data", manyPuppies}
	main()
	expected := "{1 dog white 2}\n{1 dog white 2}\n"
	actual := buf.String()
	assert.Equal(expected, actual)
}

func TestMainError(t *testing.T) {
	assert := assert.New(t)
	args = []string{"-d", duplicatePuppies}
	assert.Panics(main)
}

func TestReadFile(t *testing.T) {
	assert := assert.New(t)
	expect := puppy.Puppy{ID: 1, Breed: "dog", Colour: "white", Value: "2"}
	actual := puppy.Puppy{}
	buff := readFile(defaultPuppy)
	err := json.Unmarshal(buff, &actual)
	assert.NoError(err)
	assert.Equal(expect, actual)
}

func TestReadFileBadPath(t *testing.T) {
	assert := assert.New(t)
	assert.Panics(func() { readFile("bad path") })
}

func TestUnmarshallingError(t *testing.T) {
	assert := assert.New(t)
	args = []string{"--data", invalidPuppyJSON}
	assert.Panics(main)
}

func TestInitialisePuppyMapStoreError(t *testing.T) {
	assert := assert.New(t)
	puppy := puppy.Puppy{ID: 1, Breed: "dog", Colour: "white", Value: "1"}
	mapStore := store.MapStore{}
	syncStore := store.SyncStore{}
	_ = mapStore.CreatePuppy(puppy)
	err := initialisePuppyStore(&mapStore, &syncStore, duplicatePuppies)
	assert.Error(err)
}

func TestInitialisePuppySyncStoreError(t *testing.T) {
	assert := assert.New(t)
	puppy := puppy.Puppy{ID: 1, Breed: "dog", Colour: "white", Value: "1"}
	mapStore := store.MapStore{}
	syncStore := store.SyncStore{}
	_ = syncStore.CreatePuppy(puppy)
	err := initialisePuppyStore(&mapStore, &syncStore, duplicatePuppies)
	assert.Error(err)
}
