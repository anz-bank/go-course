package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "1\n1\n2\n3\n5\n8\n13"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestFibonnaci1(t *testing.T) {
	//Given
	assert := assert.New(t)
	//when
	fibonnaciArray := Fibonacci(1)

	//Then
	assert.Equal([]int{1}, fibonnaciArray)

}
func TestFibonnaci2(t *testing.T) {
	//Given
	assert := assert.New(t)
	//when
	fibonnaciArray := Fibonacci(-1)

	//Then
	assert.Equal([]int{}, fibonnaciArray)

}
func TestFibonnaci3(t *testing.T) {
	//Given
	assert := assert.New(t)
	//when
	fibonnaciArray := Fibonacci(0)

	//Then
	assert.Equal([]int{}, fibonnaciArray)

}
