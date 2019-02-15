package main

import "fmt"

type MapStore struct {
	m map[uint32]*Puppy
}

func NewMapStore() *MapStore {
	return &MapStore{make(map[uint32]*Puppy)}
}

func (ms *MapStore) CreatePuppy(puppy *Puppy) error {
	if _, exists := ms.m[puppy.ID]; exists {
		return &Error{ErrDuplicate, fmt.Sprintf("A puppy with ID: %d already exists", puppy.ID)}
	}
	ms.m[puppy.ID] = puppy
	return nil
}

func (ms *MapStore) ReadPuppy(id uint32) (*Puppy, error) {
	puppy, ok := ms.m[id]
	if !ok {
		return nil, &Error{ErrNotFound, fmt.Sprintf("No puppy exists with id %d", id)}
	}
	return puppy, nil
}

func (ms *MapStore) UpdatePuppy(id uint32, puppy *Puppy) error {
	if id != puppy.ID {
		return &Error{ErrInvalidInput, fmt.Sprintf("The id:%d passed and the puppy's id:%d doesnt match", id, puppy.ID)}
	}
	ms.m[id] = puppy
	return nil
}

func (ms *MapStore) DeletePuppy(id uint32) (bool, error) {
	_, ok := ms.m[id]
	if ok {
		delete(ms.m, id)
		return ok, nil
	}
	return ok, &Error{ErrDuplicate, fmt.Sprintf("No puppy exists with id %d", id)}
}
