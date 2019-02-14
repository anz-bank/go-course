package main

// MapStore struct
type MapStore struct {
	m     map[int]Puppy
	maxID int
}

// NewMapStore constructor
func NewMapStore() *MapStore {
	return &MapStore{}
}

//CreatePuppy to create
func (ms *MapStore) CreatePuppy(pup Puppy) int {
	ms.m = make(map[int]Puppy)
	ms.m[ms.maxID] = pup
	pup.ID = ms.maxID
	ms.maxID++
	return pup.ID
}

//ReadPuppy to read
func (ms *MapStore) ReadPuppy(pupID int) Puppy {
	return ms.m[pupID]
}

//UpdatePuppy to read
func (ms *MapStore) UpdatePuppy(pupID int, pup Puppy) {
	ms.m[pupID] = pup
	//return pup
}

//DeletePuppy to delete
func (ms *MapStore) DeletePuppy(pupID int) bool {
	delete(ms.m, pupID)
	return true
}
