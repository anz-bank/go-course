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
	numeronymStrings := make([]string, 0)
	for _, val := range vals {
		if val == "" {
			numeronymStrings = append(numeronymStrings, "")
		} else {
			runes := []rune(val)
			firstChar := string(runes[0])
			middleChar := ""
			lastChar := ""
			if len(runes) > 1 {
				lastChar = string(runes[len(runes)-1])
				if len(runes) == 3 {
					middleChar = string(runes[1])
				} else if len(runes) > 3 {
					middleChar = fmt.Sprint(len(runes) - 2)
				}
			}
			numeronymStrings = append(numeronymStrings, fmt.Sprintf("%s%s%s", firstChar, middleChar, lastChar))
		}
	}
	return numeronymStrings
}
