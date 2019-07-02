package main

import (
	"errors"
	"fmt"
)

type MapStore map[uint32]*Puppy

func (msp *MapStore) CreatePuppy(p *Puppy) error {
	if _, exist := (*msp)[p.ID]; exist {
		return fmt.Errorf("puppy %d already exists", p.ID)
	}
	(*msp)[p.ID] = p.Clone()
	return nil
}

func (msp *MapStore) ReadPuppy(id uint32) (*Puppy, error) {
	p, err := (*msp)[id]
	if !err {
		return nil, errors.New("no such puppy")
	}
	return p, nil
}

func (msp *MapStore) DeletePuppy(id uint32) (bool, error) {
	p, exist := (*msp)[id]
	if !exist {
		return false, errors.New("no such puppy")
	}
	delete(*msp, p.ID)
	return true, nil
}

func (msp *MapStore) UpdatePuppy(id uint32, puppy *Puppy) error {
	if puppy == nil {
		return fmt.Errorf("cant update to a nil puppy")
	}
	p, exist := (*msp)[id]
	if !exist {
		return errors.New("no such puppy")
	}
	if p.ID != puppy.ID {
		return errors.New("bad puppy")
	}
	(*msp)[id] = puppy.Clone()
	return nil
}
