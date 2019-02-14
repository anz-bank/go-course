package main

import (
	"fmt"
	"io"
	"os"
)

type IPAddr [4]byte

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, IPAddr{127, 0, 0, 1})
}

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}
