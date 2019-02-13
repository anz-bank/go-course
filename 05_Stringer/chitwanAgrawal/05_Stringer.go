package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

type IPAddr struct {
	networkIDFirst byte
	networkIDSecond byte
	hostIDFirst byte
	hostIDSecond byte
}

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
   return fmt.Sprintf("%v.%v.%v.%v", ip.networkIDFirst, ip.networkIDSecond, ip.hostIDFirst, ip.hostIDSecond)
} 

func main() {
	fmt.Fprintln(out, IPAddr{127, 0, 0, 1})
}

