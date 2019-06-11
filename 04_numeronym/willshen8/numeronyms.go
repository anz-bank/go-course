package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	returnNumeronyms := []string{}
	for _, val := range vals {
		returnNumeronyms = append(returnNumeronyms, convertStringToNumeronym(val))
	}
	return returnNumeronyms
}

func convertStringToNumeronym(input string) string {
	regEx := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	formattedInput := regEx.ReplaceAllString(input, "")

	if len(formattedInput) <= 3 {
		return formattedInput
	}

	firstChar := string(formattedInput[0])
	lastChar := string(formattedInput[len(formattedInput)-1])
	return firstChar + strconv.Itoa(len(formattedInput)-2) + lastChar
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
