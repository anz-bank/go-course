package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {

	store := new(MapStore)
	fmt.Fprintln(out, store)

}
