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

func (i IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", i[0], i[1], i[2], i[3])
}
