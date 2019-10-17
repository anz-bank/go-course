package main

import (
	"fmt"
	"io"
	"os"
)

type ipAddr [4]byte

func (i ipAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintf(out, "%v", ipAddr{127, 0, 0, 1})
}
