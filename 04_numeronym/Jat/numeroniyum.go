package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	var result = make([]string, 0)
	for _, val := range vals {
		result = append(result, getNumeronym(val))
	}
	return result
}

func getNumeronym(val string) string {

	lenght := len(val)
	if lenght <= 3 {
		return val
	}
	numeronym := string(val[0]) + strconv.Itoa(lenght-2) + string(val[lenght-1])
	return numeronym
}
