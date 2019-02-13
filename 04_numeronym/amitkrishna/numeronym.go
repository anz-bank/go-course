package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	var slice []string
	var s string
	for _, val := range vals {
		if len(val) < 4 {
			s = val
		} else {
			s = string(val[0]) + strconv.Itoa(len(val)-2) + string(val[len(val)-1])
		}
		slice = append(slice, s)
	}
	return slice
}
