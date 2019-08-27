package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputsToLetters(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(map[rune]int{}, letters(""), "empty string")
	assert.Equal(map[rune]int{'F': 1}, letters("F"), "letter")
	assert.Equal(map[rune]int{'\n': 1}, letters("\n"), "new line")
	assert.Equal(map[rune]int{'*': 1}, letters("*"), "symbol")
	assert.Equal(map[rune]int{'✌': 1}, letters("✌"), "emoticon")
}

func TestOutputToLetters(t *testing.T) {
	assert := assert.New(t)
	longString := "################################################"
	expectedOutputForMixedInput := map[rune]int{'H': 1, 'h': 1, 'e': 4, 'r': 2, ' ': 2, '!': 1, '2': 1}

	assert.Equal(map[rune]int{'#': 48}, letters(longString), "long string")
	assert.Equal(expectedOutputForMixedInput, letters("Here here 2!"), "mixed input")
}

func TestSortLetters(t *testing.T) {
	assert := assert.New(t)

	longInput := "5555555555555555bbbbbbbbbbbbbbbbbbb✌✌✌✌✌✌✌✌✌✌✌????????????"
	output1 := []string{" :8", "T:1", "a:1", "b:1", "c:1", "d:1", "e:3", "f:1", "g:1", "h:2", "i:1", "j:1", "k:1", "l:1"}
	output2 := []string{"m:1", "n:1", "o:4", "p:1", "q:1", "r:2", "s:1", "t:1", "u:2", "v:1", "w:1", "x:1", "y:1", "z:1"}
	outputManyChars := append(output1, output2...)

	assert.Equal([]string{"a:1", "b:1", "c:1", "d:1", "e:1"}, sortLetters(letters("edcba")), "unsorted input")
	assert.Equal([]string{"5:16", "?:12", "b:19", "✌:11"}, sortLetters(letters(longInput)), "long input")
	assert.Equal(outputManyChars, sortLetters(letters("The quick brown fox jumps over the lazy dog")), "many characters")
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	assert.Equal(t, "a:2\nb:1\n", buf.String())
}
