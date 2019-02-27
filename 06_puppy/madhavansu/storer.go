package main

import "sync"

/**
* Storer implementation in below scenarios
* Dynamic Map ID: MapStore implementation of Storer backed by a map
* Static Map ID: SyncStore implementation of Storer backed by a sync.Map
 */
//Sync store
type syncStore struct {
	sync.Mutex
	sync.Map
}

func newSyncStore() *syncStore {
	return &syncStore{}
}

// Map Store
type mapStore struct {
	ms    map[uint]Puppy
	mapID uint
}

func newMapStore() *mapStore {
	return &mapStore{ms: make(map[uint]Puppy)}
}
