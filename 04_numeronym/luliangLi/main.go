package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

func transfer(in string) string {
	if len(in) > 3 {
		head := in[0:1]
		tail := in[len(in)-1:]
		body := strconv.Itoa(len(in) - 2)

		return head + body + tail
	}
	return in
}

func numeronyms(vals ...string) []string {
	rlt := make([]string, len(vals))

	i := 0
	for _, val := range vals {
		rlt[i] = transfer(val)
		i++
	}

	return rlt
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
