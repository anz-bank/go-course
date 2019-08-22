package main

import (
	"fmt"
	"sync"
)

type SyncStore struct {
	sync.Map
	sync.RWMutex
	nextID uint32
}

func InitSyncStore() *SyncStore {
	return &SyncStore{
		nextID: 1,
	}
}

func (store *SyncStore) CreatePuppy(puppy *Puppy) error {
	if puppy.ID == 0 {
		puppy.ID = store.nextID
		store.nextID++
	}
	if _, ok := store.Load(puppy.ID); !ok {
		store.Store(puppy.ID, puppy)
		return nil
	}
	return fmt.Errorf("store already has a puppy with id : %d", puppy.ID)
}

func (store *SyncStore) ReadPuppy(id uint32) (*Puppy, error) {
	if p, ok := store.Load(id); ok {
		return p.(*Puppy), nil
	}
	return nil, fmt.Errorf("could not load puppy %d", id)
}

func (store *SyncStore) UpdatePuppy(id uint32, puppy *Puppy) error {
	if puppy == nil {
		return fmt.Errorf("cant update to a nil puppy")
	}
	if p, ok := store.Load(id); ok {
		if p.(*Puppy).ID == puppy.ID {
			store.Store(id, puppy.Clone())
			return nil
		}
		return fmt.Errorf("bad puppy : %d != %d", p.(*Puppy).ID, id)
	}
	return fmt.Errorf("could not find puppy id = %d", id)
}

func (store *SyncStore) DeletePuppy(id uint32) (bool, error) {
	if _, ok := store.Load(id); ok {
		store.Delete(id)
		return true, nil
	}
	return false, fmt.Errorf("could not find puppy id = %d", id)
}
