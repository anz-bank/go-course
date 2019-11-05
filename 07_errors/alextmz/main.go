package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	s1 := NewMapStore()
	s1p1 := Puppy{Breed: "Dogo", Colour: "White", Value: 500}
	fmt.Printf("Creating puppy on Mapstore : ")
	fmt.Printf("%#v (ID should start at 0 and be populated after creation)\n", s1p1)
	_ = s1.CreatePuppy(&s1p1)

	fmt.Fprintf(out, "Created puppy              : %#v\n", s1p1)
	fmt.Printf("Reading puppy back         : ")
	s1p2, _ := s1.ReadPuppy(s1p1.ID)

	fmt.Printf("%#v\n", s1p2)

	s2 := NewMapStore()
	s2p1 := Puppy{Breed: "Fila", Colour: "Golden", Value: 900}
	fmt.Printf("Creating puppy on SyncStore: ")
	fmt.Printf("%#v (ID should start at 0 and be populated after creation)\n", s2p1)
	_ = s2.CreatePuppy(&s2p1)

	fmt.Fprintf(out, "Created puppy              : %#v \n", s2p1)
	fmt.Printf("Reading puppy back         : ")
	s2p2, _ := s2.ReadPuppy(s2p1.ID)

	fmt.Printf("%#v\n", s2p2)
}
