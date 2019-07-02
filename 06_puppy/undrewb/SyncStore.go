package main

import (
	"fmt"
	"sync"
)

type SyncStore struct {
	sync.Map
}

func (store *SyncStore) CreatePuppy(p *Puppy) error {
	if _, ok := store.Load(p.ID); !ok {
		store.Store(p.ID, p)
		return nil
	}
	return fmt.Errorf("store already has a puppy with id : %d", p.ID)
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
