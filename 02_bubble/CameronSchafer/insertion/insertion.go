package insertion

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//https://en.wikipedia.org/wiki/Insertion_sort
func insertion(s []int) []int {
	//loop through the given slice
	for i := 0; i < len(s); i++ {
		a, firstPart, secondPart := s[i], s[:i], s[i+1:]

		for j := len(firstPart) - 1; j >= 0; j-- {
			if a < firstPart[j] {
				newA, b := a, firstPart[j]
				//create the new slice and reset the parent loop
				s, i = append(append(firstPart[:j], newA), append([]int{b}, secondPart...)...), 0
				break
			}
		}
	}
	return s
}

func main() {
	fmt.Fprintln(out, insertion([]int{3, 2, 1, 5}))
}
