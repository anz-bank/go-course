package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

var out io.Writer = os.Stdout

//base fibonacci function
func fib(n int) {
	var sequence []int

	if n >= 0 {
		sequence = calculateNormalFib(n) //normal fib sequence.
	} else if n < 0 {
		sequence = calculateNegaFib(int(math.Abs(float64(n)))) //nega fib sequence
	}
	printFibSequence(sequence)
}

//calculates the normal fibonacci sequence
//returns the sequence as an integer array.
func calculateNormalFib(x int) []int {
	var calcdSequence []int
	//starting values.
	n1 := 0
	n2 := 1

	//loop until end of sequence.
	for count := 0; count < x; count++ {
		calcdSequence = append(calcdSequence, n2) //store n2 into array
		//calc next value in the sequence + assign n1 to the old value of n2
		n1, n2 = n2, calcNextInSequence(n1, n2)
	}

	return calcdSequence
}

func calculateNegaFib(x int) []int {
	var calcdSequence []int
	//starting values.
	n1 := 0
	n2 := 1
	//the next negative num in the negafib sequence
	var nextNum int
	//loop until end of sequence.
	for count := 0; count < x; count++ {
		//calc next value in the sequence + assign n1 to the old value of n2
		n1, n2 = n2, calcNextInSequence(n1, n2)
		nextNum = int(float64(n1) * math.Pow(-1, float64(count))) //calc the next negafib number
		calcdSequence = append(calcdSequence, nextNum)            //store n2 into array
	}

	return calcdSequence
}

//function calculates the next number in the fibonacci sequence.
func calcNextInSequence(n1 int, n2 int) int {
	nextInSequence := n1 + n2
	return nextInSequence
}

//loop through and print the sequence.
func printFibSequence(sequence []int) {
	for _, num := range sequence {
		fmt.Fprintln(out, num)
	}
}

func main() {
	fib(7)
	//fib(-7)		//uncomment this line to show the negafib sequence.
}
