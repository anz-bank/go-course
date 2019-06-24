package main

import (
	"fmt"
	"io"
	"os"
)
var out io.Writer = os.Stdout

func main() {
  store1 := NewSyncStore()
  store1.CreatePuppy(Puppy{
    ID: "1",
    Colour: "Red",
  })

  puppy, _ := store1.ReadPuppy("1")
  fmt.Fprint(out, puppy.Colour)
}
