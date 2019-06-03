package main

import (
	"fmt"
	"strconv"
)

func numeronyms(vals ...string) []string {
	var stringList []string
	for _, s := range(vals) {
		if (len(s) <= 3) {
			stringList = append(stringList, s)
		} else {
			stringList = append(stringList, string(s[0]) + strconv.Itoa(len(s) - 2) + string(s[len(s) - 1]))
		}
	}
	return stringList
}

func main() {
	fmt.Println(numeronyms("accessibility", "Kubernetes", "abc"))
}