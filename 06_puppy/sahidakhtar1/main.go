package main

import (
	"fmt"
	"sync"
)

type Puppy struct {
	ID                   uint
	Breed, Colour, Value string
}

type Storer interface {
	CreatePuppy(Puppy)
	ReadPuppy(uint) Puppy
	UpdatePuppy(uint, Puppy)
	DeletePuppy(uint) bool
}

//Map Store Starts
type MapStore struct {
	m map[uint]Puppy
}

func newMapStore() *MapStore {
	ms := MapStore{}
	ms.m = make(map[uint]Puppy, 0)
	return &ms
}
func (ms *MapStore) CreatePuppy(p Puppy) {
	_, ok := ms.m[p.ID]
	if ok {
		return
	}
	ms.m[p.ID] = p
}
func (ms *MapStore) ReadPuppy(id uint) Puppy {
	return ms.m[id]
}
func (ms *MapStore) UpdatePuppy(id uint, p Puppy) {
	ms.m[id] = p
}

func (ms *MapStore) DeletePuppy(id uint) bool {
	_, ok := ms.m[id]
	if !ok {
		return false
	}
	delete(ms.m, id)
	return true
}

//Map Store Ends

//Sync Store Ends
type SyncStore struct {
	sync.Mutex
	sync.Map
}

func newSyncStore() *SyncStore {
	return &SyncStore{}
}
func (m *SyncStore) CreatePuppy(p Puppy) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.Load(p.ID); !ok {
		m.Store(p.ID, p)
	}
}

func (m *SyncStore) ReadPuppy(id uint) Puppy {
	p, ok := m.Load(id)
	if !ok {
		return Puppy{}
	}
	puppy := p.(Puppy)
	return puppy
}
func (m *SyncStore) UpdatePuppy(id uint, p Puppy) {
	m.Store(id, p)
}

func (m *SyncStore) DeletePuppy(id uint) bool {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.Load(id); !ok {
		return false
	}
	m.Delete(id)
	return true
}

//Sync Store Ends

func main() {
	ms := newMapStore()
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
	ms.DeletePuppy(2)
	rp2 = ms.ReadPuppy(2)
	fmt.Println(rp2)

	s := newSyncStore()
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
	s.DeletePuppy(5)
	rp5 := s.ReadPuppy(5)
	fmt.Println(rp5)

}
