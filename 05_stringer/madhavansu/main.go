package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

// IPAddr export
type IPAddr [4]byte

func main() {
	fmt.Fprint(out, IPAddr{127, 0, 0, 1})
}

func (ipAddr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}
