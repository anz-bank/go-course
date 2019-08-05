package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

// IPAddr is an IPv4 Address
type IPAddr [4]byte

func (a IPAddr) String() string {
	return fmt.Sprint(a[0], ".", a[1], ".", a[2], ".", a[3])
}

func main() {
	fmt.Fprintln(out, IPAddr{127, 0, 0, 1})
}
