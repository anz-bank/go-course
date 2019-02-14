package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func numeronym(vals ...string) []string {
	numeronymArray := make([]string, len(vals))
	for i, val := range vals {
		target := []rune(strings.Trim(val, "\n\t "))
		length := len(target)
		if length <= 3 {
			numeronymArray[i] = string(target)
		} else {
			numeronymArray[i] = fmt.Sprintf("%c%d%c", target[0], length-2, target[length-1])
		}
	}
	return numeronymArray
}

func main() {
	fmt.Fprint(out, numeronym("accessibility", "Kubernetes", "abc"))
}
