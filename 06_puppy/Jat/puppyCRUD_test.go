package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLettersMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "Items remaining in Store 1\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestMapCreationSuccess(t *testing.T) {
	r := require.New(t)
	// When
	pupStore := MapStore{}
	pupStore.Initialize()
	// Then
	expected := 0
	actual := len(pupStore.m)
	r.EqualValues(expected, actual, "Unexpected output in TestMapCreation Success")
}
func TestCreatandReadPuppySuccess(t *testing.T) {
	r := require.New(t)
	// When
	pupStore := MapStore{}
	pupStore.Initialize()
	npup := pupStore.Create("BullDog", "brown", "Tyson")
	pup := pupStore.Read(npup.ID)
	// Then
	expected := "Tyson"
	actual := pup.Value
	r.EqualValues(expected, actual, "Unexpected output in TestMapCreation Success")
}
func TestCreatandReadandDeletePuppySuccess(t *testing.T) {
	r := require.New(t)
	// When
	pupStore := MapStore{}
	pupStore.Initialize()
	npup := pupStore.Create("BullDog", "brown", "Tyson")
	pupStore.Delete(npup)
	// Then
	expected := 0
	actual := len(pupStore.m)
	r.EqualValues(expected, actual, "Unexpected output in TestMapCreation Success")
}
func TestCreatandUpdateandReadPuppySuccess(t *testing.T) {
	r := require.New(t)
	// When
	pupStore := MapStore{}
	pupStore.Initialize()
	npup := pupStore.Create("BullDog", "brown", "Tyson")
	npup.Value = "Mike"
	pupStore.Update(npup)
	upupup := pupStore.Read(npup.ID)
	// Then
	expected := "Mike"
	actual := upupup.Value
	r.EqualValues(expected, actual, "Unexpected output in TestMapCreation Success")
}

func TestUpdateNonExisting(t *testing.T) {
	r := require.New(t)
	// When
	pupStore := MapStore{}
	pupStore1 := MapStore{}
	pupStore.Initialize()
	pupStore1.Initialize()
	npup := pupStore.Create("BullDog", "brown", "Tyson")
	npup.Value = "Mike"
	res := pupStore1.Update(npup)
	// Then
	expected := -1
	actual := res
	r.EqualValues(expected, actual, "Unexpected output in TestMapCreation Success")
}

func TestDeleteNonExisting(t *testing.T) {
	r := require.New(t)
	// When
	pupStore := MapStore{}
	pupStore1 := MapStore{}
	pupStore.Initialize()
	pupStore1.Initialize()
	npup := pupStore.Create("BullDog", "brown", "Tyson")
	npup.Value = "Mike"
	res := pupStore1.Delete(npup)
	// Then
	expected := -1
	actual := res
	r.EqualValues(expected, actual, "Unexpected output in TestMapCreation Success")
}

// Test for SyncMap
func TestCreatandReadPuppySyncSuccess(t *testing.T) {
	r := require.New(t)
	// When
	pupStore := SyncMapStore{}
	pupStore.Initialize()
	npup := pupStore.Create("BullDog", "brown", "Tyson")
	pup := pupStore.Read(npup.ID)
	// Then
	expected := "Tyson"
	actual := pup.Value
	r.EqualValues(expected, actual, "Unexpected output in TestMapCreation Success")
}
func TestCreatandReadandDeleteSyncPuppySuccess(t *testing.T) {
	r := require.New(t)
	// When
	pupStore := SyncMapStore{}
	pupStore.Initialize()
	npup := pupStore.Create("BullDog", "brown", "Tyson")
	pupStore.Delete(npup)
	npup1 := pupStore.Read(npup.ID)
	// Then
	expected := npup1.Value
	actual := ""
	r.EqualValues(expected, actual, "Unexpected output in TestMapCreation Success")
}
func TestCreatandUpdateandReadSyncPuppySuccess(t *testing.T) {
	r := require.New(t)
	// When
	pupStore := SyncMapStore{}
	pupStore.Initialize()
	npup := pupStore.Create("BullDog", "brown", "Tyson")
	npup.Value = "Mike"
	pupStore.Update(npup)
	upupup := pupStore.Read(npup.ID)
	// Then
	expected := "Mike"
	actual := upupup.Value
	r.EqualValues(expected, actual, "Unexpected output in TestMapCreation Success")
}

func TestUpdateNonSyncExisting(t *testing.T) {
	r := require.New(t)
	// When
	pupStore := SyncMapStore{}
	pupStore1 := SyncMapStore{}
	pupStore.Initialize()
	pupStore1.Initialize()
	npup := pupStore.Create("BullDog", "brown", "Tyson")
	npup.Value = "Mike"
	res := pupStore1.Update(npup)
	// Then
	expected := -1
	actual := res
	r.EqualValues(expected, actual, "Unexpected output in TestMapCreation Success")
}

func TestDeleteNonSyncExisting(t *testing.T) {
	r := require.New(t)
	// When
	pupStore := SyncMapStore{}
	pupStore1 := SyncMapStore{}
	pupStore.Initialize()
	pupStore1.Initialize()
	npup := pupStore.Create("BullDog", "brown", "Tyson")
	npup.Value = "Mike"
	res := pupStore1.Delete(npup)
	// Then
	expected := -1
	actual := res
	r.EqualValues(expected, actual, "Unexpected output in TestMapCreation Success")
}
