package main

import (
	"fmt"
)

type MapStore struct {
	pmap   map[int]Puppy
	nextID int
}

func NewMapStore() MapStore {
	a := MapStore{pmap: make(map[int]Puppy)}
	return a
}

func (m MapStore) CreatePuppy(p *Puppy) error {
	if p == nil {
		return fmt.Errorf("puppy pointer is nil")
	}
	if p.ID != 0 {
		return fmt.Errorf("trying to create a puppy already initialized with ID %d", p.ID)
	}
	m.nextID++
	p.ID = m.nextID
	m.pmap[p.ID] = *p
	return nil
}

func (m MapStore) ReadPuppy(id int) (Puppy, error) {
	v, ok := m.pmap[id]
	if !ok {
		return Puppy{}, fmt.Errorf("puppy ID %d being read does not exist", id)
	}
	return v, nil
}

func (m MapStore) UpdatePuppy(p Puppy) error {
	if _, ok := m.pmap[p.ID]; !ok {
		return fmt.Errorf("puppy ID %d being updated does not exist", p.ID)
	}
	m.pmap[p.ID] = p
	return nil
}

func (m MapStore) DeletePuppy(id int) error {
	if _, ok := m.pmap[id]; !ok {
		return fmt.Errorf("puppy ID %d being deleted does not exist", id)
	}
	delete(m.pmap, id)
	return nil
}
