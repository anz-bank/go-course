package main

type Puppy struct {
	ID     uint32
	Breed  string
	Colour string
	Value  int
}

type Storer interface {
	CreatePuppy(puppy *Puppy)
	ReadPuppy(ID uint32) *Puppy
	UpdatePuppy(ID uint32, pet *Puppy)
	DeletePuppy(ID uint32) bool
}

type MapStore struct {
	m map[uint32]Puppy
}

func CreateStore() *MapStore {
	return &MapStore{make(map[uint32]Puppy)}
}

func (ms *MapStore) CreatePuppy(puppy *Puppy) {
	ms.m[puppy.ID] = *puppy
}

func (ms *MapStore) ReadPuppy(id uint32) *Puppy {
	puppy := ms.m[id]
	return &puppy
}

func (ms *MapStore) UpdatePuppy(id uint32, puppy *Puppy) {
	ms.m[id] = *puppy
}

func (ms *MapStore) DeletePuppy(id uint32) bool {
	_, ok := ms.m[id]
	if ok {
		delete(ms.m, id)
	}
	return ok
}
