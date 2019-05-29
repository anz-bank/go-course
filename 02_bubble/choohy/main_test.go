package main

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

func TestBubble(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`Bubble Sort!
[1 2 3 5]
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf(expected)
		t.Errorf(actual)
	}
}

func TestBubble0(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fmt.Fprintln(out, bubble([]int{0, 0, 0, 0}))

	expected := strconv.Quote(`[0 0 0 0]
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf(expected)
		t.Errorf(actual)
	}
}

func TestBubbleNegative(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fmt.Fprintln(out, bubble([]int{-1, -6, -3, -10}))

	expected := strconv.Quote(`[-10 -6 -3 -1]
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf(expected)
		t.Errorf(actual)
	}
}

func TestBubblePositiveNegative(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fmt.Fprintln(out, bubble([]int{-1, 6, 3, -10}))

	expected := strconv.Quote(`[-10 -1 3 6]
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf(expected)
		t.Errorf(actual)
	}
}

func TestBubbleDuplicates(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fmt.Fprintln(out, bubble([]int{6, 6, 0, 0}))

	expected := strconv.Quote(`[0 0 6 6]
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf(expected)
		t.Errorf(actual)
	}
}

func TestBubbleLongArray(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fmt.Fprintln(out, bubble([]int{7, -23, 70, 52, -12, 111, -6, 54, 33, -29, 0}))

	expected := strconv.Quote(`[-29 -23 -12 -6 0 7 33 52 54 70 111]
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf(expected)
		t.Errorf(actual)
	}
}
