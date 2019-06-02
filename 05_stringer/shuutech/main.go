package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ipAddress IPAddr) String() string {
	ip := fmt.Sprintf("%d.%d.%d.%d", ipAddress[0], ipAddress[1], ipAddress[2], ipAddress[3])
	return ip
}

func main() {
	fmt.Println(IPAddr{127, 0, 0, 1})
}
