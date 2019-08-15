package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

type IPAddr [4]byte

func (ip4 IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip4[0], ip4[1], ip4[2], ip4[3])
}

func main() {
	fmt.Fprintln(out, IPAddr{127, 0, 0, 1})
}
