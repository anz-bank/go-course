package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	result := make([]string, 0)
	for _, val := range vals {
		numeronym := ""
		length := len([]rune(val))
		if length <= 3 {
			numeronym = val
		} else {
			// convert rune slice to string using typecast string()
			// why use typecast []rune? because special character can take more than 1 byte space
			numeronym = string([]rune(val)[0]) + strconv.Itoa(length-2) + string([]rune(val)[length-1])
		}
		result = append(result, numeronym)
	}

	return result
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
