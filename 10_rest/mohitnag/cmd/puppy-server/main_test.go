package main

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/anz-bank/go-course/10_rest/mohitnag/pkg/puppy"

	"github.com/stretchr/testify/assert"
)

const (
	defaultPuppy     = "./../../pkg/puppy/store/testdata/default-puppy.json"
	invalidPuppyJSON = "./../../pkg/puppy/store/testdata/invalid-puppy.json"
	duplicatePuppies = "./../../pkg/puppy/store/testdata/duplicate-puppies.json"
	manyPuppies      = "./../../pkg/puppy/store/testdata/many-puppies.json"
)

func TestMainError(t *testing.T) {
	assert := assert.New(t)
	srvCh = make(chan *http.Server)
	args = []string{"-d", manyPuppies}
	go assert.Panics(main)
	srv := <-srvCh
	_ = srv.Shutdown(context.Background())
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

func TestInitialisePuppyStoreWithFile(t *testing.T) {
	assert := assert.New(t)
	s := createStore("map")
	err := initialisePuppyStoreWithFile(s, manyPuppies)
	actual, _ := s.ReadPuppy(1)
	assert.NoError(err)
	assert.Equal("white", actual.Colour)
}
func TestInitialisePuppyMapStoreError(t *testing.T) {
	assert := assert.New(t)
	s := createStore("map")
	err := initialisePuppyStoreWithFile(s, duplicatePuppies)
	assert.Error(err)
}

func TestInitialisePuppySyncStoreError(t *testing.T) {
	assert := assert.New(t)
	s := createStore("sync")
	err := initialisePuppyStoreWithFile(s, duplicatePuppies)
	assert.Error(err)
}
