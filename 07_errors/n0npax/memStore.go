package main

// NewMemStore creates new storer for map
func NewMemStore() *MemStore {
	return &MemStore{m: map[int]Puppy{}}
}

// CreatePuppy creates puppy
func (m *MemStore) CreatePuppy(p *Puppy) (int, error) {
	if p.Value < 0 {
		return -1, Errorf(ErrCodeInvalidInput, "puppy value have to be positive number")
	}
	id := m.nextID
	m.nextID++
	p.ID = id
	m.m[id] = *p
	return id, nil
}

// ReadPuppy reads puppy from backend
func (m *MemStore) ReadPuppy(id int) (*Puppy, error) {
	if puppy, ok := m.m[id]; ok {
		return &puppy, nil
	}
	return nil, Errorf(ErrCodeNotFound, "puppy with ID (%v) not found", id)
}

// UpdatePuppy updates puppy
func (m *MemStore) UpdatePuppy(id int, puppy *Puppy) error {
	if puppy.Value < 0 {
		return Errorf(ErrCodeInvalidInput, "puppy value have to be positive number")
	}
	if id != puppy.ID {
		return Errorf(ErrCodeInvalidInput, "ID is corrupted. Please ensure object ID matched provided ID")
	}
	if _, ok := m.m[id]; !ok {
		return Errorf(ErrCodeNotFound, "puppy with ID (%v) not found", id)
	}
	m.m[id] = *puppy
	return nil
}

// DeletePuppy deletes puppy
func (m *MemStore) DeletePuppy(id int) (bool, error) {
	if _, ok := m.m[id]; !ok {
		return false, Errorf(ErrCodeNotFound, "puppy with ID (%v) not found", id)
	}
	delete(m.m, id)
	return true, nil
}
