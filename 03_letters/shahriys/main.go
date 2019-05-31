package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var outletter io.Writer = os.Stdout

func letters(s string) map[rune]int {

	alphamap := make(map[rune]int)

	for i := 0; i < len(s); i++ {
		if val, ok := alphamap[rune(s[i])]; ok {
			alphamap[rune(s[i])] = val + 1
		} else {
			alphamap[rune(s[i])] = 1
		}

	}
	//fmt.Println(alphamap)
	return alphamap
}

// As rune type can't be sorted so creating new type
type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sortLetters(m map[rune]int) []string {

	var keys = make([]rune, 0, 256)
	for key := range m {
		keys = append(keys, key)
	}
	var runes RuneSlice = keys
	sort.Sort(runes)
	var returnval = make([]string, 0, 256)

	for _, k := range runes {
		s := string(k) + ":" + strconv.FormatInt(int64(m[k]), 10)
		returnval = append(returnval, s)
	}
	//fmt.Printf("%v  %T", returnval, returnval)
	return returnval
}
func main() {
	fmt.Fprint(outletter, strings.Join(sortLetters(letters("aba")), "\n"))
}
