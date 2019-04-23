package store

import (
	"fmt"

	types "github.com/anz-bank/go-training/08_project/mohankrishna/pkg/mohankrishna-puppy"
)

type MapStore struct {
	m map[uint32]*types.Puppy
}

func NewMapStore() *MapStore {
	return &MapStore{make(map[uint32]*types.Puppy)}
}

func (ms *MapStore) CreatePuppy(puppy *types.Puppy) error {
	if _, exists := ms.m[puppy.ID]; exists {
		return &types.Error{Code: types.ErrDuplicate,
			Message: fmt.Sprintf("A puppy with ID: %d already exists", puppy.ID)}
	}
	ms.m[puppy.ID] = puppy
	return nil
}

func (ms *MapStore) ReadPuppy(id uint32) (*types.Puppy, error) {
	puppy, ok := ms.m[id]
	if !ok {
		return nil, &types.Error{Code: types.ErrNotFound, Message: fmt.Sprintf("No puppy exists with id %d", id)}
	}
	return puppy, nil
}

func (ms *MapStore) UpdatePuppy(id uint32, puppy *types.Puppy) error {
	if id != puppy.ID {
		return &types.Error{Code: types.ErrInvalidInput,
			Message: fmt.Sprintf("The id:%d passed and the puppy's id:%d doesnt match", id, puppy.ID)}
	}
	ms.m[id] = puppy
	return nil
}

func (ms *MapStore) DeletePuppy(id uint32) error {
	_, ok := ms.m[id]
	if ok {
		delete(ms.m, id)
		return nil
	}
	return &types.Error{Code: types.ErrNotFound, Message: fmt.Sprintf("No puppy exists with id %d", id)}
}
