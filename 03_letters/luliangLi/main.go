package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var out io.Writer = os.Stdout

type Pair struct {
	key   rune
	value int
}

func (p Pair) String() string {
	return fmt.Sprintf("%s:%d", string(p.key), p.value)
}

func letters(s string) map[rune]int {
	rlt := make(map[rune]int)

	for _, char := range s {
		val := rlt[char]
		rlt[char] = val + 1
	}

	return rlt
}

func sortLetters(m map[rune]int) []string {
	pairs := make([]Pair, len(m))
	rlt := make([]string, len(m))

	i := 0
	for k, v := range m {
		pairs[i] = Pair{k, v}
		i++
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value > pairs[j].value
	})

	j := 0
	for _, pair := range pairs {
		rlt[j] = pair.String()
		j++
	}

	return rlt
}

func main() {
	fmt.Fprintln(out, strings.Join(sortLetters(letters("aba")), "\n"))
}
