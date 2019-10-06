package main

import (
	"fmt"
	"io"
	"os"
)

type ipAddr [4]byte

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, ipAddr{127, 0, 0, 1})
}

func (i ipAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", i[0], i[1], i[2], i[3])
}
