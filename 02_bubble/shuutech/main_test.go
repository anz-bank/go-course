package main
import (
	"strconv"
	"testing"
	"fmt"
	"os"
	
)

func TestBubble(t *testing.T) {
	var unsorted = []int{3, 2, 1, 5}
	var sorted = bubble(unsorted)
	var s = ""

for i := range sorted {
	s =  s + "," + strconv.Itoa(sorted[i])
}
	var expected = ",1,2,3,5"
	if expected != s {
		t.Errorf("expected %v, got %v", expected, s)
		t.Fail()
}
}

func TestMain(m *testing.M) {
	main()
	fmt.Printf("test!")
	os.Exit(m.Run())
}
