package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	var numeronymsArray = []string{}
	for _, val := range vals {
		var lengthVal = len(val)
		if lengthVal < 4 {
			numeronymsArray = append(numeronymsArray, val)
		} else {
			numeronymsArray = append(numeronymsArray, fmt.Sprintf("%c%d%c", val[0], lengthVal-2, val[lengthVal-1]))
		}
	}
	return numeronymsArray
}
