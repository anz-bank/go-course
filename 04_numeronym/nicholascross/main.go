package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	nyms := make([]string, 0, len(vals))
	for _, str := range vals {
		nyms = append(nyms, numeronym(str))
	}
	return nyms
}

func numeronym(str string) string {
	count := len(str)
	if count > 3 {
		return fmt.Sprintf("%c%d%c", str[0], count-2, str[count-1])
	}
	return str
}
