package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

func nemeronyms(vals ...string) []string {
	cp := make([]string, len(vals))
	for i := 0; i < len(vals); i++ {
		temp := vals[i]
		if len(temp) > 3 {
			temp1 := ""
			num := 0
			for j, r := range temp {
				if j == 0 {
					temp1 = string(r)
				} else {
					num++
				}
				if j == len(temp)-1 {
					num = num - 1
					temp1 = temp1 + strconv.Itoa(num) + string(r)
				}
			}
			cp[i] = temp1
		} else {
			cp[i] = temp
		}

	}
	return cp
}

func main() {
	fmt.Fprint(out, nemeronyms("accessibility", "Kubernetes", "abc", "abcd"))
}
