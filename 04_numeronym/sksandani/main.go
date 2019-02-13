package main

import (
	"bytes"
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
	results := make([]string, len(vals))
	in := []string{}
	copy(in, vals)
	i := 0
	for _, val := range vals {
		results[i] = getNumeronyms(val)
		i++
	}
	return results
}

func getNumeronyms(val string) string {
	var buffer bytes.Buffer
	in := strings.Trim(val, " ")
	if len(in) <= 3 {
		buffer.WriteString(in)
	} else {
		buffer.WriteString(in[0:1])
		buffer.WriteString(strconv.Itoa(len(in) - 2))
		buffer.WriteString(in[len(in)-1:])
	}
	return buffer.String()
}
