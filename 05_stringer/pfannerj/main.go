package main

import (
	"fmt"
	"io"
	"os"
)

var mainout io.Writer = os.Stdout

//IPAddr is an Internet Protocol version 4 (IPv4) network address
type IPAddr [4]byte

func (i IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", i[0], i[1], i[2], i[3])
}

func main() {

	fmt.Fprintln(mainout, IPAddr{127, 0, 0, 1})

}
