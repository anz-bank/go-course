package puppy

import "sync"

type MapStore struct {
	ms     map[uint32]Puppy
	nextID uint32
	mu     sync.Mutex
}

// NewMapStore initialise a new MapStore, ID starts at 1
func NewMapStore() *MapStore {
	return &MapStore{ms: map[uint32]Puppy{}, nextID: 1}
}

// CreatePuppy create a new puppy and store in mapStore.
func (m *MapStore) CreatePuppy(p *Puppy) (uint32, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if err := p.Validate(); err != nil {
		return 0, err
	}
	p.ID = m.nextID
	m.nextID++
	m.ms[p.ID] = *p
	return p.ID, nil
}

// ReadPuppy read a puppy given its id.
// It returns the pointer to that puppy.
func (m *MapStore) ReadPuppy(id uint32) (*Puppy, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if p, ok := m.ms[id]; ok {
		return &p, nil
	}
	return nil, Errorf(ErrNotFound, "Puppy ID (%v) not found", id)
}

// UpdatePuppy updates the store with key of id with the new puppy.
func (m *MapStore) UpdatePuppy(id uint32, p *Puppy) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.ms[id]; !ok {
		return Errorf(ErrNotFound, "Puppy ID can't be found, update operation failed")
	}
	if err := p.Validate(); err != nil {
		return err
	}
	p.ID = id
	m.ms[id] = *p
	return nil
}

// DeletePuppy delete the puppy given the id.
func (m *MapStore) DeletePuppy(id uint32) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.ms[id]; ok {
		delete(m.ms, id)
		return nil
	}
	return Errorf(ErrNotFound, "Puppy ID can't be found, delete operation failed")
}
