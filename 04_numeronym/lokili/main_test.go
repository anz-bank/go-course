package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var testSet = []struct {
	input  []string
	output []string
}{
	{[]string{"世界Thrives！"}, []string{"世8！"}},
	{[]string{""}, []string{""}},
	{[]string{"&《 44za  "}, []string{"&7 "}},
	{[]string{"？will他 reach the airport at 10 am ?"}, []string{"？33?"}},
}

func TestMainOutput(t *testing.T) {
	r := require.New(t)
	// Given
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := `[a11y K8s abc]`
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestNumeronyms(t *testing.T) {
	r := require.New(t)
	for _, data := range testSet {
		r.Equalf(data.output, numeronyms(data.input...), "Unexpected output in main()")
	}
}
