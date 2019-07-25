package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	log.SetFlags(log.Ldate)
	fib(7)
}

func fib(n int) {
	if n < 0 {
		log.Printf("Fib of %d is not available. Defaulting to minimum fib of 0", n)
		n = 0
	} else if n > 45 {
		log.Printf("Fib of %d is not available. Defaulting to maximum fib of 45", n)
		n = 45
	}
	n1, n2 := 1, 1
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, n1)
		n1, n2 = n2, n1+n2
	}
}
