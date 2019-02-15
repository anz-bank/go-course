package main

type MapStore struct {
	m     map[uint]Puppy
	maxID uint
}

func newMapStore() *MapStore {
	ms := MapStore{}
	ms.m = make(map[uint]Puppy)
	return &ms
}
func (ms *MapStore) CreatePuppy(p Puppy) uint {
	ms.maxID++
	p.ID = ms.maxID
	ms.m[p.ID] = p
	return p.ID
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
