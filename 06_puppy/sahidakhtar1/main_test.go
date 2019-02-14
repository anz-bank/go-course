package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainOutPut(t *testing.T) {
	main()
}

func TestCreatePuppyInMapStore(t *testing.T) {
	//Given
	r := require.New(t)
	m := newMapStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}
	//When
	p2 := m.ReadPuppy(1)
	m.CreatePuppy(p1)
	p3 := m.ReadPuppy(1)
	p4 := Puppy{1, "Bulldog", "White", "100"}
	m.CreatePuppy(p4)
	p5 := m.ReadPuppy(1)
	//Then
	r.Equalf(Puppy{}, p2, "CreatePuppy Fail")
	r.Equalf(p3, p1, "Unexpected output in main()")
	r.Equalf(p5, p1, "Unexpected output in main()")
}

func TestReadPuppyInMapStore(t *testing.T) {
	//Given
	r := require.New(t)
	m := newMapStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}
	m.CreatePuppy(p1)
	//When
	p2 := m.ReadPuppy(1)
	p3 := m.ReadPuppy(2)

	//Then
	r.Equalf(p2, p1, "Unexpected output in main()")
	r.Equalf(Puppy{}, p3, "Unexpected output in main()")

}
func TestUpdatePuppyInMapStore(t *testing.T) {
	//Given
	r := require.New(t)
	m := newMapStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}
	//When
	m.CreatePuppy(p1)
	p2 := m.ReadPuppy(1)
	p3 := Puppy{1, "Poddle", "black", "100"}
	m.UpdatePuppy(1, p3)
	p4 := m.ReadPuppy(1)
	//Then
	r.Equalf(p1, p2, "Unexpected output in main()")
	r.Equalf(p3, p4, "Unexpected output in main()")
}
func TestDeletePuppyInMapStore(t *testing.T) {
	//Given
	r := require.New(t)
	m := newMapStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}
	//When
	m.CreatePuppy(p1)
	p2 := m.ReadPuppy(1)
	delete := m.DeletePuppy(1)
	p3 := m.ReadPuppy(1)
	deleteAgain := m.DeletePuppy(1)
	//Then
	r.Equalf(p1, p2, "Unexpected output in main()")
	r.Equalf(Puppy{}, p3, "Unexpected output in main()")
	r.Equalf(true, delete, "Unexpected output in main()")
	r.Equalf(false, deleteAgain, "Unexpected output in main()")
}

func TestCreatePuppyInSyncStore(t *testing.T) {
	//Given
	r := require.New(t)
	s := newSyncStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}
	//When
	p2 := s.ReadPuppy(1)
	s.CreatePuppy(p1)
	p3 := s.ReadPuppy(1)
	//Then
	r.Equalf(Puppy{}, p2, "CreatePuppy Fail")
	r.Equalf(p3, p1, "Unexpected output in main()")
}
func TestReadPuppyInSyncStore(t *testing.T) {
	//Given
	r := require.New(t)
	m := newSyncStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}
	m.CreatePuppy(p1)
	//When
	p2 := m.ReadPuppy(1)
	p3 := m.ReadPuppy(2)

	//Then
	r.Equalf(p2, p1, "Unexpected output in main()")
	r.Equalf(Puppy{}, p3, "Unexpected output in main()")

}
func TestUpdatePuppyInSyncStore(t *testing.T) {
	//Given
	r := require.New(t)
	s := newSyncStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}
	//When
	s.CreatePuppy(p1)
	p2 := s.ReadPuppy(1)
	p3 := Puppy{1, "Poddle", "black", "100"}
	s.UpdatePuppy(1, p3)
	p4 := s.ReadPuppy(1)
	//Then
	r.Equalf(p1, p2, "Unexpected output in main()")
	r.Equalf(p3, p4, "Unexpected output in main()")
}
func TestDeletePuppyInSyncStore(t *testing.T) {
	//Given
	r := require.New(t)
	s := newSyncStore()
	p1 := Puppy{1, "Bulldog", "White", "100"}
	//When
	s.CreatePuppy(p1)
	p2 := s.ReadPuppy(1)
	delete := s.DeletePuppy(1)
	p3 := s.ReadPuppy(1)
	deleteAgain := s.DeletePuppy(1)
	//Then
	r.Equalf(p1, p2, "Unexpected output in main()")
	r.Equalf(Puppy{}, p3, "Unexpected output in main()")
	r.Equalf(true, delete, "Unexpected output in main()")
	r.Equalf(false, deleteAgain, "Unexpected output in main()")
}
