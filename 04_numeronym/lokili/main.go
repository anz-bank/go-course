package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	results := make([]string, len(vals))
	for i, val := range vals {
		arr := []rune(val)
		arrLen := len(arr)
		if arrLen < 4 {
			results[i] = val
		} else {
			results[i] = fmt.Sprintf("%c%d%c", arr[0], arrLen-2, arr[arrLen-1])
		}
	}

	return results
}

func main() {
	fmt.Fprint(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
