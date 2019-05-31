package main

import ("fmt")

func fib(n int) int {
	if (n == 0) {
		return n + 1
	}
	x,y := n, n-1 
	return x + y
}

func main() {
	for i:= 0; i <= 7; i++ {
		fmt.Printf("%d \n",fib(i))
	}
}