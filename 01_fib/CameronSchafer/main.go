package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//base fibonacci function
func fib(n int){
	//testing the normal fib sequence function.
	sequence := calculateNormalFib(7)
	printFibSequence(sequence)
}

//calculates the normal fibonacci sequence 
//returns the sequence as an integer array.
func calculateNormalFib(x int) []int{
	calcd_sequence := []int{2, 3, 4}
	return calcd_sequence
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
