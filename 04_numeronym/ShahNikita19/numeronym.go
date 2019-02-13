package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, numeronymsFind("accessibility", "Kubernetes", "abc"))
}

func numeronymsFind(vals ...string) []string {
	numeronyms := []string{}
	for _, val := range vals {
		numeronyms = append(numeronyms, generateNumeronym(val))
	}
	return numeronyms
}

func generateNumeronym(input string) string {
	inputSize := len(input)
	if inputSize <= 3 {
		return input
	}
	return fmt.Sprintf("%c%d%c", input[0], inputSize-2, input[inputSize-1])
}
