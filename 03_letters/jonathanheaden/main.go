package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var out io.Writer = os.Stdout

func sortLetters(m map[rune]int) []string {
	var rv = make([]string, 0)
	var keys = make([]int, 0, len(m))
	var sm = make(map[int]int)
	for k := range m {
		sm[int(k)] = m[k]
		keys = append(keys, int(k))
	}

	sort.Ints(keys)
	for _, i := range keys {
		rs, _ := strconv.Unquote(strconv.QuoteRune(rune(i)))
		rv = append(rv, fmt.Sprintf("%s:%d", rs, sm[i]))
	}
	return rv
}

func letters(s string) map[rune]int {
	rm := make(map[rune]int)
	for _, rn := range s {
		rm[rn]++
	}
	return rm
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
