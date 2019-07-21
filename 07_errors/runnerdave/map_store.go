package main

// MapStore represents a simple map storage for the Puppy store
type MapStore struct {
	puppies map[uint16]Puppy
}

// NewMapStore creates a new in-memory store with map intialised
func NewMapStore() *MapStore {
	return &MapStore{puppies: map[uint16]Puppy{}}
}

// CreatePuppy saves new puppy if not in store, if it is already returns error
func (m *MapStore) CreatePuppy(p *Puppy) error {
	if err := validateValue(p.Value); err != nil {
		return err
	}
	if _, ok := m.puppies[p.ID]; ok {
		return Errorf(ErrUnknown, "puppy with id %d already exists", p.ID)
	}
	m.puppies[p.ID] = *p
	return nil
}

// ReadPuppy reads store by Puppy ID
func (m *MapStore) ReadPuppy(id uint16) (Puppy, error) {
	if puppy, ok := m.puppies[id]; ok {
		return puppy, nil
	}
	return Puppy{}, Errorf(ErrIDNotFound, "puppy with ID:%d not found", id)
}

// UpdatePuppy updates puppy with new value if ID present otherwise error
func (m *MapStore) UpdatePuppy(id uint16, p *Puppy) error {
	if err := validateValue(p.Value); err != nil {
		return err
	}
	if _, ok := m.puppies[id]; !ok {
		return Errorf(ErrIDNotFound, "puppy with ID:%d not found", id)
	}
	m.puppies[id] = *p
	return nil
}

// DeletePuppy deletes a puppy by id from the store
func (m *MapStore) DeletePuppy(id uint16) error {
	if _, ok := m.puppies[id]; ok {
		delete(m.puppies, id)
		return nil
	}
	return Errorf(ErrIDNotFound, "puppy with ID:%d not found", id)
}
