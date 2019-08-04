package main

import (
	"fmt"
	"io"
	"os"
)

type IPAddr struct {
	ip1, ip2, ip3, ip4 byte
}

var out io.Writer = os.Stdout

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip.ip1, ip.ip2, ip.ip3, ip.ip4)
}

func main() {
	fmt.Fprintln(out, IPAddr{127, 0, 0, 1})
}
