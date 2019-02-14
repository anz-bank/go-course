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
		target := strings.Trim(val, "\n\t ")
		len := len(target)
		if len <= 3 {
			numeronymArray[i] = target
			continue
		}
		target = fmt.Sprintf("%c%d%c", target[0], len-2, target[len-1])
		numeronymArray[i] = target
	}
	return numeronymArray
}

func main() {
	fmt.Fprint(out, numeronym("accessibility", "Kubernetes", "abc"))
}
