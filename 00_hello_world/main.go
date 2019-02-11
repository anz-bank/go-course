package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fib(7)
}

func fib(n int){
	f1 := 0
	f2 := 1
	fmt.Printf("%d\n",f2)
	for counter := 1; counter < n; counter++ {
		next := f1 + f2
		fmt.Printf("%d\n",next)
		f1 = f2
		f2 = next
	}
}
