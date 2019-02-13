package main

import (
	"fmt"
	"strconv"
	"bytes"
)

func main() {
   fmt.Println(numeronyms("accessibility", "Kubernetes", "abc"))
}

func numeronyms(vals ...string) []string {
	resultSlice := []string{}
    for _, val := range vals {
		//fmt.Println("value ", val)
		if len(val) < 4 {
			resultSlice = append(resultSlice, val)
		} else {
			var buf bytes.Buffer
            buf.WriteString(string(val[0]))
			buf.WriteString(strconv.Itoa(int(len(val)-2)))
			buf.WriteString(string(val[len(val)-1]))
            resultSlice = append(resultSlice,  buf.String())
		}
	}
	return resultSlice
}