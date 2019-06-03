package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func Fib(number int) {
	firstNum, nextNum := int64(0), int64(1)

	if number == 0 {
		fmt.Fprintln(out, number)
	} else if number < 0 { // Negative Fibonacci series
		for i := -1; i >= number; i-- {
			firstNum, nextNum = nextNum, firstNum-nextNum
			fmt.Fprintln(out, firstNum)
			if i <= -92 { //-92 is the last good value for 64 bit int
				fmt.Fprintln(out, "max fibo series reached")
				break
			}
		}
	} else { // positive Fibonacci series
		for i := 1; i <= number; i++ {
			firstNum, nextNum = nextNum, firstNum+nextNum
			fmt.Fprintln(out, firstNum)
			if i >= 92 { //92 is the last good value for 64 bit int
				fmt.Fprintln(out, "max fibo series reached")
				break
			}
		}
	}
}

func main() {
	Fib(7)

}
