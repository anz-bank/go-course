package store

import (
	"os"
	"testing"

	types "github.com/anz-bank/go-training/08_project/mohankrishna/pkg/mohankrishna-puppy"
	tassert "github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	// Given
	assert := tassert.New(t)
	pup := types.Puppy{
		ID:     0x1236,
		Breed:  "Sheep herder",
		Colour: "Brown",
		Value:  1000,
	}
	levelDBStore := NewLevelDBStore()
	defer cleanUp(levelDBStore)
	err := levelDBStore.CreatePuppy(&pup)

	assert.NoError(err, "Should be able to create puppy")

	actual, err := levelDBStore.GetAll()
	expected := []*types.Puppy{&pup}
	if assert.NoError(err, "Should be able to read all the data") {
		assert.EqualValuesf(actual, expected, "Read data should be identical to the one passed to Create")
	}
}

func cleanUp(levelDBStore *LevelDBStore) {
	iter := levelDBStore.ldb.NewIterator(nil, nil)
	for iter.Next() {
		_ = levelDBStore.ldb.Delete(iter.Key(), nil)
	}
	iter.Release()
	levelDBStore.CloseDB()
	os.RemoveAll("./storage/")
}
