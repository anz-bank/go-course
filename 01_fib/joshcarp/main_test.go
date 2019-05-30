package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFib(t *testing.T) {
	tests := map[int]string{7: "1\n1\n2\n3\n5\n8\n13\n",
		-7: "1\n-1\n2\n-3\n5\n-8\n13\n"}

	for input, output := range tests {
		re := captureStdout(fib, input)
		fmt.Printf(re)
		require.Equal(t, re, output)
	}
	main()
	fmt.Println("exit")
}

func captureStdout(f func(int), n int) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f(n)
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
