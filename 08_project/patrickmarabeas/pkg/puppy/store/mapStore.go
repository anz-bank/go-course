package store

import (
	p "github.com/anz-bank/go-course/08_project/patrickmarabeas/pkg/puppy"
)

type MapStore struct {
	uuid  int
	store map[int]p.Puppy
}

// NewMapStore returns a pointer to a new instance of the MapStore struct which implements the Storer interface.
func NewMapStore() Storer {
	return &MapStore{
		uuid:  0,
		store: map[int]p.Puppy{},
	}
}

// Create increments the uuid and adds the provided Puppy struct to the store with this identifier.
func (store *MapStore) Create(puppy p.Puppy) (int, error) {
	if puppy.Value < 0 {
		return -1, p.NewError(p.NegativeValue)
	}

	puppy.ID = store.uuid
	store.store[puppy.ID] = puppy
	store.uuid++

	return puppy.ID, nil
}

// Read returns the puppy matching the provided uuid.
// An empty Puppy struct is returned if the identifier does not exist.
func (store *MapStore) Read(id int) (p.Puppy, error) {
	if _, ok := store.store[id]; ok {
		return store.store[id], nil
	}

	return p.Puppy{}, p.NewError(p.IDNotFound)
}

// Update modifies the puppy matching the provided uuid in the store with the provided Puppy struct.
// It returns a bool whether a matching identifier was modified or not.
func (store *MapStore) Update(id int, puppy p.Puppy) (bool, error) {
	if _, ok := store.store[id]; !ok {
		return false, p.NewError(p.IDNotFound)
	}
	if puppy.Value < 0 {
		return false, p.NewError(p.NegativeValue)
	}

	puppy.ID = id
	store.store[id] = puppy
	return true, nil
}

// Destroy removes the puppy matching the provided uuid from the store.
// It returns a bool whether a matching identifier was deleted or not.
func (store *MapStore) Destroy(id int) (bool, error) {
	if _, ok := store.store[id]; !ok {
		return false, p.NewError(p.IDNotFound)
	}

	delete(store.store, id)
	return true, nil
}
