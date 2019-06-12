package main

import (
	"errors"
)

func NewMemStore() MemStore {
	return MemStore{}
}

func (m MemStore) CreatePuppy(puppy *Puppy) int {
	id := len(m)
	m[id] = puppy
	return id
}

func (m MemStore) ReadPuppy(id int) (*Puppy, error) {
	puppy, ok := m[id]
	if !ok {
		return nil, errors.New("doesn't exists")
	}
	return puppy, nil
}

func (m MemStore) UpdatePuppy(id int, puppy *Puppy) error {
	m[id] = puppy
	return nil
}

func (m MemStore) DeletePuppy(id int) (bool, error) {
	if _, ok := m[id]; !ok {
		return false, errors.New("doesn't exist")
	}
	delete(m, id)
	return true, nil
}
