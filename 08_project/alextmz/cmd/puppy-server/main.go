package main

import (
	"fmt"
	"io"
	"os"

	"github.com/anz-bank/go-course/08_project/alextmz/pkg/puppy"
	"github.com/anz-bank/go-course/08_project/alextmz/pkg/puppy/store"
)

var out io.Writer = os.Stdout

func main() {
	s1 := store.NewMapStore()
	s1p1 := puppy.Puppy{Breed: "Dogo", Colour: "White", Value: 500}
	fmt.Fprintf(out, "%-27s : ", "Creating puppy on Mapstore")
	_ = s1.CreatePuppy(&s1p1)

	fmt.Fprintf(out, "%s : %#v\n", "Created puppy", s1p1)
	fmt.Fprintf(out, "%-27s : ", "Reading puppy back")
	s1p2, _ := s1.ReadPuppy(s1p1.ID)

	fmt.Fprintf(out, "%#v\n", s1p2)

	s2 := store.NewMapStore()
	s2p1 := puppy.Puppy{Breed: "Fila", Colour: "Golden", Value: 900}
	fmt.Fprintf(out, "%-27s : ", "Creating puppy on SyncStore")
	_ = s2.CreatePuppy(&s2p1)

	fmt.Fprintf(out, "%s : %#v \n", "Created puppy", s2p1)
	fmt.Fprintf(out, "%-27s : ", "Reading puppy back")
	s2p2, _ := s2.ReadPuppy(s2p1.ID)

	fmt.Fprintf(out, "%#v\n", s2p2)
}
