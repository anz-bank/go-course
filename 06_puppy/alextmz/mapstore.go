package main

import "fmt"

type MapStore map[int]Puppy

func NewmapStore() MapStore {
	a := MapStore{}
	return a
}

// ID a Puppy and Store
func (m MapStore) CreatePuppy(p *Puppy) error {
	if p.ID == 0 {
		p.ID = len(m) + 1
		m[p.ID] = *p
		return nil
	}
	if _, ok := m[p.ID]; ok {
		return fmt.Errorf("puppy ID %d being created already exists", p.ID)
	}
	return fmt.Errorf("trying to create a puppy that is already identified with ID %d", p.ID)
}

func (m MapStore) ReadPuppy(id int) (Puppy, error) {
	if v, ok := m[id]; ok {
		return v, nil
	}
	return Puppy{}, fmt.Errorf("puppy ID %d being read does not exist", id)
}

func (m MapStore) UpdatePuppy(p Puppy) error {
	if _, ok := m[p.ID]; ok {
		m[p.ID] = p
		return nil
	}
	return fmt.Errorf("puppy ID %d being updated does not exist", p.ID)
}

func (m MapStore) DeletePuppy(id int) error {
	if _, ok := m[id]; ok {
		delete(m, id)
		return nil
	}
	return fmt.Errorf("puppy ID %d being deleted does not exist", id)
}
