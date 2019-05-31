package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var outnum io.Writer = os.Stdout

func numeronyms(vals ...string) []string {
	var returnval = make([]string, 0, len(vals))
	for _, val := range vals {
		switch len(val) {
		case 0, 1, 2, 3:
			returnval = append(returnval, val)
		default:
			l := len(val) - 2
			s := string(val[0]) + strconv.FormatInt(int64(l), 10) + string(val[l+1])
			returnval = append(returnval, s)
		}

	}

	return returnval
}

func main() {
	fmt.Fprint(outnum, numeronyms("accessibility", "Kubernetes", "abc"))
}
