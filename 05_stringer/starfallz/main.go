package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, IPAddr{127, 0, 0, 1})
}

type IPAddr struct {
	add1, add2, add3, add4 uint8
}

func (ipAddr IPAddr) String() string {
	return strings.Join([]string{
		strconv.Itoa(int(ipAddr.add1)),
		strconv.Itoa(int(ipAddr.add2)),
		strconv.Itoa(int(ipAddr.add3)),
		strconv.Itoa(int(ipAddr.add4))}, ".")
}
