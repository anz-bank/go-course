package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//base fibonacci function
func fib(n int){
	//testing the sequence printer
	sequence := []int{2, 3, 4}
	printFibSequence(sequence)
}

func calculateNormalFib(x int){
	
}

func calculateNegaFib(x int){
	
}

//loop through and print the sequence.
func printFibSequence(sequence []int){
	for _, num := range sequence {
        fmt.Fprintln(out, num)
	}
}

func main() {
	fib(7)
}
