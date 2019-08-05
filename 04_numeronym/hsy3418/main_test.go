package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestNumeronym(t *testing.T) {
	actual := numeronyms("accessibility", "Kubernetes", "abc", "Android Enterprise", "", "ab", "a", "34teaf1223%783!",
		"😄🐷🙈🐷🏃😄", "😄🏃😄", "阿拉奥利嗷嗷嗷", "阿拉奥")
	expected := []string{"a11y", "K8s", "abc", "A16e", "", "ab", "a", "313!", "😄4😄", "😄🏃😄",
		"阿5嗷", "阿拉奥"}
	for index, val := range actual {
		if expected[index] != val {
			t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected[index], val)
		}
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	expected := strconv.Quote("[a11y K8s abc]\n")
	actual := strconv.Quote(buf.String())
	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}
