package main

import "fmt"

type IPAddr struct {
	net1, net2, subNet, host byte
}

func main() {
	fmt.Println(IPAddr{127, 0, 0, 1})
}

func (address IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", address.net1, address.net2, address.subNet, address.host)
}
