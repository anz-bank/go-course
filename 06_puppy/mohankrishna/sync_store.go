package main

import (
	"fmt"
	"sync"
)

type SyncStore struct {
	sm sync.Map
}

func NewSyncStore() *SyncStore {
	return &SyncStore{}
}

func (ss *SyncStore) CreatePuppy(puppy *Puppy) error {
	if _, ok := ss.sm.Load(puppy.ID); ok {
		return &Error{ErrDuplicate, fmt.Sprintf("A puppy with ID: %d already exists", puppy.ID)}
	}
	ss.sm.Store(puppy.ID, puppy)
	return nil
}

func (ss *SyncStore) ReadPuppy(id uint32) (*Puppy, error) {
	v, ok := ss.sm.Load(id)
	if ok {
		puppy, ok := v.(*Puppy)
		if ok {
			return puppy, nil
		}
	}
	return nil, &Error{ErrNotFound, fmt.Sprintf("A puppy with ID: %d doesn't exist", id)}
}

func (ss *SyncStore) UpdatePuppy(id uint32, puppy *Puppy) error {
	if id != puppy.ID {
		return &Error{ErrInvalidInput, fmt.Sprintf("The id:%d passed and the puppy's id:%d doesnt match", id, puppy.ID)}
	}
	ss.sm.Store(id, puppy)
	return nil
}

func (ss *SyncStore) DeletePuppy(id uint32) (bool, error) {
	_, ok := ss.sm.Load(id)
	if ok {
		ss.sm.Delete(id)
		return ok, nil
	}
	return ok, &Error{ErrDuplicate, fmt.Sprintf("No puppy exists with id %d", id)}
}
