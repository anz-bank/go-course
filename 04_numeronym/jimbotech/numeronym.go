package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

// Program need not convert words with length 3, and should not convert shorter words.
// If an l (lowercase ell) would occur before a 1 (one), it should be made uppercase.
// If an I (uppercase eye) would occur before a 1 (one), it should be made lowercase.
func wordMixer(rawInput string) string {
	knownConversions := map[string]string{
		"year 2038":        "Y2K38",
		"year 2000":        "Y2K",
		"i love you":       "143",
		"eyjafjallajökull": "E15",
	}
	rawInput = strings.TrimSpace(rawInput)
	rw := []rune(rawInput)
	if len(rw) <= 3 {
		return rawInput
	}
	if known, found := knownConversions[strings.ToLower(rawInput)]; found {
		return known
	}
	fc := rw[0]
	lc := rw[len(rw)-1]
	l := len(rw) - 2
	if l >= 10 && l <= 19 {
		if fc == 'I' {
			fc = 'i'
		} else if fc == 'l' {
			fc = 'L'
		}
	}
	return fmt.Sprintf("%c%d%c", fc, l, lc)
}
func numeronyms(vals ...string) []string {
	results := []string{}
	for _, val := range vals {
		results = append(results, wordMixer(val))
	}
	return results
}
func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
