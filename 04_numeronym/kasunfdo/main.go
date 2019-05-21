package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

var out io.Writer = os.Stdout

func numeronyms(s ...string) []string {
	out := make([]string, len(s))
	for i, str := range s {
		out[i] += getNumeronym(str)
	}
	return out
}

func getNumeronym(str string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	processedString := reg.ReplaceAllString(str, "")
	length := len(processedString)

	if length <= 3 {
		return processedString
	}

	return fmt.Sprintf("%c%d%c", processedString[0], length-2, processedString[length-1])
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
