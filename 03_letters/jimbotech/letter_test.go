package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestLetters(t *testing.T) {
	var sentence = "marry had a little lamb"
	expected := map[rune]int{
		'm': 2,
		'a': 4,
		'r': 2,
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

	res := letters(sentence)

	eq := reflect.DeepEqual(res, expected)
	if !eq {
		t.Errorf("Result maps %v differs from expected %v", res, expected)
	}
}

func TestNoLetters(t *testing.T) {
	var sentence = ""

	res := letters(sentence)

	if len(res) > 0 {
		t.Error("Retured map is NOT empty")
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
	res := sortLetters(testMap)

	eq := reflect.DeepEqual(res, expected)
	if !eq {
		t.Errorf("Result %v differs from expected %v", res, expected)
	}
}
func TestEmptySorter(t *testing.T) {
	testMap := map[rune]int{}

	res := sortLetters(testMap)

	if len(res) > 0 {
		t.Error("Retured map is NOT empty")
	}
}

func TestMain(t *testing.T) {

	want := "a:2\nb:1\n"
	var buf bytes.Buffer
	out = &buf
	main()
	result := buf.String()

	if result != want {
		t.Errorf("expected %v, got %v", want, result)
	}
}
