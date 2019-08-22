package main

import (
	"errors"
	"fmt"
)

type MapStore struct {
	store  map[uint32]*Puppy
	nextID uint32
}

func InitMapStore() *MapStore {
	return &MapStore{
		store:  map[uint32]*Puppy{},
		nextID: 1,
	}
}

func (msp *MapStore) CreatePuppy(p *Puppy) error {
	if p.ID == 0 {
		p.ID = msp.nextID
		msp.nextID++
	}

	if _, exist := msp.store[p.ID]; exist {
		return fmt.Errorf("puppy %d already exists", p.ID)
	}

	msp.store[p.ID] = p.Clone()
	return nil
}

func (msp *MapStore) ReadPuppy(id uint32) (*Puppy, error) {
	p, err := msp.store[id]
	if !err {
		return nil, errors.New("no such puppy")
	}
	return p, nil
}

func (msp *MapStore) DeletePuppy(id uint32) (bool, error) {
	p, exist := msp.store[id]
	if !exist {
		return false, errors.New("no such puppy")
	}
	delete(msp.store, p.ID)
	return true, nil
}

func (msp *MapStore) UpdatePuppy(id uint32, puppy *Puppy) error {
	if puppy == nil {
		return fmt.Errorf("cant update to a nil puppy")
	}
	p, exist := msp.store[id]
	if !exist {
		return errors.New("no such puppy")
	}
	if p.ID != puppy.ID {
		return errors.New("bad puppy")
	}
	msp.store[id] = puppy.Clone()
	return nil
}
