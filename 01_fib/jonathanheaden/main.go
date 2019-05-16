package main

import (
	"fmt"
)

var fibs = map[int]int{}


func fib(n int) {
	prefix := ""
	if (n < 0){
		n = n*-1
		prefix = "-"
	}
	fibs[0] = 1
	fibs[1] = 1
	for i := 0; i < n; i++ {
		if (i > 1) {
			fibs[i] = fibs[i-1] + fibs[i-2]
		}
		fmt.Printf("%s%d\n",prefix,fibs[i])
	}
}

func main() {
	fib(-7)
}
