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

func TestSingleNumeronymUnicode(t *testing.T) {
	assert.Equal(t, "", numeronym(""))
	assert.Equal(t, "\u3128", numeronym("\u3128"))
	assert.Equal(t, "\u3127\u3128", numeronym("\u3127\u3128"))
	assert.Equal(t, "\u31262\u3129", numeronym("\u3126\u3127\u3128\u3129"))
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
