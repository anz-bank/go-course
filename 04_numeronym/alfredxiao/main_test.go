package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumeronym(t *testing.T) {
	assert.Equal(t, "", numeronym(""))
	assert.Equal(t, "a", numeronym("a"))
	assert.Equal(t, "ab", numeronym("ab"))
	assert.Equal(t, "abc", numeronym("abc"))
	assert.Equal(t, "a2d", numeronym("abcd"))
	assert.Equal(t, "i18n", numeronym("internationalisation"))
}

func TestNumeronyms(t *testing.T) {
	assert.Equal(t, []string{}, numeronyms())
	assert.Equal(t, []string{"a"}, numeronyms("a"))
	assert.Equal(t, []string{"a", "", "a2d"}, numeronyms("a", "", "abcd"))
}

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	assert.Equal(t, "[a2d]", buf.String())
}
