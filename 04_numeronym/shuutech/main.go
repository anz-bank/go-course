package main

import (
	"fmt"
	"strconv"
)

func numeronyms(vals ...string) []string {
	numeronym := []string{}
	var stringToConvert string
	for i := range vals {
		if len(vals[i]) <= 3 {
			numeronym = append(numeronym, vals[i])
		}
		if len(vals[i]) > 3 {
			stringToConvert = vals[i]
			numeronym = append(numeronym,
				string(stringToConvert[0])+
					strconv.Itoa(len(stringToConvert)-2)+
					string(stringToConvert[len(stringToConvert)-1]))
		}

	}
	return numeronym
}

func main() {
	fmt.Println(numeronyms("accessibility", "Kubernetes", "abc"))
}
