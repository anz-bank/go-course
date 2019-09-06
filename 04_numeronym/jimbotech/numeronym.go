package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

var out io.Writer = os.Stdout

// following rules from https://www.revolvy.com/page/Numeronym
// Program need not convert words with length 3, and should not convert shorter words.
// If an l (lowercase ell) would occur before a 1 (one), it should be made uppercase.
// If an I (uppercase eye) would occur before a 1 (one), it should be made lowercase.
// Input will be printable ASCII with spaces.
func wordMixer() func(inputStr string) string {

	knownConversions := map[string]string{
		"year 2038":        "Y2K38",
		"year 2000":        "Y2K",
		"i love you":       "143",
		"eyjafjallajökull": "E15",
	}

	return func(inputStr string) string {

		inputStr = strings.TrimSpace(inputStr)
		runes := []rune(inputStr)

		if len(runes) <= 3 {
			return inputStr
		}

		lowerRunes := []rune(inputStr)
		for i, v := range lowerRunes {
			lowerRunes[i] = unicode.ToLower(v)
		}
		if value, found := knownConversions[string(lowerRunes)]; found {
			return value
		}

		firstChar := runes[0]
		lastChar := runes[len(runes)-1]
		len := len(runes) - 2

		if len > 9 && len < 20 {
			if firstChar == 'I' {
				firstChar = 'i'
			} else if firstChar == 'l' {
				firstChar = 'L'
			}
		}
		return fmt.Sprintf("%c%d%c", firstChar, len, lastChar)
	}
}

func numeronyms(vals ...string) []string {

	f := wordMixer()
	results := []string{}

	for _, val := range vals {
		results = append(results, f(val))
	}
	return results
}

func main() {
	fmt.Fprintln(out, numeronyms("accessibility", "Kubernetes", "abc"))
}
