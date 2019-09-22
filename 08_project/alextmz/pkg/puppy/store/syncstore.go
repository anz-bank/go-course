package store

import (
	"sync"

	"github.com/anz-bank/go-course/08_project/alextmz/pkg/puppy"
)

type SyncStore struct {
	nextID int
	sync.Map
	sync.Mutex
}

func NewSyncStore() *SyncStore {
	a := SyncStore{}
	return &a
}

func (m *SyncStore) CreatePuppy(p *puppy.Puppy) error {
	if p == nil {
		return puppy.Error{Code: puppy.ErrNilPuppyPointer}
	}
	if p.Value < 0 {
		return puppy.Errorp(puppy.ErrNegativePuppyValueOnCreate, p.Value)
	}
	if p.ID != 0 {
		return puppy.Errorp(puppy.ErrPuppyAlreadyIdentified, p.ID)
	}
	m.Lock()
	defer m.Unlock()
	p.ID = m.nextID + 1
	m.nextID++
	m.Store(p.ID, *p)
	return nil
}

func (m *SyncStore) ReadPuppy(id int) (puppy.Puppy, error) {
	v, ok := m.Load(id)
	if !ok {
		return puppy.Puppy{}, puppy.Errorp(puppy.ErrPuppyNotFoundOnRead, id)
	}
	m.Lock()
	defer m.Unlock()
	return v.(puppy.Puppy), nil
}

func (m *SyncStore) UpdatePuppy(p puppy.Puppy) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.Load(p.ID); !ok {
		return puppy.Errorp(puppy.ErrPuppyNotFoundOnUpdate, p.ID)
	}
	if p.Value < 0 {
		return puppy.Errorp(puppy.ErrNegativePuppyValueOnUpdate, p.Value)
	}
	m.Store(p.ID, p)
	return nil
}

func (m *SyncStore) DeletePuppy(id int) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.Load(id); !ok {
		return puppy.Errorp(puppy.ErrPuppyNotFoundOnDelete, id)
	}
	m.Delete(id)
	return nil
}
