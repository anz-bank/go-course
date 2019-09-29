package store

import (
	"github.com/anz-bank/go-course/10_rest/alextmz/pkg/puppy"
)

type MapStore map[int]puppy.Puppy

func NewmapStore() MapStore {
	a := MapStore{}
	return a
}

func (m MapStore) CreatePuppy(p *puppy.Puppy) error {
	switch {
	case p == nil:
		return puppy.NewError(puppy.Err400BadRequest)
	case p.ID == 0:
		p.ID = len(m) + 1
		m[p.ID] = *p
		return nil
	default:
		return puppy.NewError(puppy.Err400BadRequest)
	}
}

func (m MapStore) ReadPuppy(id int) (puppy.Puppy, error) {
	if id < 0 {
		return puppy.Puppy{}, puppy.NewError(puppy.Err400BadRequest)
	}
	if v, ok := m[id]; ok {
		return v, nil
	}
	return puppy.Puppy{}, puppy.NewError(puppy.Err404NotFound)
}

func (m MapStore) UpdatePuppy(p puppy.Puppy) error {
	if p.ID < 0 {
		return puppy.NewError(puppy.Err400BadRequest)
	}
	if _, ok := m[p.ID]; ok {
		m[p.ID] = p
		return nil
	}
	return puppy.NewError(puppy.Err404NotFound)
}

func (m MapStore) DeletePuppy(id int) error {
	if id < 0 {
		return puppy.NewError(puppy.Err400BadRequest)
	}
	if _, ok := m[id]; ok {
		delete(m, id)
		return nil
	}
	return puppy.NewError(puppy.Err404NotFound)
}
