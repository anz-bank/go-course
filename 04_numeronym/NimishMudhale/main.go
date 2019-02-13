package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func numeronym(vals ...string) []string {
	numeronymArray := []string{}
	for _, val := range vals {
		target := strings.TrimRight(val, "\n")
		len := len(target)
		if len <= 3 {
			numeronymArray = append(numeronymArray, target)
			continue
		}
		target = fmt.Sprintf("%c%d%c", target[0], len-2, target[len-1])
		numeronymArray = append(numeronymArray, target)
	}
	return numeronymArray
}
func main() {
	fmt.Fprint(out, numeronym("accessibility", "Kubernetes", "abc"))
}
