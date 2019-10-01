package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	var stringList []string
	for _, s := range vals {
		if len(s) <= 3 {
			stringList = append(stringList, s)
		} else {
			stringList = append(stringList, string(s[0])+strconv.Itoa(len(s)-2)+string(s[len(s)-1]))
		}
	}
	return stringList
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
