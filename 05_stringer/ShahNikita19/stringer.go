package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//IPAddr will satisfy stringer
type IPAddr [4]byte

func (ipAddress IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipAddress[0], ipAddress[1], ipAddress[2], ipAddress[3])
}

func main() {
	fmt.Fprintln(out, IPAddr{127, 0, 0, 1})
}
