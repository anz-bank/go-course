package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	numeronyms := []string{}
	for _, val := range vals {
		numeronyms = append(numeronyms, generateNumeronym(val))
	}
	return numeronyms
}

func generateNumeronym(input string) string {
	inputLen := len(input)
	//minimum 4 chars are requried
	if inputLen < 4 {
		return input
	}
	return fmt.Sprintf("%c%d%c", input[0], inputLen-2, input[inputLen-1])
}

func main() {
	fmt.Fprint(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
