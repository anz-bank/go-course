package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

type IPAddr [4]byte

func (ipAdd IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipAdd[0], ipAdd[1], ipAdd[2], ipAdd[3])
}

func main() {
	fmt.Fprint(out, IPAddr{127, 0, 0, 1})
}
