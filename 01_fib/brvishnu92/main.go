package main

import (
	"fmt"
	"io"
	"os"
	"math"
)

var out io.Writer = os.Stdout

func main() {
	fib(7)
}


func fib(n int){

	var first,second,count,sum int = 1, 0,1,0
	n= int(math.Abs(float64(n)))
	for count<=n {
	sum=first+second
	fmt.Fprintln(out,sum)
	first=second
	second=sum
	count+=1
	}
}
