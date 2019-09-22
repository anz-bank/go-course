package store

import (
	"github.com/anz-bank/go-course/08_project/alextmz/pkg/puppy"
)

type MapStore struct {
	pmap   map[int]puppy.Puppy
	nextID int
}

func NewMapStore() *MapStore {
	a := MapStore{pmap: make(map[int]puppy.Puppy)}
	return &a
}

func (m *MapStore) CreatePuppy(p *puppy.Puppy) error {
	if p == nil {
		return puppy.Error{Code: puppy.ErrNilPuppyPointer}
	}
	if p.Value < 0 {
		return puppy.Errorp(puppy.ErrNegativePuppyValueOnCreate, p.Value)
	}
	if p.ID != 0 {
		return puppy.Errorp(puppy.ErrPuppyAlreadyIdentified, p.ID)
	}
	m.nextID++
	p.ID = m.nextID
	m.pmap[p.ID] = *p
	return nil
}

func (m *MapStore) ReadPuppy(id int) (puppy.Puppy, error) {
	v, ok := m.pmap[id]
	if !ok {
		return puppy.Puppy{}, puppy.Errorp(puppy.ErrPuppyNotFoundOnRead, id)
	}
	return v, nil
}

func (m *MapStore) UpdatePuppy(p puppy.Puppy) error {
	if _, ok := m.pmap[p.ID]; !ok {
		return puppy.Errorp(puppy.ErrPuppyNotFoundOnUpdate, p.ID)
	}
	if p.Value < 0 {
		return puppy.Errorp(puppy.ErrNegativePuppyValueOnUpdate, p.Value)
	}
	m.pmap[p.ID] = p
	return nil
}

func (m *MapStore) DeletePuppy(id int) error {
	if _, ok := m.pmap[id]; !ok {
		return puppy.Errorp(puppy.ErrPuppyNotFoundOnDelete, id)
	}
	delete(m.pmap, id)
	return nil
}
