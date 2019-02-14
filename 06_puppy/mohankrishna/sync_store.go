package main

import (
	"sync"
)

type SyncStore struct {
	sm sync.Map
}

func CreateSyncStore() *SyncStore {
	return &SyncStore{}
}

func (ss *SyncStore) CreatePuppy(puppy *Puppy) {
	ss.sm.Store(puppy.ID, *puppy)
}

func (ss *SyncStore) ReadPuppy(id uint32) *Puppy {
	v, ok := ss.sm.Load(id)
	if ok {
		if puppy, ok := v.(Puppy); ok {
			return &puppy
		}
	}
	//This should return an error, we will be implementing the errors in next lab
	return &Puppy{}
}

func (ss *SyncStore) UpdatePuppy(id uint32, puppy *Puppy) {
	ss.sm.Store(id, *puppy)
}

func (ss *SyncStore) DeletePuppy(id uint32) bool {
	_, ok := ss.sm.Load(id)
	if ok {
		ss.sm.Delete(id)
	}
	return ok
}
