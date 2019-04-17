package fib

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

var out io.Writer = os.Stdout

func TestFibWithNEqualsToZero(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	Fib(0)

	//then
	r.Equalf("", buf.String(), "Unexpected output in main()")
}

func TestFibWithNEqualsToOne(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	Fib(1)

	//then
	r.Equalf("", buf.String(), "Unexpected output in main()")
}

func TestFibWithNEqualsToSeven(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	Fib(7)

	//then
	r.Equalf("", buf.String(), "Unexpected output in main()")
}

func TestFibWithNEqualsToOneHundred(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	Fib(100)

	//then
	r.Equalf("", buf.String(), "Unexpected output in main()")
}
