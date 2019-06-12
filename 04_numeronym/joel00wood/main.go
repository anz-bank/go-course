package main

import (
	"fmt"
)

func main() {
	fmt.Println(numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronym(s string) string {
	out := s
	sLen := len(s)
	midLen := sLen - 2
	if midLen > 1 {
		out = fmt.Sprintf("%s%v%s",
			s[:1], midLen, s[sLen-1:])
	}
	return out
}

func numeronyms(vals ...string) []string {
	out := make([]string, len(vals))
	for i, val := range vals {
		out[i] = numeronym(val)
	}
	return out
}
