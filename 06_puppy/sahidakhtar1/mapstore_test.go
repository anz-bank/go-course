package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePuppyInMapStore(t *testing.T) {
	//Given
	r := require.New(t)
	var m Storer = newMapStore()
	p1 := Puppy{0, "Bulldog", "White", "100"}
	p2 := Puppy{0, "Bulldog", "White", "100"}
	p3 := Puppy{0, "Bulldog", "White", "100"}
	p4 := Puppy{0, "Bulldog", "White", "100"}
	//When
	p1Id := m.CreatePuppy(p1)
	p2Id := m.CreatePuppy(p2)
	p3Id := m.CreatePuppy(p3)
	p4Id := m.CreatePuppy(p4)
	//Then
	r.Equalf(uint(1), p1Id, "ID for puppy is not set properly")
	r.Equalf(uint(2), p2Id, "ID2 for puppy is not set properly")
	r.Equalf(uint(3), p3Id, "ID for puppy is not set properly")
	r.Equalf(uint(4), p4Id, "ID for puppy is not set properly")
}

func TestReadPuppyInMapStore(t *testing.T) {
	//Given
	r := require.New(t)
	var m Storer = newMapStore()
	p1 := Puppy{0, "Bulldog", "White", "100"}
	//When
	p1ID := m.CreatePuppy(p1)
	p2 := m.ReadPuppy(p1ID)
	p3 := m.ReadPuppy(2)
	expected := Puppy{1, "Bulldog", "White", "100"}
	//Then
	r.Equalf(expected, p2, "ReadPuppy: Reads some other obejct")
	r.Equalf(Puppy{}, p3, "ReadPuppy: Returns non empty Puppy")
}
func TestUpdatePuppyInMapStore(t *testing.T) {
	//Given
	r := require.New(t)
	var m Storer = newMapStore()
	p1 := Puppy{0, "Bulldog", "White", "100"}
	p3 := Puppy{2, "Poddle", "black", "100"}
	//When
	p1ID := m.CreatePuppy(p1)
	p2 := m.ReadPuppy(p1ID)
	m.UpdatePuppy(p1ID, p3)
	p4 := m.ReadPuppy(p1ID)
	//Then
	r.Equalf(Puppy{1, "Bulldog", "White", "100"}, p2, "ReadPuppy: Reads some other obejct")
	r.Equalf(p3, p4, "UpdatePuppy: Doesn't update the desired Puppy")
}
func TestDeletePuppyInMapStore(t *testing.T) {
	//Given
	r := require.New(t)
	var m Storer = newMapStore()
	p1 := Puppy{0, "Bulldog", "White", "100"}
	//When
	p1ID := m.CreatePuppy(p1)
	delete := m.DeletePuppy(p1ID)
	p3 := m.ReadPuppy(p1ID)
	deleteAgain := m.DeletePuppy(p1ID)
	//Then
	r.Equalf(Puppy{}, p3, "ReadPuppy: Returns non empty Puppy")
	r.Equalf(true, delete, "Fails to delete the desired puppy")
	r.Equalf(false, deleteAgain, "When Puppy is not there it fails to delete")
}
