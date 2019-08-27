package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/anz-bank/go-course/09_json/mohitnag/pkg/puppy"
	"github.com/anz-bank/go-course/09_json/mohitnag/pkg/puppy/store"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	out = &buf
	args = []string{"--data", "./../../pkg/puppy/store/testdata/many-puppies.json"}
	main()
	expected := "{1 dog white 2}\n{1 dog white 2}\n"
	actual := buf.String()
	assert.Equal(expected, actual)
}

func TestMainError(t *testing.T) {
	assert := assert.New(t)
	args = []string{"-d", "./../../pkg/puppy/store/testdata/invalid-puppies.json"}
	assert.Panics(main)
}

func TestReadJson(t *testing.T) {
	t.Run("Valid Json File should pass", func(t *testing.T) {
		assert := assert.New(t)
		expect := puppy.Puppy{ID: 1, Breed: "dog", Colour: "white", Value: "2"}
		actual := puppy.Puppy{}
		buff := readJSON("./../../pkg/puppy/store/testdata/default-puppy.json")
		err := json.Unmarshal([]byte(buff), &actual)
		assert.NoError(err)
		assert.Equal(expect, actual)
	})
	t.Run("Non existing file should fail", func(t *testing.T) {
		assert := assert.New(t)
		assert.Panics(func() { readJSON("bad path") })
	})
}

func TestUnmarshallingError(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	out = &buf
	args = []string{"--data", "./../../pkg/puppy/store/testdata/invalid-puppy.json"}
	assert.Panics(main)
}

func TestInitialisePuppyStoreErrors(t *testing.T) {
	assert := assert.New(t)
	puppy := puppy.Puppy{ID: 1, Breed: "dog", Colour: "white", Value: "1"}
	mapStore := store.MapStore{}
	syncStore := store.SyncStore{}
	_ = mapStore.CreatePuppy(puppy)
	_ = syncStore.CreatePuppy(puppy)
	err := initialisePuppyStore(&mapStore, &syncStore, "./../../pkg/puppy/store/testdata/invalid-puppies.json")
	assert.Error(err)
	newmapStore := store.MapStore{}
	err = initialisePuppyStore(&newmapStore, &syncStore, "./../../pkg/puppy/store/testdata/invalid-puppies.json")
	assert.Error(err)
}
