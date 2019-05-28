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
	if(n>0){
		printFib(n,"positive")
	}else if(n<0){
		printFib(n,"negative")
	}
	}


func printFib(n int,valtype string){

	var first,second,count,sum int = 1, 0,1,0
	n= int(math.Abs(float64(n)))
	for count<=n {
	sum=first+second
	if(valtype=="negative" && count%2==0){
		
		fmt.Fprintln(out,sum*-1)	
	}else{
		fmt.Fprintln(out,sum)
	}
	first=second
	second=sum
	count+=1
	}
}