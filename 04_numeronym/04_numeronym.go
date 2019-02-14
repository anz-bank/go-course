package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode/utf8"
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
			numeronymsArray = append(numeronymsArray, string(val[0])+strconv.Itoa(utf8.RuneCountInString(val)-2)+string(val[len(val)-1]))
		}
	}
	return numeronymsArray
}
