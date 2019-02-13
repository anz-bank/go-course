package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	var arr = make([]string, len(vals))
	for i := 0; i < len(vals); i++ {
		val := vals[i]
		if len(vals[i]) > 4 {
			val = val[1:]
			val = val[:len(val)-1]
			arr[i] = string(vals[i][0]) + strconv.Itoa(len(val)) + string(vals[i][len(vals[i])-1])
		} else {
			arr[i] = val
		}
	}
	return arr
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
