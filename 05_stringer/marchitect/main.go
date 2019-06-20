package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad

func main() {
	fmt.Fprintln(out, IPAddr{127, 0, 0, 1})
}
