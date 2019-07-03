package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var out io.Writer = os.Stdout

func letters(s string) map[rune]int {
	var ltrs = make(map[rune]int)
	for _, v := range s {
		ltrs[v]++
	}
	return ltrs
}

func sortLetters(m map[rune]int) []string {
	type ltrCount struct {
		r     rune
		count int
	}

	// Convert to slice
	var ltrCounts = []ltrCount{}
	for r, count := range m {
		ltrCounts = append(ltrCounts, ltrCount{r, count})
	}

	lessFn := func(i int, j int) bool {
		// Secondary sort by rune to make result deterministic
		if ltrCounts[i].count == ltrCounts[j].count {
			return ltrCounts[i].r < ltrCounts[j].r
		}
		return ltrCounts[i].count > ltrCounts[j].count
	}

	sort.Slice(ltrCounts, lessFn)

	// Format
	var ltrStrings = []string{}
	for _, v := range ltrCounts {
		str := fmt.Sprintf("%v:%v", string(v.r), v.count)
		ltrStrings = append(ltrStrings, str)
	}

	return ltrStrings
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
