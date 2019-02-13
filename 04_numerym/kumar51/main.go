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
		result = append(result, getResponse(val))
	}
	return result
}

func getResponse(val string) string {
	lenght := len(val)
	if lenght < 4 {
		return val
	}
	res := string(val[0]) + strconv.Itoa(lenght-2) + string(val[lenght-1])
	return res
}
