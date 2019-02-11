
package main

import (
"fmt"
"io"
"os"
)

var out io.Writer = os.Stdout

var (
    num1 = 0
    num2 = 1
    sum = 0
)

func main() {
    Fib(7)
}

func Fib(n int) int {
    fmt.Fprintln(out, num2)
    for i := 1; i < n; i++ {
        sum = num1 + num2
        num1 = num2
        num2 = sum
        fmt.Fprintln(out, sum)
    }
    return sum
}
