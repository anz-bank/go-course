package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

type IPAddr [4]byte

func (ipNum IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipNum[0], ipNum[1], ipNum[2], ipNum[3])
}

func main() {
	fmt.Fprint(out, IPAddr{127, 0, 0, 1})
}
