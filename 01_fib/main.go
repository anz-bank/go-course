package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

func main() {
	//fmt.Fprintln(out, " Fibnacci Series")
	fib(7)
}

func fib(limit int) {
	firstNo := 0
	secondNo := firstNo + 1
	var finalString string
	for i := 0; i < limit; i++ {
		//fmt.Println(firstNo)
		finalString += strconv.Itoa(firstNo)
		temp := secondNo
		secondNo += firstNo
		firstNo = temp
	}
	//fmt.Println(finalString)
	fmt.Fprintln(out, finalString)
}
