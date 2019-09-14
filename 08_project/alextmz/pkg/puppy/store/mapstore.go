package store

import (
	pp "github.com/anz-bank/go-course/08_project/alextmz/pkg/puppy"
)

type MapStore map[pp.Pid]*pp.Puppy

func NewmapStore() *MapStore {
	a := MapStore{}
	return &a
}

func (m *MapStore) CreatePuppy(p *pp.Puppy) error {
	if _, ok := (*m)[p.ID]; !ok {
		if p.Value < 0 {
			return pp.NewError(pp.ErrValueLessThanZero)
		}
		(*m)[p.ID] = p
		return nil
	}
	return pp.NewError(pp.ErrIDBeingCreatedAlreadyExists)
}

func (m *MapStore) ReadPuppy(id pp.Pid) (*pp.Puppy, error) {
	if v, ok := (*m)[id]; ok {
		return v, nil
	}
	return nil, pp.NewError(pp.ErrIDBeingReadDoesNotExist)
}

func (m *MapStore) UpdatePuppy(id pp.Pid, p *pp.Puppy) error {
	if _, ok := (*m)[id]; ok {
		if p.Value < 0 {
			return pp.NewError(pp.ErrValueLessThanZero)
		}
		(*m)[id] = p
		return nil
	}
	return pp.NewError(pp.ErrIDBeingUpdatedDoesNotExist)

}

func (m *MapStore) DeletePuppy(id pp.Pid) error {
	if _, ok := (*m)[id]; ok {
		delete(*m, id)
		return nil
	}
	return pp.NewError(pp.ErrIDBeingDeletedDoesNotExist)
}
