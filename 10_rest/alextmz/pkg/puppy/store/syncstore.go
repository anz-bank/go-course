package store

import (
	"sync"

	"github.com/anz-bank/go-course/10_rest/alextmz/pkg/puppy"
)

type SyncStore struct {
	size int
	sync.Map
	sync.Mutex
}

func NewSyncStore() *SyncStore {
	a := SyncStore{}
	return &a
}

// func CreatePuppy takes
func (m *SyncStore) CreatePuppy(p *puppy.Puppy) error {
	switch {
	case p == nil:
		return puppy.NewError(puppy.Err400BadRequest)
	case p.ID == 0:
		m.Lock()
		defer m.Unlock()
		p.ID = m.size + 1
		m.size++
		m.Store(p.ID, *p)
		return nil
	}
	return puppy.NewError(puppy.Err400BadRequest)
}

func (m *SyncStore) ReadPuppy(id int) (puppy.Puppy, error) {
	switch {
	case id < 0:
		return puppy.Puppy{}, puppy.NewError(puppy.Err400BadRequest)
	default:
		m.Lock()
		defer m.Unlock()
		if r, ok := m.Load(id); ok {
			puppy := r.(puppy.Puppy)
			return puppy, nil
		}
		return puppy.Puppy{}, puppy.NewError(puppy.Err404NotFound)
	}
}

func (m *SyncStore) UpdatePuppy(p puppy.Puppy) error {
	switch {
	case p.ID < 0:
		return puppy.NewError(puppy.Err400BadRequest)
	default:
		m.Lock()
		defer m.Unlock()
		if _, ok := m.Load(p.ID); ok {
			m.Store(p.ID, p)
			return nil
		}
		return puppy.NewError(puppy.Err404NotFound)
	}
}

func (m *SyncStore) DeletePuppy(id int) error {
	switch {
	case id < 0:
		return puppy.NewError(puppy.Err400BadRequest)
	default:
		m.Lock()
		defer m.Unlock()
		if _, ok := m.Load(id); ok {
			m.Delete(id)
			return nil
		}
		return puppy.NewError(puppy.Err404NotFound)
	}
}
