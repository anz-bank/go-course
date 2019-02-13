package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

type IPAddr [4]byte

func main() {
	// fmt.Println(iPAddr{127, 0, 0, 1})
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for _, ip := range hosts {
		fmt.Fprint(out, ip)
	}
}
