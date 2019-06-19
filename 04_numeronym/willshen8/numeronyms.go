package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	returnNumeronyms := make([]string, len(vals))
	for i, val := range vals {
		returnNumeronyms[i] = convertStringToNumeronym(val)
	}
	return returnNumeronyms
}

func convertStringToNumeronym(input string) string {
	regEx := regexp.MustCompile(`[^a-zA-Z]+`)
	formattedInput := regEx.ReplaceAllString(input, "")

	if len(formattedInput) <= 3 {
		return input
	}

	return fmt.Sprintf("%c%d%c", formattedInput[0], len(formattedInput)-2, formattedInput[len(formattedInput)-1])
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
