package main

import (
	"bytes"
	"testing"
)

func Test_main(t *testing.T) {
	want := "1\n1\n2\n3\n5\n8\n13\n"

	var buf bytes.Buffer
	out = &buf
	main()
	result := buf.String()

	if result != want {
		t.Errorf("expected %v, got %v", want, result)
	}
}

func Test_printfib(t *testing.T) {

	tests := []struct {
		name string
		args int
		want string
	}{
		{"Print Fibonacci sequence for 0", 0, ""},
		{"Print Fibonacci sequence for 1", 1, "1\n"},
		{"Print Fibonacci sequence for 2", 2, "1\n1\n"},
		{"Print Fibonacci sequence for 3", 3, "1\n1\n2\n"},
		{"Print Fibonacci sequence for 7", 7, "1\n1\n2\n3\n5\n8\n13\n"},
	}

	for _, tr := range tests {
		tr2 := tr

		t.Run(tr2.name, func(t *testing.T) {
			var buf bytes.Buffer
			out = &buf
			printfib(tr2.args)
			result := buf.String()

			if result != tr2.want {
				t.Errorf("expected %v, got %v", tr2.want, result)
			}

		})
	}
}

func Test_fibN(t *testing.T) {

	tests := []struct {
		name string
		args int
		want int
	}{
		{"Fibonacci of 0", 0, 0},
		{"Fibonacci of 1", 1, 1},
		{"Fibonacci of 2", 2, 1},
		{"Fibonacci of 3", 3, 2},
		{"Fibonacci of 7", 7, 13},
		{"Fibonacci of 20", 20, 6765},
	}

	for _, tr := range tests {
		tr2 := tr

		t.Run(tr2.name, func(t *testing.T) {
			result := fibN(tr2.args)
			if result != tr2.want {
				t.Errorf("expected %v, got %v", tr2.want, result)
			}
		})
	}
}
