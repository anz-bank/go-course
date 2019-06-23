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
  fmt.Fprintln(out, puppy)

  puppy.Colour = "Blue"
  store1.UpdatePuppy(puppy)
  puppy, _ = store1.ReadPuppy("1")
  fmt.Fprintln(out, puppy)
}
