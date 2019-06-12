package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func Fib(number int) {
	sign, firstNum, nextNum := 1, int64(0), int64(1)
	switch number {
	case 0:
		fmt.Fprintln(out, number)
	default:
		if number > 92 {
			number = 92
			fmt.Fprintln(out, "92 is max positive fibonacci series.")
		} else if number < -92 {
			number = -92
			fmt.Fprintln(out, "-92 is max negative fibonacci series.")
		}
		if number < 0 {
			sign = -1
		}
		number *= sign
		for i := 0; i < number; i++ {
			if sign > 0 {
				firstNum, nextNum = nextNum, firstNum+nextNum
			} else {
				firstNum, nextNum = nextNum, firstNum-nextNum
			}
			fmt.Fprintln(out, firstNum)
		}

		// if number < 0 { // Negative Fibonacci series
		// 	for i := -1; i >= number; i-- {
		// 		firstNum, nextNum = nextNum, firstNum-nextNum
		// 		fmt.Fprintln(out, firstNum)
		// 		if i <= -92 { //-92 is the last good value for 64 bit int
		// 			fmt.Fprintln(out, "max fibo series reached")
		// 			break
		// 		}
		// 	}
		// } else { // positive Fibonacci series
		// 	for i := 1; i <= number; i++ {
		// 		firstNum, nextNum = nextNum, firstNum+nextNum
		// 		fmt.Fprintln(out, firstNum)
		// 		if i >= 92 { //92 is the last good value for 64 bit int
		// 			fmt.Fprintln(out, "max fibo series reached")
		// 			break
		// 		}
		// 	}
		// }

	}
}

func main() {
	Fib(7)

}
