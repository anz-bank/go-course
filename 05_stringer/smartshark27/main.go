package main

import (
	"fmt"
	"io"
	"os"
)

type IPAddr [4]byte

var out io.Writer = os.Stdout

func (a IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", a[0], a[1], a[2], a[3])
}

func main() {
	fmt.Fprintln(out, IPAddr{127, 0, 0, 1})
}
