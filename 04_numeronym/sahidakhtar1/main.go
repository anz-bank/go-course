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
	var count int
	var s1, s2 string
	for i, v := range trimedString {
		count++
		if i == 0 {
			s1 = string(v)
		}
		s2 = string(v)
	}
	if count <= 3 {
		numeronym = trimedString
	} else {
		numeronym = s1 + strconv.Itoa(count-2) + s2
	}
	return numeronym
}
