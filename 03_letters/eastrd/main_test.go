package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("a:2\nb:1\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestLetters(t *testing.T) {
	expected := map[rune]int{[]rune("c")[0]: 3, []rune("a")[0]: 4, []rune("0")[0]: 5}
	actual := letters("0a0ca0cac0a0")

	if len(expected) != len(actual) {
		t.Errorf("Unexpected output in main()")
	}
	for k, v := range actual {
		v2 := expected[k]
		if v2 != v {
			t.Errorf("Unexpected output in main()")
		}
	}
}

func TestLetters2(t *testing.T) {
	expected := map[rune]int{}
	actual := letters("")

	if len(expected) != len(actual) {
		t.Errorf("Unexpected output in main()")
	}
	for k, v := range actual {
		v2 := expected[k]
		if v2 != v {
			t.Errorf("Unexpected output in main()")
		}
	}
}

func TestSortLetters(t *testing.T) {
	expected := []string{"1:3", "a:2", "b:1"}
	actual := sortLetters(map[rune]int{[]rune("a")[0]: 2, []rune("b")[0]: 1, []rune("1")[0]: 3})

	if len(expected) != len(actual) || reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected output in main()")
	}
}

func TestSortLetters2(t *testing.T) {
	expected := []string{"1:10"}
	actual := sortLetters(map[rune]int{[]rune("1")[0]: 10})

	if len(expected) != len(actual) {
		t.Errorf("Unexpected output in main()")
	}
	for idx := range expected {
		if expected[idx] != actual[idx] {
			t.Errorf("Unexpected output in main()")
		}
	}
}

func TestSortLetters3(t *testing.T) {
	expected := []string{}
	actual := sortLetters(map[rune]int{})

	if len(expected) != len(actual) {
		t.Errorf("Unexpected output in main()")
	}
	for idx := range expected {
		if expected[idx] != actual[idx] {
			t.Errorf("Unexpected output in main()")
		}
	}
}
