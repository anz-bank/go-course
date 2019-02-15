package main

type MapStore struct {
	m map[uint]Puppy
}

func newMapStore() *MapStore {
	ms := MapStore{}
	ms.m = make(map[uint]Puppy)
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
