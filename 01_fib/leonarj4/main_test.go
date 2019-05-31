package main

import "testing"

func TestFib(t *testing.T){

	for i:= 0; i <= 7; i++ {
		
		switch total := fib(i)
		total {
		case 1,3,5,7,9,11,13:
			t.Logf("The test passed with current number: %d", total)
		default:
			t.Errorf("The test did not pass with current number: %d", total)
		}
	}
	
	 	



}