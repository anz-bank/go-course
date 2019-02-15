package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//IPAddr will satisfy stringer
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	fmt.Fprint(out, IPAddr{12, 23, 23, 24})
}
