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
		result = append(result, prepareResponse(val))
	}
	return result
}

func prepareResponse(value string) string {
	lenght := len(value)
	if lenght <= 3 {
		return value
	}
	res := string(value[0]) + strconv.Itoa(lenght-2) + string(value[lenght-1])
	return res
}
