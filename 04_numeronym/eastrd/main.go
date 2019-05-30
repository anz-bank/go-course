package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func numeronyms(vals ...string) []string {
	/*
		Returns a slice of numeronyms for its input strings
	*/
	res := make([]string, 0)
	for _, v := range vals {
		if len(v) < 4 {
			res = append(res, v)
		} else {
			res = append(res, string(v[0])+strconv.Itoa(len(v)-2)+string(v[len(v)-1]))
		}
	}

	return res
}

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
