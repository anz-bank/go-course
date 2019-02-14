package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var out io.Writer = os.Stdout

func main() {
	pupStore := MapStore{}
	pupStore.Initialize()

	pup1 := pupStore.Create("Beagle", "brown", "Sam")
	pup2 := pupStore.Read(pup1.ID)
	pup2.Value = "Bingo"
	pupStore.Update(pup2)
	pupStore.Read(1)
	pup4 := pupStore.Create("Bull", "brown", "Tyson")
	pup5 := pupStore.Read(pup4.ID)
	pupStore.Delete(pup1)
	pupStore.Delete(pup5)
	pupStore.Create("rottweiler", "black", "Tipu")

	pupSyncStore := SyncMapStore{}
	pupSyncStore.Initialize()

	pupSync1 := pupSyncStore.Create("Pom", "Black", "Sim")
	pupSync2 := pupSyncStore.Read(pupSync1.ID)
	pupSync2.Value = "Dingo"
	pupSyncStore.Update(pupSync2)
	pupSyncStore.Read(1)
	pupSync4 := pupSyncStore.Create("Pug", "white", "Dim")
	pupSync5 := pupSyncStore.Read(pupSync4.ID)
	pupSyncStore.Delete(pupSync1)
	pupSyncStore.Delete(pupSync5)
	pupSyncStore.Create("Woulf", "Ash", "Togo")
	fmt.Fprintln(out, fmt.Sprintf("Items remaining in Store %d", len(pupStore.m)))
}

type Puppy struct {
	ID     int
	Breed  string
	Colour string
	Value  string
}
type Initializer interface {
	Initialize()
}
type Storer interface {
	Create(string, string, string) Puppy
	Read(int) Puppy
	Update(Puppy) int
	Delete(Puppy) int
}

//Map Store
type MapStore struct {
	m map[int]Puppy
}

//Initializer Implementation for Mapstore
func (store *MapStore) Initialize() {
	store.m = make(map[int]Puppy)
}

//Storer  Implementation for Mapstore
func (store *MapStore) Create(breed string, colour string, value string) Puppy {
	id := len(store.m) + 1
	newPup := Puppy{id, breed, colour, value}
	store.m[id] = newPup
	return newPup
}
func (store MapStore) Update(pup Puppy) int {
	_, ok := store.m[pup.ID]
	if !ok {
		return -1
	}
	store.m[pup.ID] = pup
	return 0
}
func (store MapStore) Delete(pup Puppy) int {
	_, ok := store.m[pup.ID]
	if !ok {
		return -1
	}
	delete(store.m, pup.ID)
	return 0
}

func (store MapStore) Read(id int) Puppy {
	pup := store.m[id]
	return pup
}

//SuncMapstore
type SyncMapStore struct {
	m    sync.Map
	sqid int
}

//Initializer Implementation for SyncMapstore
func (store *SyncMapStore) Initialize() {
	store.m = sync.Map{}
}

//Storer  Implementation for Mapstore
func (store *SyncMapStore) Create(breed string, colour string, value string) Puppy {
	store.sqid++
	id := store.sqid
	newPup := Puppy{id, breed, colour, value}
	store.m.Store(id, newPup)
	return newPup
}
func (store *SyncMapStore) Update(pup Puppy) int {
	_, ok := store.m.Load(pup.ID)
	if !ok {
		return -1
	}
	store.m.Store(pup.ID, pup)
	return 0
}
func (store *SyncMapStore) Delete(pup Puppy) int {
	_, ok := store.m.Load(pup.ID)
	if !ok {
		return -1
	}
	store.m.Delete(pup.ID)
	return 0
}

func (store *SyncMapStore) Read(id int) Puppy {
	pup, _ := store.m.Load(id)
	p, _ := pup.(Puppy)
	return p
}
