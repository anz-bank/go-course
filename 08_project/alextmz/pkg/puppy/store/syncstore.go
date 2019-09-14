package store

import (
	"sync"

	pp "github.com/anz-bank/go-course/08_project/alextmz/pkg/puppy"
)

type SyncStore struct {
	sync.Map
	sync.Mutex
}

func NewSyncStore() *SyncStore {
	a := SyncStore{}
	return &a
}

func (m *SyncStore) CreatePuppy(p *pp.Puppy) error {
	if _, ok := m.Load(p.ID); !ok {
		if p.Value < 0 {
			return pp.NewError(pp.ErrValueLessThanZero)
		}
		m.Lock()
		m.Store(p.ID, *p)
		m.Unlock()
		return nil
	}
	return pp.NewError(pp.ErrIDBeingCreatedAlreadyExists)
}

func (m *SyncStore) ReadPuppy(id pp.Pid) (*pp.Puppy, error) {
	r, ok := m.Load(id)
	if ok {
		puppy := r.(pp.Puppy)
		return &puppy, nil
	}
	return nil, pp.NewError(pp.ErrIDBeingReadDoesNotExist)
}

func (m *SyncStore) UpdatePuppy(id pp.Pid, p *pp.Puppy) error {
	if _, ok := m.Load(id); ok {
		if p.Value < 0 {
			return pp.NewError(pp.ErrValueLessThanZero)
		}
		m.Lock()
		m.Store(id, *p)
		m.Unlock()
		return nil
	}
	return pp.NewError(pp.ErrIDBeingUpdatedDoesNotExist)
}

func (m *SyncStore) DeletePuppy(id pp.Pid) error {
	_, ok := m.Load(id)
	if ok {
		m.Lock()
		m.Delete(id)
		m.Unlock()
		return nil
	}
	return pp.NewError(pp.ErrIDBeingDeletedDoesNotExist)
}
