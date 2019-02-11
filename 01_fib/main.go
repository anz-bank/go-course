package main

import (
	"fmt"
)

func main() {
	fmt.Println(" Fibnacci Series")
	fib(7)
}

func fib(limit int) {
	firstNo := 0
	secondNo := firstNo + 1
	for i := 0; i < limit; i++ {
		fmt.Println(firstNo)
		temp := secondNo
		secondNo = secondNo + firstNo
		firstNo = temp
	}
}
