package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var out io.Writer = os.Stdout

type Puppy struct {
	id    uint32
	breed string
	color string
	value string
}

type Storer interface {
	CreatePuppy(*Puppy)
	ReadPuppy(id uint32) *Puppy
	UpdatePuppy(puppy *Puppy)
	DeletePuppy(id uint32) bool
	HasPuppy(id uint32) bool
}

type MapStore map[uint32]Puppy

type SynStore struct {
	store *sync.Map
}

func (ms MapStore) CreatePuppy(pup *Puppy) {
	if !ms.HasPuppy(pup.id) {
		ms[pup.id] = *pup
	}
}
func (ms MapStore) ReadPuppy(id uint32) *Puppy {
	if puppy, has := ms[id]; has {
		return &puppy
	}
	return nil
}

func (ms MapStore) UpdatePuppy(puppy *Puppy) {
	if ms.HasPuppy(puppy.id) {
		ms[puppy.id] = *puppy
	}
}

func (ms MapStore) DeletePuppy(id uint32) bool {
	if ms.HasPuppy(id) {
		delete(ms, id)
		return true
	}
	return false
}
func (ms MapStore) HasPuppy(id uint32) bool {
	_, has := ms[id]
	return has
}

func (ss SynStore) CreatePuppy(pup *Puppy) {
	if !ss.HasPuppy(pup.id) {
		ss.store.Store(pup.id, pup)
	}
}

func (ss SynStore) ReadPuppy(id uint32) *Puppy {
	if pup, ok := ss.store.Load(id); ok {
		return pup.(*Puppy)
	}
	return nil
}

func (ss SynStore) UpdatePuppy(puppy *Puppy) {
	if ss.HasPuppy(puppy.id) {
		ss.store.Store(puppy.id, puppy)
	}
}

func (ss SynStore) DeletePuppy(id uint32) bool {
	if ss.HasPuppy(id) {
		ss.store.Delete(id)
		return true
	}
	return false
}

func (ss SynStore) HasPuppy(id uint32) bool {
	_, ok := ss.store.Load(id)
	return ok
}

func main() {
	tempPuppy := Puppy{id: 1, breed: "dog", color: "Black", value: "1"}
	var store Storer = make(MapStore)
	store.CreatePuppy(&tempPuppy)
	fmt.Fprintln(out, *store.ReadPuppy(tempPuppy.id))
	tempPuppy2 := Puppy{id: 1, breed: "cat", color: "Brown", value: "1"}
	var store2 Storer = SynStore{store: &sync.Map{}}
	store2.CreatePuppy(&tempPuppy2)
	fmt.Fprintln(out, *store2.ReadPuppy(tempPuppy.id))
}
