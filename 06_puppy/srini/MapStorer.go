package main

import (
	"fmt"
	"strconv"
	"sync"
)

// MapStore implemenmtation
type MapStore struct {
	ms map[string]PuppyRecord
}

//Generate Puppy ID
var (
	ID       int
	mutex    sync.Mutex
	mapStore *MapStore
)

//GetMapStore returns the singleton instance of MapStore
func GetMapStore() *MapStore {
	if mapStore == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if mapStore == nil {
			ms := make(map[string]PuppyRecord)
			mapStore = &MapStore{ms}
		}

	}
	return mapStore
}

//CreatePuppy a Puppy
func (mapStore *MapStore) CreatePuppy(p *Puppy) int {
	ID++
	mapStore.ms[strconv.Itoa(ID)] = PuppyRecord{id, *p}
	return ID
}

//ReadPuppy Read a puppy
func (mapStore *MapStore) ReadPuppy(id int) (*Puppy, error) {
	puppyRecord, err := mapStore.ReadPuppyRecord(id)
	if err == nil {
		return &puppyRecord.puppy, nil
	}
	return nil, err
}

//ReadPuppyRecord Read a puppy
func (mapStore *MapStore) ReadPuppyRecord(id int) (*PuppyRecord, error) {
	puppyRecord, ok := mapStore.ms[strconv.Itoa(id)]
	if ok {
		return &puppyRecord, nil
	}
	return nil, fmt.Errorf("no puppy exists with id %d", id)
}

//UpdatePuppy Update a puppy
func (mapStore *MapStore) UpdatePuppy(id int, puppy *Puppy) error {
	puppyRecord, err := mapStore.ReadPuppyRecord(id)
	if err == nil {
		puppyRecord.puppy = *puppy
		mapStore.ms[strconv.Itoa(ID)] = *puppyRecord
		return nil
	}
	return fmt.Errorf("no puppy exists with id %d", id)
}

//DeletePuppy Delete a puppy
func (mapStore *MapStore) DeletePuppy(id int) error {
	_, err := mapStore.ReadPuppyRecord(id)
	if err == nil {
		delete(mapStore.ms, strconv.Itoa(ID))
		return nil
	}

	return fmt.Errorf("could not delete the puppy with %d", id)
}
