package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	var arr = make([]string, len(vals))
	for pos, val := range vals {
		if len(val) > 3 {
			trimmedVal := []rune(strings.Trim(val, "\n\t "))
			len := len(trimmedVal)
			arr[pos] = string(trimmedVal[0]) + strconv.Itoa(len-2) + string(trimmedVal[len-1])
		} else {
			arr[pos] = val
		}
	}
	return arr
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
