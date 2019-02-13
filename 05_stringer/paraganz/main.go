package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//IPAddr will satisfy stringer
type IPAddr [4]byte

func (ipaddr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3])
}

func main() {
	fmt.Fprintln(out, IPAddr{127, 0, 0, 1})
}
