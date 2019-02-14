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
	if strings.ContainsAny(trimedString, " \t\n") {
		return s
	}
	runes := []rune(trimedString)
	count := len(runes)
	if count <= 3 {
		numeronym = trimedString
	} else {
		numeronym = string(runes[0]) + strconv.Itoa(count-2) + string(runes[count-1])
	}
	return numeronym
}
