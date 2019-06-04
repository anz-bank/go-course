package main

import (
	"fmt"
	"io"
	"os"
)

var outstringer io.Writer = os.Stdout

type IPAddr [4]byte

// String satisfies the fmt.Stringer interface for the IPAddr type
func (u IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", u[0], u[1], u[2], u[3])
}

func main() {
	fmt.Fprint(outstringer, IPAddr{127, 0, 0, 1})
}
