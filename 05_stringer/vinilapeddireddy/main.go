package main

import (
	"fmt"
	"io"
	"os"
)

type IPAddr [4]byte

var out io.Writer = os.Stdout

// TODO: Add a "String() string" method to IPAddr.
func (addr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", addr[0], addr[1], addr[2], addr[3])
}
func main() {
	fmt.Fprint(out, IPAddr{127, 0, 0, 1})
}
