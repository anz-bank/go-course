package main

import (
	"fmt"
	"io"
	"os"
)

// IPAddr is below, violets are blue and this is here to shut up golint :-)
type IPAddr [4]byte

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, IPAddr{127, 0, 0, 1})
}

func (i IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", i[0], i[1], i[2], i[3])
}
