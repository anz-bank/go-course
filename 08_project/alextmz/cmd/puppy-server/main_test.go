package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	want := `Creating puppy on Mapstore  : Created puppy : puppy.Puppy{ID:1, Breed:"Dogo", Colour:"White", Value:500}
Reading puppy back          : puppy.Puppy{ID:1, Breed:"Dogo", Colour:"White", Value:500}
Creating puppy on SyncStore : Created puppy : puppy.Puppy{ID:1, Breed:"Fila", Colour:"Golden", Value:900} 
Reading puppy back          : puppy.Puppy{ID:1, Breed:"Fila", Colour:"Golden", Value:900}
`

	t.Run("main test", func(t *testing.T) {
		var buf bytes.Buffer
		out = &buf
		main()
		got := buf.String()
		assert.Equal(t, want, got)
	})
}
