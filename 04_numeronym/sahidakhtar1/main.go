package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	result := make([]string, len(vals))
	for i, v := range vals {
		result[i] = createNumeronym(v)
	}
	return result
}

func createNumeronym(s string) string {
	trimedString := strings.TrimSpace(s)
	var numeronym string
	//If the string contains space or \t or \n then it will return the same string
	if strings.Contains(trimedString, " ") ||
		strings.Contains(trimedString, "\t") ||
		strings.Contains(trimedString, "\n") {
		return s
	}
	if len(trimedString) <= 3 {
		numeronym = trimedString
	} else {
		numeronym += trimedString[:1]
		numeronym += strconv.Itoa(len(trimedString) - 2)
		numeronym += trimedString[len(trimedString)-1:]
	}
	return numeronym
}
