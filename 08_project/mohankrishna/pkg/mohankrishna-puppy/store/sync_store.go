package store

import (
	"fmt"
	"sync"

	types "github.com/anz-bank/go-training/08_project/mohankrishna/pkg/mohankrishna-puppy"
)

type SyncStore struct {
	sm sync.Map
}

func NewSyncStore() *SyncStore {
	return &SyncStore{}
}

func (ss *SyncStore) CreatePuppy(puppy *types.Puppy) error {
	if _, ok := ss.sm.Load(puppy.ID); ok {
		return &types.Error{Code: types.ErrDuplicate,
			Message: fmt.Sprintf("A puppy with ID: %d already exists", puppy.ID)}
	}
	ss.sm.Store(puppy.ID, puppy)
	return nil
}

func (ss *SyncStore) ReadPuppy(id uint32) (*types.Puppy, error) {
	v, ok := ss.sm.Load(id)
	if ok {
		puppy, ok := v.(*types.Puppy)
		if ok {
			return puppy, nil
		}
	}
	return nil, &types.Error{Code: types.ErrNotFound,
		Message: fmt.Sprintf("A puppy with ID: %d doesn't exist", id)}
}

func (ss *SyncStore) UpdatePuppy(id uint32, puppy *types.Puppy) error {
	if id != puppy.ID {
		return &types.Error{Code: types.ErrInvalidInput,
			Message: fmt.Sprintf("The id:%d passed and the puppy's id:%d doesnt match", id, puppy.ID)}
	}
	ss.sm.Store(id, puppy)
	return nil
}

func (ss *SyncStore) DeletePuppy(id uint32) error {
	_, ok := ss.sm.Load(id)
	if ok {
		ss.sm.Delete(id)
		return nil
	}
	return &types.Error{Code: types.ErrNotFound, Message: fmt.Sprintf("No puppy exists with id %d", id)}
}
