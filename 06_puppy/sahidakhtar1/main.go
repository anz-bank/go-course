package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	var ms Storer = newMapStore()
	p1 := Puppy{0, "Bulldog", "White", "100"}
	p2 := Puppy{0, "Poddle", "Black", "200"}
	p1Id := ms.CreatePuppy(p1)
	p2Id := ms.CreatePuppy(p2)
	rp1 := ms.ReadPuppy(p1Id)
	fmt.Fprintln(out, rp1)
	p1.Value = "300"
	ms.UpdatePuppy(p1Id, p1)
	rp1 = ms.ReadPuppy(p1Id)
	fmt.Println(rp1)
	rp2 := ms.ReadPuppy(p2Id)
	fmt.Println(rp2)
	delete := ms.DeletePuppy(p2Id)
	fmt.Println(delete)

	var s Storer = newSyncStore()
	p4 := Puppy{0, "Beagle", "White", "400"}
	p5 := Puppy{0, "Pug", "Black", "500"}
	p4ID := s.CreatePuppy(p4)
	p5ID := s.CreatePuppy(p5)
	rp4 := s.ReadPuppy(p4ID)
	fmt.Println(rp4)
	p4.Value = "600"
	s.UpdatePuppy(p4ID, p4)
	rp4 = s.ReadPuppy(p4ID)
	fmt.Println(rp4)
	delete = s.DeletePuppy(p5ID)
	fmt.Println(delete)

}
