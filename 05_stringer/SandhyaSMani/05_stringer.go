package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

type IPAddress [4]byte

func (ipaddr IPAddress) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3])
}
func main() {
	fmt.Fprintln(out, IPAddress{127, 0, 0, 1})
}
