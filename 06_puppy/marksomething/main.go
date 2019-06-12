package main

import "sync"

// Puppy represents a small dog
type Puppy struct {
	ID     uint32
	Breed  string
	Colour string
	Value  int
}

// PuppyStorer stores an object in a key-value store
type PuppyStorer interface {
	CreatePuppy(p Puppy)
	ReadPuppy(k uint32) Puppy
	UpdatePuppy(k uint32, p Puppy)
	DeletePuppy(k uint32)
}

// MapStore is a storage mechanism for Puppy
type MapStore map[uint32]Puppy

// CreatePuppy adds a Puppy to storage unless a Puppy of the same ID is stored in which case nothing is done
func (m MapStore) CreatePuppy(p Puppy) {
	if _, found := m[p.ID]; found {
		return
	}
	m[p.ID] = p
}

// ReadPuppy returns a Puppy from storage, or a blank Puppy where no Puppy of that id is found
func (m MapStore) ReadPuppy(id uint32) Puppy {
	return m[id]
}

// UpdatePuppy updates an existing puppy if id matches the Puppy and id is found in the storage, otherwise nothing
func (m MapStore) UpdatePuppy(id uint32, p Puppy) {
	if id != p.ID {
		return
	}
	if _, ok := m[id]; !ok {
		return
	}
	m[id] = p
}

// DeletePuppy removes a puppy from storage if it exists
func (m MapStore) DeletePuppy(id uint32) {
	delete(m, id)
}

// SyncStore is a storage mechanism for Puppy
type SyncStore struct{ sync.Map }

// CreatePuppy adds a Puppy to storage unless a Puppy of the same ID is stored in which case nothing is done
func (s *SyncStore) CreatePuppy(p Puppy) {
	if _, found := s.Load(p.ID); found {
		return
	}
	s.Store(p.ID, p)
}

// ReadPuppy returns a Puppy from storage, or a blank Puppy where no Puppy of that id is found
func (s *SyncStore) ReadPuppy(id uint32) Puppy {
	p, ok := s.Load(id)
	if !ok {
		return Puppy{}
	}
	return p.(Puppy)
}

// UpdatePuppy updates an existing puppy if id matches the Puppy and id is found in the storage, otherwise nothing
func (s *SyncStore) UpdatePuppy(id uint32, p Puppy) {
	if id != p.ID {
		return
	}
	if _, ok := s.Load(id); !ok {
		return
	}
	s.Store(id, p)
}

// DeletePuppy removes a puppy from storage if it exists
func (s *SyncStore) DeletePuppy(id uint32) {
	s.Delete(id)
}

func main() {

}
