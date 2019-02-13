package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//IPAddr To export
type IPAddr [4]byte

func (ipadd IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipadd[0], ipadd[1], ipadd[2], ipadd[3])
}

func main() {
	fmt.Fprint(out, IPAddr{127, 0, 0, 1})
}
