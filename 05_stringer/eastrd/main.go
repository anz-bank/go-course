package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	res := ""
	for i, v := range ip {
		res += strconv.Itoa(int(v))
		if i < len(ip)-1 {
			res += "."
		}
	}
	return res
}

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, IPAddr{127, 0, 0, 1})
}
