package main

import (
	"bytes"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buffib bytes.Buffer
	outfib = &buffib
	main()
	expected :=
		`1
1
2
3
5
8
13
1
-1
2
-3
5
-8
13
`
	actual := buffib.String()

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}
func TestFibPos(t *testing.T) {
	var buffib bytes.Buffer
	outfib = &buffib

	fib(7)

	expected := `1
1
2
3
5
8
13
`
	actual := buffib.String()

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}
func TestFibNeg(t *testing.T) {
	var buffib bytes.Buffer
	outfib = &buffib

	fib(-7)

	expected := `1
-1
2
-3
5
-8
13
`
	actual := buffib.String()

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}
