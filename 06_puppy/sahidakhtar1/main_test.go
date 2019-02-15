package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainOutPut(t *testing.T) {
	//Just not checking the console out put. Keeping this for code coverage
	//Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf
	var m Storer = newMapStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}

	//When
	main()
	p1ID := m.CreatePuppy(p1)
	p2 := m.ReadPuppy(p1ID)
	fmt.Println(p2)

	// Then
	expected := "{1 Bulldog White 100}\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
