package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMapStoreWithoutContructor may not be required as it may not be
// possible to create an instance of that type without using the constructor
func TestMapStoreWithoutContructor(t *testing.T) {
	var puppyStore MapStore
	pup := Puppy{1, "kelpie", "brown", "indispensable"}

	_, err := puppyStore.CreatePuppy(&pup)
	assert.Equal(t, ErrNotConstructed, err)
	err = puppyStore.UpdatePuppy(1, &pup)
	assert.Equal(t, ErrNotConstructed, err)
	_, err = puppyStore.ReadPuppy(1)
	assert.Equal(t, ErrNotConstructed, err)
	err = puppyStore.DeletePuppy(1)
	assert.Equal(t, ErrNotConstructed, err)
}
