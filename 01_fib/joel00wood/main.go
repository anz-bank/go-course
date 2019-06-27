package main

import (
	"fmt"
	"math"
)

func fibCalc(n float64) float64 {
	switch n {
	case 0, 1:
		return n
	default:
		return fibCalc(n-1) + fibCalc(n-2)
	}
}

func fibN(n float64) float64 {
	if n < 0 {
		var i float64
		n := math.Abs(n)
		i = math.Pow(-1, n+1)
		return fibCalc(n) * i
	}
	return fibCalc(n)
}

func fib(n int) {
	for i := 0.0; i <= math.Abs(float64(n)); i++ {
		if n < 0 && i > 0 {
			fmt.Println(fibN(i * -1))
		} else {
			fmt.Println(fibN(i))
		}
	}
}

func main() {
	fib(7)
}
