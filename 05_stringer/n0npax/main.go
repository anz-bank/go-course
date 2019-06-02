package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// IPAddr is a type to encapsulate IPv4 (4 bytes)
type IPAddr [4]byte

// Stringer for IPAddr
func (addr IPAddr) String() string {
	values := make([]string, len(addr))
	for i, v := range addr {
		values[i] = fmt.Sprint(v)
	}
	return strings.Join(values, ".")
}

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, IPAddr{127, 0, 0, 1})
}
