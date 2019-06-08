package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var stdout = os.Stdout

var lettersTests = []struct {
	input       string
	expectedMap map[rune]int
	expectedStr []string
}{
	{"aba", map[rune]int{'a': 2, 'b': 1}, []string{"a:2", "b:1"}},
	{"ywbgjzhoiumsnqdexprlctvfka", map[rune]int{
		'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1, 'g': 1, 'h': 1, 'i': 1,
		'j': 1, 'k': 1, 'l': 1, 'm': 1, 'n': 1, 'o': 1, 'p': 1, 'q': 1, 'r': 1,
		's': 1, 't': 1, 'u': 1, 'v': 1, 'w': 1, 'x': 1, 'y': 1, 'z': 1},
		[]string{"a:1", "b:1", "c:1", "d:1", "e:1", "f:1", "g:1", "h:1", "i:1",
			"j:1", "k:1", "l:1", "m:1", "n:1", "o:1", "p:1", "q:1", "r:1",
			"s:1", "t:1", "u:1", "v:1", "w:1", "x:1", "y:1", "z:1"}},
	{`.#jV#Rz;ZO`, map[rune]int{
		'#': 2, '.': 1, ';': 1, 'O': 1, 'R': 1, 'V': 1, 'Z': 1, 'j': 1, 'z': 1},
		[]string{"#:2", ".:1", ";:1", "O:1", "R:1", "V:1", "Z:1", "j:1", "z:1"}},
}

func TestMain(t *testing.T) {
	r, w := captureStart()
	main()
	result := captureStop(r, w)
	expected := strings.Join(lettersTests[0].expectedStr, "\n") + "\n"
	require.Equal(t, expected, result)
}

func TestLetters(t *testing.T) {
	for _, test := range lettersTests {
		require.Equal(t, test.expectedMap, letters(test.input))
	}
}

func TestSortLetters(t *testing.T) {
	for _, test := range lettersTests {
		require.Equal(t, test.expectedStr, sortLetters(test.expectedMap))
	}
}

// captureStart diverts stdio to another file object
func captureStart() (io.Reader, io.Closer) {
	r, w, _ := os.Pipe()
	os.Stdout = w
	return r, w
}

// captureStop copies a file buffer and returns a string of the file
func captureStop(r io.Reader, w io.Closer) string {
	var buf bytes.Buffer
	w.Close()
	os.Stdout = stdout
	_, err := io.Copy(&buf, r)
	if err != nil {
		panic("file not opened")
	}
	return buf.String()
}
