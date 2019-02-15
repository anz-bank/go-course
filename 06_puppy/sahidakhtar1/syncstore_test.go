package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePuppyInSyncStore(t *testing.T) {
	fmt.Println("TestCreatePuppyInSyncStore")
	//Given
	r := require.New(t)
	var s Storer = newSyncStore()
	p1 := Puppy{0, "Bulldog", "White", "100"}
	p2 := Puppy{0, "Popdle", "Black", "300"}
	p3 := Puppy{0, "Beagle", "Cream", "400"}
	p4 := Puppy{0, "Pug", "Mix", "600"}
	//When
	p1Id := s.CreatePuppy(p1)
	p2Id := s.CreatePuppy(p2)
	p3Id := s.CreatePuppy(p3)
	p4Id := s.CreatePuppy(p4)
	//Then
	r.Equalf(uint(1), p1Id, "ID for puppy doesn't set properly in SyncStore")
	r.Equalf(uint(2), p2Id, "ID for puppy is not set properly in SyncStore")
	r.Equalf(uint(3), p3Id, "ID for puppy is not set properly in SyncStore")
	r.Equalf(uint(4), p4Id, "ID for puppy is not set properly in SyncStore")
}
func TestReadPuppyInSyncStore(t *testing.T) {
	fmt.Println("TestReadPuppyInSyncStore")
	//Given
	r := require.New(t)
	var s Storer = newSyncStore()
	p1 := Puppy{0, "Bulldog", "White", "100"}

	//When
	p1ID := s.CreatePuppy(p1)
	p2 := s.ReadPuppy(p1ID)
	p3 := s.ReadPuppy(2)
	expected := Puppy{1, "Bulldog", "White", "100"}

	//Then
	r.Equalf(expected, p2, "ReadPuppy: Reads some other obejct")
	r.Equalf(Puppy{}, p3, "ReadPuppy: Returns non empty Puppy")
}
func TestUpdatePuppyInSyncStore(t *testing.T) {
	fmt.Println("TestUpdatePuppyInSyncStore")
	//Given
	r := require.New(t)
	var s Storer = newSyncStore()
	p1 := Puppy{0, "Bulldog", "White", "100"}
	p3 := Puppy{2, "Poddle", "black", "100"}

	//When
	p1ID := s.CreatePuppy(p1)
	p2 := s.ReadPuppy(p1ID)
	s.UpdatePuppy(p1ID, p3)
	p4 := s.ReadPuppy(p1ID)

	//Then
	r.Equalf(Puppy{1, "Bulldog", "White", "100"}, p2, "ReadPuppy: Reads some other obejct")
	r.Equalf(p3, p4, "UpdatePuppy: Doesn't update the desired Puppy")
}
func TestDeletePuppyInSyncStore(t *testing.T) {
	fmt.Println("TestDeletePuppyInSyncStore")
	//Given
	r := require.New(t)
	var s Storer = newSyncStore()
	p1 := Puppy{0, "Bulldog", "White", "100"}

	//When
	p1ID := s.CreatePuppy(p1)
	delete := s.DeletePuppy(p1ID)
	p3 := s.ReadPuppy(p1ID)
	deleteAgain := s.DeletePuppy(p1ID)

	//Then
	r.Equalf(Puppy{}, p3, "ReadPuppy: Returns non empty Puppy")
	r.Equalf(true, delete, "Fails to delete the desired puppy")
	r.Equalf(false, deleteAgain, "When Puppy is not there it fails to delete")
}
