package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	want := `Created puppy              : main.Puppy{ID:1, Breed:"Dogo", Colour:"White", Value:500}
Created puppy              : main.Puppy{ID:1, Breed:"Fila", Colour:"Golden", Value:900} 
`
	var buf bytes.Buffer
	out = &buf
	main()
	got := buf.String()
	assert.Equal(t, want, got)

}
