package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

var mainout io.Writer = os.Stdout

func main() {
	fmt.Fprintln(mainout, numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	//Regular expression can be changed or removed altogether if additional chars need to be included
	regEx := regexp.MustCompile("[^a-zA-Z0-9]+")
	valsOut := vals
	for i := range vals {
		alphaNumVal := regEx.ReplaceAllString(vals[i], "") //Strips all non-alpha-numeric characters
		length := len(alphaNumVal)
		if length > 3 {
			valsOut[i] = fmt.Sprintf("%v%d%v", string(alphaNumVal[0]), length-2, string(alphaNumVal[length-1]))
		}
	}
	return valsOut
}
