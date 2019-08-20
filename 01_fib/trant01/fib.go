package main

import "fmt"

func fibo (n int) int {
    if n == 0 || n == 1{
        return n
    }
    return fibo(n-1) + fibo(n-2)
}

func main() {
    var n int;
    fmt.Println("Please input an integer greater than 2:")
    fmt.Scan(&n)
    fmt.Printf("\nThe %d-number Fibonacci sequence is: \n", n)
    fmt.Println(0)
    fmt.Println(1)
    for i := 2; i < n; i++ {
        fmt.Println(fibo(i))
    }
}