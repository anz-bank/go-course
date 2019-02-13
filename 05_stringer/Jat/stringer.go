package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

func main() {
	ip := IPAddr{127, 0, 0, 1}
	fmt.Fprintln(out, ip)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	seperator := "."
	return strconv.Itoa(int(ip[0])) + seperator +
		strconv.Itoa(int(ip[1])) + seperator +
		strconv.Itoa(int(ip[2])) + seperator + strconv.Itoa(int(ip[3]))
}
