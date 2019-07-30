package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

// IPAddr holds an IP Address by storing 4 bytes.
type IPAddr [4]byte

// String implements `fmt.Stringer` to print the address as a dotted quad.
func (part IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", part[0], part[1], part[2], part[3])
}

func main() {
	fmt.Fprint(out, IPAddr{127, 0, 0, 1})
}
