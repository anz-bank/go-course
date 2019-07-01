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

	numeronyms := make([]string, 0)
	for _, val := range vals {
		if val == "" {
			numeronyms = append(numeronyms, "")
		} else {
			strArr := []rune(val)
			staVal := string(strArr[0])
			midVal := ""
			endVal := ""
			switch {
			case len(strArr) == 2:
				endVal = string(strArr[len(strArr)-1])
			case len(strArr) == 3:
				midVal = string(strArr[1])
				endVal = string(strArr[len(strArr)-1])
			case len(strArr) > 3:
				midVal = fmt.Sprint(len(strArr) - 2)
				endVal = string(strArr[len(strArr)-1])
			}
			numeronyms = append(numeronyms, fmt.Sprintf("%s%s%s", staVal, midVal, endVal))
		}
	}
	return numeronyms
}
