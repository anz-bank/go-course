package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestLetters(t *testing.T) {
	var sentence = "mary had a little lamb"
	expected := map[rune]int{
		'm': 2,
		'a': 4,
		'r': 1,
		'y': 1,
		'h': 1,
		'd': 1,
		'i': 1,
		'l': 3,
		't': 2,
		'e': 1,
		'b': 1,
		' ': 4,
	}

	actual := letters(sentence)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestNoLetters(t *testing.T) {
	var sentence = ""

	actual := letters(sentence)

	if len(actual) > 0 {
		t.Errorf("actual map %v, is NOT empty", actual)
	}
}

func TestSorter(t *testing.T) {
	testMap := map[rune]int{
		'å':    -1,
		'a':    93,
		'ß':    100000,
		'\x00': 0,
	}
	expected := []string{
		"\x00:0", "a:93", "ß:100000", "å:-1",
	}
	actual := sortLetters(testMap)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestEmptySorter(t *testing.T) {
	testMap := map[rune]int{}

	actual := sortLetters(testMap)

	if len(actual) > 0 {
		t.Errorf("actual map %v, is NOT empty", actual)
	}
}

func TestMain(t *testing.T) {

	expected := "a:2\nb:1\n"
	var buf bytes.Buffer
	out = &buf
	main()
	actual := buf.String()

	if actual != expected {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
