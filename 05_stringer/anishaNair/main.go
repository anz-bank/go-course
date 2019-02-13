package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

type IPAddr [4]byte

func (r IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", r[0], r[1], r[2], r[3])
}

func main() {
	fmt.Fprintln(out, IPAddr{192, 0, 0, 1})
}
