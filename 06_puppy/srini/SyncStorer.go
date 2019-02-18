package main

import (
	"fmt"
	"sync"
)

// SyncStore implemenmtation
type SyncStore struct {
	sync.Mutex
	sync.Map
}

//Generate Puppy ID
var (
	id        int
	mu        sync.Mutex
	syncStore *SyncStore
)

//GetSyncStore returns the singleton instance of Syncstore
func GetSyncStore() *SyncStore {
	if syncStore == nil {
		mu.Lock()
		defer mu.Unlock()
		if syncStore == nil {
			syncStore = &SyncStore{}
		}
	}
	return syncStore
}

//CreatePuppy a Puppy
func (syncStr *SyncStore) CreatePuppy(p *Puppy) int {
	syncStr.Lock()
	defer syncStr.Unlock()
	id++
	syncStr.Store(id, PuppyRecord{id, *p})
	return id // the calling program should maintain the reference of puppy ID
}

//ReadPuppy Read a puppy
func (syncStr *SyncStore) ReadPuppy(id int) (*Puppy, error) {
	puppyRecord, err := syncStore.ReadPuppyRecord(id)
	if err == nil {
		return &puppyRecord.puppy, nil
	}
	return nil, err
}

//UpdatePuppy Update a puppy
func (syncStr *SyncStore) UpdatePuppy(id int, puppy *Puppy) error {
	mu.Lock()
	defer mu.Unlock()
	_, err := syncStore.ReadPuppyRecord(id)
	if err == nil {
		syncStr.Store(id, PuppyRecord{id, *puppy})
		return nil
	}
	return fmt.Errorf("no puppy exists with id %d", id)
}

//DeletePuppy Delete a puppy
func (syncStr *SyncStore) DeletePuppy(id int) error {
	mu.Lock()
	defer mu.Unlock()
	_, err := syncStore.ReadPuppyRecord(id)
	if err == nil {
		syncStr.Delete(id)
		return nil
	}
	return fmt.Errorf("could not delete the puppy with %d", id)
}

//ReadPuppyRecord Read a puppy
func (syncStr *SyncStore) ReadPuppyRecord(id int) (*PuppyRecord, error) {
	val, ok := syncStr.Load(id)
	if ok {
		puppyRecord, _ := val.(PuppyRecord)
		return &puppyRecord, nil
	}
	return nil, fmt.Errorf("no puppy exists with id %d", id)
}
