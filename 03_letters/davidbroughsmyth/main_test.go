package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	poem1 = `There was movement at the station, for the word had passed around
    That the colt from old Regret had got away,
   And had joined the wild bush horses — he was worth a thousand pound,
    So all the cracks had gathered to the fray.
   All the tried and noted riders from the stations near and far
    Had mustered at the homestead overnight,
   For the bushmen love hard riding where the wild bush horses are,
    And the stock-horse snuffs the battle with delight.`

	poem2 = `大云寺赞公房四首 (一)
心在水精域
衣沾春雨时
洞门尽徐步
深院果幽期
到扉开复闭
撞钟斋及兹
醍醐长发性
饮食过扶衰
把臂有多日
开怀无愧辞
黄鹂度结构
紫鸽下罘罳
愚意会所适
花边行自迟
汤休起我病
微笑索题诗`
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	expected := strconv.Quote("a:2\nb:1\n")
	actual := strconv.Quote(buf.String())
	assert.Equal(t, expected, actual, "Unexpected output from main()")
}

var tests = map[string]struct {
	input       string
	wantLetters map[rune]int
	wantSort    []string
}{
	"Empty": {
		input:       "",
		wantLetters: map[rune]int{},
		wantSort:    []string{},
	},
	"doubles": {
		input:       "ZZaabbYYdd",
		wantLetters: map[rune]int{'Z': 2, 'd': 2, 'b': 2, 'a': 2, 'Y': 2},
		wantSort:    []string{"Y:2", "Z:2", "a:2", "b:2", "d:2"},
	},
	"Alphanumeric": {
		input:       "1x222B5jS0x",
		wantLetters: map[rune]int{'0': 1, '1': 1, '2': 3, '5': 1, 'j': 1, 'x': 2, 'B': 1, 'S': 1},
		wantSort:    []string{"0:1", "1:1", "2:3", "5:1", "B:1", "S:1", "j:1", "x:2"},
	},
	"SpecialChars": {
		input:       "...@$%^&*!!.",
		wantLetters: map[rune]int{'.': 4, '!': 2, '@': 1, '$': 1, '%': 1, '^': 1, '&': 1, '*': 1},
		wantSort:    []string{"!:2", "$:1", "%:1", "&:1", "*:1", ".:4", "@:1", "^:1"},
	},
	"Chinese": {
		input:       "大云寺赞大云寺赞",
		wantLetters: map[rune]int{'大': 2, '云': 2, '寺': 2, '赞': 2},
		wantSort:    []string{"云:2", "大:2", "寺:2", "赞:2"},
	},
	"symbols_emoji": {
		input:       "😸♞🁻∑👾👾😸😛😛😛",
		wantLetters: map[rune]int{'♞': 1, '🁻': 1, '∑': 1, '👾': 2, '😸': 2, '😛': 3},
		wantSort:    []string{"∑:1", "♞:1", "🁻:1", "👾:2", "😛:3", "😸:2"},
	},
}

func TestLetterFreq(t *testing.T) {
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			got := letterFreq(test.input)
			assert.Equal(t, test.wantLetters, got)
		})
	}
}

func TestLetter(t *testing.T) {
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			got := letters(test.input)
			assert.Equal(t, test.wantLetters, got)
		})
	}
}

func TestSortLetters(t *testing.T) {
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			got := sortLetters(test.wantLetters)
			assert.Equal(t, test.wantSort, got)
		})
	}
}

var testConc = map[string]struct {
	input       string
	wantLetters map[rune]int
}{
	"WordSpaces": {
		input:       "My Me  Make  More Max",
		wantLetters: map[rune]int{'M': 5, 'a': 2, 'e': 3, 'k': 1, 'o': 1, 'r': 1, 'x': 1, 'y': 1},
	},
	"MixedUnicode": {
		input:       "4大 4$ 4😸N  n😛😛n 4大",
		wantLetters: map[rune]int{'4': 4, '$': 1, 'n': 2, 'N': 1, '大': 2, '😸': 1, '😛': 2},
	},
}

func TestLetterConcurrency(t *testing.T) {
	for name, test := range testConc {
		test := test
		t.Run(name, func(t *testing.T) {
			got := letters(test.input)
			assert.Equal(t, test.wantLetters, got)
		})
	}
}

func BenchmarkLetterFreq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterFreq(poem1 + poem2)
	}
}

func BenchmarkConcurrentLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letters(poem1 + poem2)
	}
}
