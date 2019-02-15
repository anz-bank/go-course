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
	numeronyms := make([]string, len(vals))
	for i, val := range vals {
		numeronyms[i] = generateNumeronym(val)
	}
	return numeronyms
}

func generateNumeronym(input string) string {
	runeRep := []rune(input)
	inputLen := len(runeRep)
	//minimum 4 chars are required
	if inputLen < 4 {
		return input
	}
	return fmt.Sprintf("%c%d%c", runeRep[0], inputLen-2, runeRep[inputLen-1])
}
