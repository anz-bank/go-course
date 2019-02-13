package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode/utf8"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	var arr = make([]string, len(vals))
	for pos, val := range vals {
		if len(val) > 3 {
			arr[pos] = string(val[0]) + strconv.Itoa(utf8.RuneCountInString(val)-2) + string(val[len(val)-1])
		} else {
			arr[pos] = val
		}
	}
	return arr
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
