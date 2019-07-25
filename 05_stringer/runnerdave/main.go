package main

import (
	"fmt"
	"io"
	"os"
)

// IPAddr represents an IP address
type IPAddr [4]byte

var out io.Writer = os.Stdout

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	fmt.Fprintln(out, IPAddr{127, 0, 0, 1})
}
