package main

import (
	"fmt"
	"io"
	"os"
)

// IPAddr type accepts up to 4 bytes and prints string in IP Address format
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, IPAddr{127, 0, 0, 1})
}
