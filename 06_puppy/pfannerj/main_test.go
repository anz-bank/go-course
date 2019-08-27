package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	mainout = &buf
	main()
	expected := `Map Puppy Created with ID: 1
Map Puppy read: {1 Labrador Brown 999.99}
Map Puppy updated: 1
Map Puppy deleted! Error returned was: <nil>
Sync Puppy Created with ID: 1
Sync Puppy read: {1 Labrador Brown 999.99}
Sync Puppy updated: 1
Sync Puppy deleted! Error returned was: <nil>
`
	actual := buf.String()
	assert.Equal(t, expected, actual, "Unexpected output from main()")
}
