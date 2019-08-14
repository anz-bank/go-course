package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumeronyms(t *testing.T) {
	testData := []struct {
		Scenario string
		input    []string
		expected []string
	}{
		{"Happy Day", []string{"ehi!@  45%!"}, []string{"e9!"}},
		{"Empty slice should return empty", []string{""}, []string{""}},
		{"String with 2 char", []string{"we"}, []string{"we"}},
		{"String with 3 char", []string{"win"}, []string{"win"}},
	}
	for _, td := range testData {
		td := td
		t.Run(td.Scenario, func(t *testing.T) {
			numeronym := numeronyms(td.input...)
			assert.Equal(t, td.expected, numeronym)
		})
	}
}
func TestMain(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	out = &buf
	main()
	expected := string("[a11y K8s abc]")
	actual := buf.String()
	actual = strings.Replace(actual, "\n", ",", -1)
	actual = strings.TrimRight(actual, ",")
	assert.Equal(expected, actual)
}

func TestNumeronym(t *testing.T) {
	testData := []struct {
		Scenario string
		input    string
		expected string
	}{
		{"Three letters", "abc", "abc"},
		{"Four letters", "test", "t2t"},
		{"Test Unicode", "🤓😀🤣🥳", "🤓2🥳"},
	}
	for _, td := range testData {
		td := td
		t.Run(td.Scenario, func(t *testing.T) {
			numeronym := numeronym(td.input)
			assert.Equal(t, td.expected, numeronym)
		})
	}
}
