package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

// calculateFibSequence creates a slice of Fibonacci numbers
// Starts from 0
// Negative integer inputs will create a Negafibonacci sequence
// Returns nil for input of 0
// Returns an error if the input would generate numbers outside of an int lower/upper values
func calculateFibSequence(n int) ([]int, error) {
	// Early bailout returning zero value of a slice
	if n == 0 {
		return nil, nil
	}
	// Early bailout and error for return values outside of int limits
	if n > 92 || n < -92 {
		return nil, fmt.Errorf("fib: %d is outside of limits of -92..92", n)
	}
	// Handle negative inputs
	// Multiplied against the initial Fibonacci value (1) to spawn either a positive or negative sequence
	sign := 1
	if n < 0 {
		n *= -1
		sign = -1
	}

	result := make([]int, n)
	// Seed initial value for the sequence
	result[0] = 0
	if n > 1 {
		// Seed second value for the sequence; accommodating for Negafibonacci
		result[1] = 1 * sign
		// Calculate remaining sequence
		for i := 2; i < n; i++ {
			result[i] = result[i-2] + result[i-1]*sign
		}
	}
	return result, nil
}

// fib prints either a loop of Fibonacci numbers or an error
func fib(n int) {
	seq, err := calculateFibSequence(n)

	if err != nil {
		fmt.Fprint(out, err)
		return
	}

	for _, num := range seq {
		fmt.Fprintln(out, num)
	}
}

func main() {
	fib(7)
}
