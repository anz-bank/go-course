package main

import "fmt"

func main(){

	fib(-7)  //Provide the input, programmed for both Positive and Negative Series
}

func fib(n int){

	var a int  = 0 //first term
	var b int  = 1 //second term
	var c int      //next term

//For Positive Series only

	for i := 0;i < n ;i++ {
	
		c = a+b
		fmt.Println(c)  //Printing next term
		if (i>0){  //variable reassignment for next iteration
		a=b
		b=c
		}
	}
//For Negative Series only
	for i := 0;i > n ;i-- {
		
		if (i<0){
		c = a-b
		fmt.Println(c)  //Printing next term
		a=b
		b=c
		} else {        //To cover exceptions in negative series
			c=a+b
			fmt.Println(c)
		}
	}
}