package main
import (
	"testing"
	"fmt"
	"os"
)

func TestNumeronym(t *testing.T){
	expectedSlice := []string{"a11y", "abc"}
	actualSlice := numeronyms("accessibility","abc")

	var expected string
	var actual string

	for i := 0; i < 2; i++{
		if (expectedSlice[i] != actualSlice[i]){
				t.Errorf("expected %v, got %v", actual, expected)
				t.Fail()
		}

	}
}
	func TestMain(m *testing.M) {
		main()
		fmt.Printf("test!")
		os.Exit(m.Run())
	}

