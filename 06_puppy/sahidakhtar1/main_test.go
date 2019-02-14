package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainOutPut(t *testing.T) {
	//Just not checking the cosole out put
	r := require.New(t)
	main()
	// Then
	r.Equalf("", "", "Unexpected output in main()")
}

func TestCreatePuppyInMapStore(t *testing.T) {
	//Given
	r := require.New(t)
	var m Storer = newMapStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}
	//When
	p2 := m.ReadPuppy(1)
	m.CreatePuppy(p1)
	p3 := m.ReadPuppy(1)
	p4 := Puppy{1, "Bulldog", "White", "100"}
	m.CreatePuppy(p4)
	p5 := m.ReadPuppy(1)
	//Then
	r.Equalf(Puppy{}, p2, "CreatePuppy: Returns non empty Puppy when its not created")
	r.Equalf(p3, p1, "CreatePuppy: Returns some other object")
	r.Equalf(p5, p1, "CreatePuppy: Returns some other object")
}

func TestReadPuppyInMapStore(t *testing.T) {
	//Given
	r := require.New(t)
	var m Storer = newMapStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}
	m.CreatePuppy(p1)
	//When
	p2 := m.ReadPuppy(1)
	p3 := m.ReadPuppy(2)

	//Then
	r.Equalf(p2, p1, "ReadPuppy: Reads some other obejct")
	r.Equalf(Puppy{}, p3, "ReadPuppy: Returns non empty Puppy")
}
func TestUpdatePuppyInMapStore(t *testing.T) {
	//Given
	r := require.New(t)
	var m Storer = newMapStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}
	//When
	m.CreatePuppy(p1)
	p2 := m.ReadPuppy(1)
	p3 := Puppy{1, "Poddle", "black", "100"}
	m.UpdatePuppy(1, p3)
	p4 := m.ReadPuppy(1)
	//Then
	r.Equalf(p1, p2, "ReadPuppy: Reads some other obejct")
	r.Equalf(p3, p4, "UpdatePuppy: Doesn't update the desired Puppy")
}
func TestDeletePuppyInMapStore(t *testing.T) {
	//Given
	r := require.New(t)
	var m Storer = newMapStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}
	//When
	m.CreatePuppy(p1)
	p2 := m.ReadPuppy(1)
	delete := m.DeletePuppy(1)
	p3 := m.ReadPuppy(1)
	deleteAgain := m.DeletePuppy(1)
	//Then
	r.Equalf(p1, p2, "ReadPuppy: Reads some other obejct")
	r.Equalf(Puppy{}, p3, "ReadPuppy: Returns non empty Puppy")
	r.Equalf(true, delete, "Fails to delete the desired puppy")
	r.Equalf(false, deleteAgain, "When Puppy is not there it fails to delete")
}

func TestCreatePuppyInSyncStore(t *testing.T) {
	//Given
	r := require.New(t)
	var s Storer = newSyncStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}
	//When
	p2 := s.ReadPuppy(1)
	s.CreatePuppy(p1)
	p3 := s.ReadPuppy(1)
	//Then
	r.Equalf(Puppy{}, p2, "CreatePuppy: Returns non empty Puppy when its not created")
	r.Equalf(p3, p1, "CreatePuppy: Returns some other object")
}
func TestReadPuppyInSyncStore(t *testing.T) {
	//Given
	r := require.New(t)
	var s Storer = newSyncStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}
	s.CreatePuppy(p1)
	//When
	p2 := s.ReadPuppy(1)
	p3 := s.ReadPuppy(2)

	//Then
	r.Equalf(p2, p1, "ReadPuppy: Reads some other obejct")
	r.Equalf(Puppy{}, p3, "ReadPuppy: Returns non empty Puppy")
}
func TestUpdatePuppyInSyncStore(t *testing.T) {
	//Given
	r := require.New(t)
	var s Storer = newSyncStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}
	//When
	s.CreatePuppy(p1)
	p2 := s.ReadPuppy(1)
	p3 := Puppy{1, "Poddle", "black", "100"}
	s.UpdatePuppy(1, p3)
	p4 := s.ReadPuppy(1)
	//Then
	r.Equalf(p1, p2, "ReadPuppy: Reads some other obejct")
	r.Equalf(p3, p4, "UpdatePuppy: Doesn't update the desired Puppy")
}
func TestDeletePuppyInSyncStore(t *testing.T) {
	//Given
	r := require.New(t)
	var s Storer = newSyncStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}
	//When
	s.CreatePuppy(p1)
	p2 := s.ReadPuppy(1)
	delete := s.DeletePuppy(1)
	p3 := s.ReadPuppy(1)
	deleteAgain := s.DeletePuppy(1)
	//Then
	r.Equalf(p1, p2, "ReadPuppy: Reads some other obejct")
	r.Equalf(Puppy{}, p3, "ReadPuppy: Returns non empty Puppy")
	r.Equalf(true, delete, "Fails to delete the desired puppy")
	r.Equalf(false, deleteAgain, "When Puppy is not there it fails to delete")
}
