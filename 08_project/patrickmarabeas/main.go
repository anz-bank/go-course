package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, "hello")
}
