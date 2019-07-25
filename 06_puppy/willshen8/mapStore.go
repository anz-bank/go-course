package main

type MapStore struct {
	ms      map[uint32]Puppy
	counter uint32
}

// NewMapStore initialise a new MapStore
func NewMapStore() *MapStore {
	var newMapStore = MapStore{}
	newMapStore.ms = map[uint32]Puppy{}
	return &newMapStore
}

// IncrementCounter increase the ID counter everytime a new Puppy is created.
func (m *MapStore) IncrementCounter() {
	m.counter++
}

// CreatePuppy create a new puppy and store in mapStore.
func (m *MapStore) CreatePuppy(p *Puppy) uint32 {
	m.IncrementCounter()
	p.ID = m.counter
	m.ms[m.counter] = *p
	return p.ID
}

// ReadPuppy read a puppy given its id.
// It returns the pointer to that puppy.
func (m *MapStore) ReadPuppy(id uint32) *Puppy {
	if p, ok := m.ms[id]; ok {
		return &p
	}
	return nil
}

// UpdatePuppy updates the store with key of id with the new puppy.
// It returns a boolean whether the operation is successful or not.
func (m *MapStore) UpdatePuppy(id uint32, puppy *Puppy) bool {
	// if no existing puppy with this id, return false
	if _, ok := m.ms[id]; !ok {
		return false
	}
	puppy.ID = id
	m.ms[id] = *puppy
	return true
}

// DeletePuppy delete the puppy given the id.
// It returns true/success or false/unsuccessful.
func (m *MapStore) DeletePuppy(id uint32) bool {
	if _, found := m.ms[id]; found {
		delete(m.ms, id)
		return true
	}
	return false
}
