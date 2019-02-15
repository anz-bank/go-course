package main

import (
	"fmt"
)

func main() {
	var ms Storer = newMapStore()
	// var ms Storer = newMapStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}
	p2 := Puppy{2, "Poddle", "Black", "200"}
	ms.CreatePuppy(p1)
	ms.CreatePuppy(p2)
	rp1 := ms.ReadPuppy(1)
	fmt.Println(rp1)
	p1.Value = "300"
	ms.UpdatePuppy(1, p1)
	rp1 = ms.ReadPuppy(1)
	fmt.Println(rp1)
	rp2 := ms.ReadPuppy(2)
	fmt.Println(rp2)
	delete := ms.DeletePuppy(2)
	fmt.Println(delete)

	var s Storer = newSyncStore()
	p4 := Puppy{4, "Beagle", "White", "400"}
	p5 := Puppy{5, "Pug", "Black", "500"}
	s.CreatePuppy(p4)
	s.CreatePuppy(p5)
	rp4 := s.ReadPuppy(4)
	fmt.Println(rp4)
	p4.Value = "600"
	s.UpdatePuppy(4, p4)
	rp4 = s.ReadPuppy(4)
	fmt.Println(rp4)
	delete = s.DeletePuppy(5)
	fmt.Println(delete)

}
