package main

import (
	"sync"
)

type SyncStore struct {
	uuid uint
	sync.Map
}

// NewSyncStore returns a pointer to a new instance of the SyncStore struct which implements the Storer interface.
func NewSyncStore() Storer {
	return &SyncStore{
		uuid: 0,
	}
}

// Create increments the uuid and adds the provided Puppy struct to the store with this identifier.
func (store *SyncStore) Create(puppy Puppy) uint {
	puppy.ID = store.uuid
	store.Store(puppy.ID, puppy)
	store.uuid++

	return puppy.ID
}

// Read returns the puppy matching the provided uuid.
// An empty Puppy struct is returned if the identifier does not exist.
func (store *SyncStore) Read(id uint) Puppy {
	if value, ok := store.Load(id); ok {
		return value.(Puppy)
	}

	return Puppy{}
}

// Update modifies the puppy matching the provided uuid in the store with the provided Puppy struct.
// It returns a bool whether a matching identifier was modified or not.
func (store *SyncStore) Update(id uint, puppy Puppy) bool {
	if _, ok := store.Load(id); !ok {
		return false
	}

	puppy.ID = id
	store.Store(id, puppy)
	return true
}

// Destroy removes the puppy matching the provided uuid from the store.
// It returns a bool whether a matching identifier was deleted or not.
func (store *SyncStore) Destroy(id uint) bool {
	if _, ok := store.Load(id); !ok {
		return false
	}

	store.Delete(id)
	return true
}
