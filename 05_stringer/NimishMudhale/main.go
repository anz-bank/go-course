package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

type IPAddr [4]byte

func (address IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", address[0], address[1], address[2], address[3])
}
func main() {
	fmt.Fprint(out, IPAddr{127, 0, 0, 1})
}
