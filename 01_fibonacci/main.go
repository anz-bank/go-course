package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	for num := 1; num < 10; num++ {
		fmt.Fprintln(out, fibinocci(num))
	}
}
func fibinocci(x int) int {
	if x < 2 {
		return x
	}
	return fibinocci(x-1) + fibinocci(x-2)

}
